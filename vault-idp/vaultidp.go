package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func createUser(ctx context.Context, client *vault.Client, name, pass string) (*vault.Response[map[string]interface{}], error) {
	vals := url.Values{}
	vals.Add("token_ttl", "60")
	vals.Add("token_max_ttl", "60")
	vals.Add("token_explicit_max_ttl", "60")
	opts := vault.WithQueryParameters(vals)
	req := schema.UserpassWriteUserRequest{Password: pass}
	rsp, err := client.Auth.UserpassWriteUser(ctx, name, req, opts)
	fmt.Printf("UserpassWriteUser resp: %+v\n", rsp)
	return rsp, err
}

func loginUser(ctx context.Context, client *vault.Client, userName, userPass string) (*vault.Response[map[string]interface{}], error) {
	res, err := client.Auth.UserpassLogin(ctx, userName, schema.UserpassLoginRequest{Password: userPass})
	if err != nil {
		return nil, err
	}
	fmt.Printf("UserpassLogin resp: %+v\n", *res.Auth)
	return res, nil
}

func setUserMetadata(ctx context.Context, client *vault.Client, entityID string, metadata map[string]any) (*vault.Response[map[string]interface{}], error) {
	// just overwrite with new metadata
	updateEntityReq := schema.EntityUpdateByIdRequest{Metadata: metadata}
	return client.Identity.EntityUpdateById(ctx, entityID, updateEntityReq)
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

func createScope(ctx context.Context, client *vault.Client, name, description, template string) (*vault.Response[map[string]interface{}], error) {
	scopeReq := schema.OidcWriteScopeRequest{
		Description: description,
		Template:    template,
	}
	return client.Identity.OidcWriteScope(ctx, name, scopeReq)
}

func createOIDCProvider(ctx context.Context, client *vault.Client, issuer string, scopes []string) (*vault.Response[map[string]interface{}], error) {
	providerReq := schema.OidcWriteProviderRequest{
		AllowedClientIds: []string{"*"},
		ScopesSupported:  scopes,
		Issuer:           issuer, // optional, "" for default
	}
	return client.Identity.OidcWriteProvider(ctx, providerName, providerReq)
}

func createAppIntegration(ctx context.Context, client *vault.Client, name, redirectURL string) (*vault.Response[map[string]interface{}], error) {
	clientReq := schema.OidcWriteClientRequest{
		RedirectUris:   []string{redirectURL},
		Assignments:    []string{"allow_all"},
		Key:            "default",
		IdTokenTtl:     "300",
		AccessTokenTtl: "3600",
	}
	return client.Identity.OidcWriteClient(ctx, name, clientReq)
}
