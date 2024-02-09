package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
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
	_, err = createOIDCProvider(ctx, client, scopes)
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

func createGroup(ctx context.Context, client *vault.Client, name string, entityIDs []string) (*vault.Response[map[string]interface{}], error) {
	groupReq := schema.GroupCreateRequest{
		Name:            name,
		MemberEntityIds: entityIDs,
		Type:            "internal",
	}
	gres, err := client.Identity.GroupCreate(ctx, groupReq)
	return gres, err
}

func logFatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createOIDCProvider(ctx context.Context, client *vault.Client, scopes []string) (*vault.Response[map[string]interface{}], error) {
	providerReq := schema.OidcWriteProviderRequest{
		AllowedClientIds: []string{"*"},
		ScopesSupported:  scopes,
	}
	return client.Identity.OidcWriteProvider(ctx, providerName, providerReq)
}

func createAppIntegration(ctx context.Context, client *vault.Client, name, redirectURL string) (*vault.Response[map[string]interface{}], error) {
	clientReq := schema.OidcWriteClientRequest{
		RedirectUris:   []string{redirectURL},
		Assignments:    []string{"allow_all"},
		Key:            "default",
		IdTokenTtl:     300,
		AccessTokenTtl: 3600,
	}
	return client.Identity.OidcWriteClient(ctx, name, clientReq)
}

func setUserMetadata(ctx context.Context, client *vault.Client, entityID string, metadata map[string]any) (*vault.Response[map[string]interface{}], error) {
	// just overwrite with new metadata
	updateEntityReq := schema.EntityUpdateByIdRequest{Metadata: metadata}
	return client.Identity.EntityUpdateById(ctx, entityID, updateEntityReq)
}

func loginUser(ctx context.Context, client *vault.Client, name, pass string) (*vault.Response[map[string]interface{}], error) {
	res, err := client.Auth.UserpassLogin(ctx, userName, schema.UserpassLoginRequest{Password: userPass})
	if err != nil {
		return nil, err
	}
	fmt.Printf("UserpassLogin resp: %+v\n", *res.Auth)
	return res, nil
}

func createUser(ctx context.Context, client *vault.Client, name, pass string) (*vault.Response[map[string]interface{}], error) {
	vals := url.Values{}
	vals.Add("token_ttl", "60")
	vals.Add("token_max_ttl", "60")
	vals.Add("token_explicit_max_ttl", "60")
	opts := vault.WithCustomQueryParameters(vals)
	req := schema.UserpassWriteUserRequest{Password: pass}
	rsp, err := client.Auth.UserpassWriteUser(ctx, name, req, opts)
	fmt.Printf("UserpassWriteUser resp: %+v\n", rsp)
	return rsp, err
}

func createScope(ctx context.Context, client *vault.Client, name, description, template string) (*vault.Response[map[string]interface{}], error) {
	scopeReq := schema.OidcWriteScopeRequest{
		Description: description,
		Template:    template,
	}
	return client.Identity.OidcWriteScope(ctx, name, scopeReq)
}
