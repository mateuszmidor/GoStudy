package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zitadel/oidc/v2/pkg/oidc"

	oidc_client "github.com/zitadel/oidc/v2/pkg/client"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"golang.org/x/oauth2"
)

// RefreshToken returns a new Token in exchange for refreshToken, using either client_assertion(priv-pub x.509 keys) or ClientSecret to authenticate.
// golang.org/x/oauth2 doesn't expose interface to refresh a token using client_assertion as client authentication method in IDP.
// This is work-around for that, using relyingParty.OAuthConfig().Exchange(...) to access oauth2-internal retrieveToken machinery
func RefreshToken(relyingParty rp.RelyingParty, clientSecret, refreshToken string) (*oauth2.Token, error) {
	const noClientSecret = "" // don't specify client secret when client_assertion is used for client authentication

	// ask authorization server for token refresh
	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("grant_type", "refresh_token"),
		oauth2.SetAuthURLParam("refresh_token", refreshToken),
	}

	// OAuthConfig().Exchange uses HttpClient extracted from ctx and we use dedicated Authrorization:basic header removing client for x.509 keys auth so need to set it here
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, relyingParty.HttpClient())

	// client authenticates itself using client_assertion (x.509 private-public keys)
	if relyingParty.Signer() != nil {
		audience := []string{
			// relyingParty.Issuer(),                        // original audience by Zitadel
			relyingParty.OAuthConfig().Endpoint.TokenURL, // but Okta requires TokenURL in audience, and this is in accordance with client_assertion spec
		}
		assertion, err := oidc_client.SignedJWTProfileAssertion(relyingParty.OAuthConfig().ClientID, audience, time.Hour, relyingParty.Signer())
		if err != nil {
			return nil, fmt.Errorf("failed to build assertion for token refresh: %w", err)
		}
		opts = append(opts, rp.WithClientAssertionJWT(assertion)()...)
		return relyingParty.OAuthConfig().Exchange(ctx, noClientSecret, opts...)
	}

	// client authenticates itself using client secret
	return relyingParty.OAuthConfig().Exchange(context.Background(), clientSecret, opts...)
}

type revokeRequest struct {
	Token         string `schema:"token"`
	TokenTypeHint string `schema:"token_type_hint"`
	ClientID      string `schema:"client_id"`
	// either ClientSecret, or ClientAssertion + ClientAssertionType should be supplied. Never mix them
	ClientSecret        string `schema:"client_secret,omitempty"`
	ClientAssertion     string `schema:"client_assertion,omitempty"`
	ClientAssertionType string `schema:"client_assertion_type,omitempty"`
}

type tokenHint string

// TokenHintAccess indicates to the IDP that access token should be revoked
const TokenHintAccess tokenHint = "id_token"

// TokenHintRefresh indicates to the IDP that refresh token should be revoked
const TokenHintRefresh tokenHint = "refresh_token"

// RevokeToken invalidates either access token or refresh token, using either client_assertion(priv-pub x.509 keys) or ClientSecret to authenticate.
// golang.org/x/oauth2 doesn't expose interface to revoke a token using client_assertion as client authentication method in IDP.
// This is a modification of "RevokeToken" from github.com/zitadel/oidc/v2/pkg/client/rp/relying_party.go
func RevokeToken(rp rp.RelyingParty, token string, hint tokenHint) error {
	revokeCaller, ok := rp.(oidc_client.RevokeCaller)
	if !ok && revokeCaller.GetRevokeEndpoint() == "" {
		return fmt.Errorf("RelyingParty does not support RevokeCaller")
	}

	request := revokeRequest{
		Token:         token,
		TokenTypeHint: string(hint),
		ClientID:      rp.OAuthConfig().ClientID,
	}

	// client authenticates itself using client_assertion (x.509 private-public keys)
	if rp.Signer() != nil {
		audience := []string{
			// rp.Issuer(),            // original audience by Zitadel
			rp.GetRevokeEndpoint(), // but Okta requires TokenURL in audience, and this is in accordance with client_assertion spec
		}
		assertion, err := oidc_client.SignedJWTProfileAssertion(rp.OAuthConfig().ClientID, audience, time.Hour, rp.Signer())
		if err != nil {
			return fmt.Errorf("failed to build assertion for token revoke: %w", err)
		}
		request.ClientAssertion = assertion
		request.ClientAssertionType = oidc.ClientAssertionTypeJWTAssertion
		return oidc_client.CallRevokeEndpoint(request, nil, revokeCaller)
	}

	// client authenticates itself using client secret
	request.ClientSecret = rp.OAuthConfig().ClientSecret
	return oidc_client.CallRevokeEndpoint(request, nil, revokeCaller)
}
