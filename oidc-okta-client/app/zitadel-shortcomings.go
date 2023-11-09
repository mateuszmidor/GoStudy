package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	oidc_client "github.com/zitadel/oidc/v2/pkg/client"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

// authAwareHttpTransport removes "Authorization" : "Basic " headers from incoming requests
type authAwareHttpTransport struct {
	http.Transport
}

func (c *authAwareHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	lowercaseAuthorizationHeader := strings.ToLower(req.Header.Get("Authorization"))
	if strings.Contains(lowercaseAuthorizationHeader, "basic ") { // only remove basic and not bearer!
		logrus.Debug("removing Authorization header")
		req.Header.Del("Authorization")
	}
	return c.Transport.RoundTrip(req)
}

// relyingParty should not use BasicAuth (Authorization in header) when using x509 keys (client_assertion in body) - Okta shouts:
// oauth2: "invalid_request" "Cannot supply multiple client credentials. Use one of the following: credentials in the Authorization header, credentials in the post body, or a client_assertion in the post body."
// such situation happens when first user authenticates with server like "https://dev-39423526.okta.com" using ClientSecret (header auth)
// and next user tries to authenticate with the same server with x509 keys (body auth)
// this is because  golang.org/x/oauth2 keeps internal cache mapping URLs to AuthMethods, and the first successfuly used method gets used forever,
// so when second user tries to authenticate, header auth gets set, but also x509 auth gets set in body, and so Okta screams :)
// golang.org/x/oauth2 seems not really ready for x509 keys client authentication as of July 2023
func basicAuthRemovingClient() *http.Client {
	return &http.Client{Transport: &authAwareHttpTransport{}}
}

// Below function is copied from github.com/zitadel/oidc/v2/pkg/client/rp and modified: add TokenURL in "audience" JWT claim
func CodeExchangeHandler[C oidc.IDClaims](callback rp.CodeExchangeCallback[C], client rp.RelyingParty, urlParam ...rp.URLParamOpt) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := tryReadStateCookie(w, r, client)
		if err != nil {
			logrus.Errorf("failed to get state: %v", err)
			http.Error(w, "failed to get state: "+err.Error(), http.StatusUnauthorized)
			return
		}
		params := r.URL.Query()
		if params.Get("error") != "" {
			logrus.Errorf("failure returned from IDP: %s", params.Get("error"))
			client.ErrorHandler()(w, r, params.Get("error"), params.Get("error_description"), state)
			return
		}
		codeOpts := make([]rp.CodeExchangeOpt, 0)
		for i, p := range urlParam {
			codeOpts[i] = rp.CodeExchangeOpt(p)
		}

		if client.IsPKCE() {
			// BEGIN MODIFIED
			const pkceCode = "pkce"
			// END MODIFIED
			codeVerifier, err := client.CookieHandler().CheckCookie(r, pkceCode)
			if err != nil {
				http.Error(w, "failed to get code verifier: "+err.Error(), http.StatusUnauthorized)
				return
			}
			codeOpts = append(codeOpts, rp.WithCodeVerifier(codeVerifier))
		}
		if client.Signer() != nil {
			// BEGIN MODIFIED
			audience := []string{
				// client.Issuer(),                        // original audience by Zitadel
				client.OAuthConfig().Endpoint.TokenURL, // but Okta requires TokenURL in audience, and this is in accordance with client_assertion spec
			}
			assertion, err := oidc_client.SignedJWTProfileAssertion(client.OAuthConfig().ClientID, audience, time.Hour, client.Signer())
			// END MODIFIED
			if err != nil {
				http.Error(w, "failed to build assertion: "+err.Error(), http.StatusUnauthorized)
				return
			}
			codeOpts = append(codeOpts, rp.WithClientAssertionJWT(assertion))
		}
		tokens, err := rp.CodeExchange[C](r.Context(), params.Get("code"), client, codeOpts...)
		if err != nil {
			logrus.Errorf("failed to exchange token: %v", err)
			http.Error(w, "failed to exchange token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		callback(w, r, tokens, state, client)
	}
}

// Below function is copied from github.com/zitadel/oidc/v2/pkg/client/rp
func tryReadStateCookie(w http.ResponseWriter, r *http.Request, rp rp.RelyingParty) (state string, err error) {
	// BEGIN MODIFIED
	const stateParam = "state"
	// END MODIFIED
	if rp.CookieHandler() == nil {
		return r.FormValue(stateParam), nil
	}
	state, err = rp.CookieHandler().CheckQueryCookie(r, stateParam)
	if err != nil {
		return "", err
	}
	rp.CookieHandler().DeleteCookie(w, stateParam)
	return state, nil
}
