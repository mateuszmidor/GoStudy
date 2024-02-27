package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
)

var userName = "user"
var userPass = "pass"
var providerName = "my-provider"
var webappName = "my-webapp"

func main() {
	log.SetFlags(log.Lshortfile)
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	logFatalOnError(err)

	// authenticate the client with a root token (insecure), so it can create users and OIDC providers
	err = client.SetToken("root-token")
	logFatalOnError(err)

	// create new user
	_, err = createUser(ctx, client, userName, userPass)
	logFatalOnError(err)

	// login user to get user information
	res, err := loginUser(ctx, client, userName, userPass)
	logFatalOnError(err)

	// setup user metadata
	metadata := map[string]any{
		"email":      "JOHN.doe@acme.com",
		"givenname":  "John",
		"familyname": "Doe",
	}
	_, err = setUserMetadata(ctx, client, res.Auth.EntityID, metadata)
	logFatalOnError(err)

	// create user groups
	entityIDs := []string{res.Auth.EntityID}
	_, err = createGroup(ctx, client, "users", entityIDs)
	logFatalOnError(err)
	_, err = createGroup(ctx, client, "admins", entityIDs)
	logFatalOnError(err)

	// create OIDC scopes
	emailTemplate := `{"email": {{identity.entity.metadata.email}}}`
	_, err = createScope(ctx, client, "email", "email", emailTemplate)
	logFatalOnError(err)
	givennameTemplate := `{"givenname": {{identity.entity.metadata.givenname}}}`
	_, err = createScope(ctx, client, "givenname", "givenname", givennameTemplate)
	logFatalOnError(err)
	familynameTemplate := `{"familyname": {{identity.entity.metadata.familyname}}}`
	_, err = createScope(ctx, client, "familyname", "familyname", familynameTemplate)
	logFatalOnError(err)
	groupsTemplate := `{"groups": {{identity.entity.groups.names}}}`
	_, err = createScope(ctx, client, "groups", "groups", groupsTemplate)
	logFatalOnError(err)

	// create OIDC provider
	scopes := []string{"email", "givenname", "familyname", "groups"}
	_, err = createOIDCProvider(ctx, client, "", scopes)
	logFatalOnError(err)

	// create OIDC client app
	redirectURL := "http://localhost:8000/auth/callback"
	_, err = createAppIntegration(ctx, client, webappName, redirectURL)
	logFatalOnError(err)

	// print OIDC client app credentials to use
	mywebapp, err := client.Identity.OidcReadClient(ctx, webappName)
	logFatalOnError(err)
	fmt.Println("ClientID:", mywebapp.Data["client_id"])
	fmt.Println("ClientSecret:", mywebapp.Data["client_secret"])

	// print OIDC issuer URL to use
	provider, err := client.Identity.OidcReadProvider(ctx, providerName)
	logFatalOnError(err)
	fmt.Println("IssuerURL:", provider.Data["issuer"])
}

func logFatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
