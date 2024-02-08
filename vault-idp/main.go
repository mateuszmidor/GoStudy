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
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("root-token"); err != nil {
		log.Fatal(err)
	}

	// // create/get user
	// ereq := schema.EntityCreateRequest{
	// 	Name: userName,
	// }
	// entity, err := client.Identity.EntityCreate(ctx, ereq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if entity == nil {
	// 	entity, err = client.Identity.EntityReadByName(ctx, userName)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// secReq := schema.MountsEnableSecretsEngineRequest{
	// 	Type: "passthrough",
	// }
	// _, err = client.System.MountsEnableSecretsEngine(ctx, "auth", secReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	req := schema.UserpassWriteUserRequest{
		Password: userPass,
	}
	vals := url.Values{}
	vals.Add("token_ttl", "60")
	vals.Add("token_max_ttl", "60")
	vals.Add("token_explicit_max_ttl", "60")
	opts := vault.WithCustomQueryParameters(vals)
	_, err = client.Auth.UserpassWriteUser(ctx, userName, req, opts)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Auth.UserpassLogin(ctx, userName, schema.UserpassLoginRequest{Password: userPass}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *res.Auth)

	// create OIDC client app
	clientReq := schema.OidcWriteClientRequest{
		RedirectUris:   []string{"http://localhost:8000/auth/callback"},
		Assignments:    []string{"allow_all"},
		Key:            "default",
		IdTokenTtl:     300,
		AccessTokenTtl: 3600,
	}
	_, err = client.Identity.OidcWriteClient(ctx, webappName, clientReq)
	if err != nil {
		log.Fatal(err)
	}

	// create scopes
	emailScope := schema.OidcWriteScopeRequest{
		Description: "email",
		Template:    `{"email": "mat@acme.com" }`,
	}
	_, err = client.Identity.OidcWriteScope(ctx, "email", emailScope)
	if err != nil {
		log.Fatal(err)
	}
	// create OIDC provider
	providerReq := schema.OidcWriteProviderRequest{
		AllowedClientIds: []string{"*"},
		ScopesSupported:  []string{"email"},
	}
	_, err = client.Identity.OidcWriteProvider(ctx, providerName, providerReq)
	if err != nil {
		log.Fatal(err)
	}

	// print oidc client app credentials to use
	mywebapp, err := client.Identity.OidcReadClient(ctx, webappName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ClientID:", mywebapp.Data["client_id"])
	fmt.Println("ClientSecret:", mywebapp.Data["client_secret"])

	// print oidc issuer URL to use
	provider, err := client.Identity.OidcReadProvider(ctx, providerName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IssuerURL:", provider.Data["issuer"])

	// add ALIAS. is this needed?
	// fmt.Printf("%+v\n", entity)
	// entityID := entity.Data["id"].(string)
	// fmt.Println("entityID:", entityID)
	// areq := schema.EntityCreateAliasRequest{
	// 	CanonicalId: entityID,
	// 	// Id:            entityID,
	// 	MountAccessor: "auth_userpass_97588e92",
	// 	Name:          name,
	// }
	// _, err = client.Identity.EntityCreateAlias(ctx, areq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
