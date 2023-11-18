package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/zitadel/oidc/v2/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v2/pkg/http"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

// Dictionary:
//	 IDP - Identity Provider; the server that holds the user profile

var (
	loginPath         = "/"
	loginCallbackPath = "/auth/callback"
	refreshTokenPath  = "/auth/refresh"
	revokeTokenPath   = "/auth/revoke" // after calling, refreshToken gets invalidated
	userInfoPath      = "/auth/userinfo"
)
var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	clientKey    = os.Getenv("CLIENT_KEY")
	clientKeyID  = os.Getenv("CLIENT_KEY_ID")
	issuer       = os.Getenv("ISSUER")
	port         = os.Getenv("PORT")
	scopes       = strings.Split(os.Getenv("SCOPES"), " ")
)

var (
	secretKey     = []byte("test1234test1234")
	redirectURI   = fmt.Sprintf("http://localhost:%v%v", port, loginCallbackPath)
	cookieHandler = httphelper.NewCookieHandler(secretKey, secretKey, httphelper.WithUnsecure())
)

var authTokens *oidc.Tokens[*oidc.IDTokenClaims] // holds logged-in user tokens
var relyingParty rp.RelyingParty

func main() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)

	// relyingParty represents the Identity Provider client
	relyingParty = makeRelyingParty()

	// register the AuthURLHandler at your preferred path.
	// the AuthURLHandler creates the auth request and redirects the user to the auth server.
	// including state handling with secure cookie and the possibility to use PKCE.
	// Prompts can optionally be set to inform the server of
	// any messages that need to be prompted back to the user.
	http.Handle(loginPath, rp.AuthURLHandler(stateFunc, relyingParty)) // , rp.WithPromptURLParam("Welcome back!") is not supported by AzureAD

	// register the CodeExchangeHandler at the callbackPath
	// the CodeExchangeHandler handles the auth response, creates the token request and calls the callback function
	// with the returned tokens from the token endpoint
	// in this example the callback function itself is wrapped by the UserinfoCallback which
	// will call the Userinfo endpoint, check the sub and pass the info into the callback function
	http.Handle(loginCallbackPath, CodeExchangeHandler(rp.UserinfoCallback(handleAuthTokens), relyingParty))

	// register the OAuth2 RefreshToken handler (this is optional, only if you plan to use token refresh flow)
	http.HandleFunc(refreshTokenPath, handleRefreshToken)

	// register revoke token handler, so the refresh token can be manually invalidated
	http.HandleFunc(revokeTokenPath, handleRevokeAccessRefreshToken)

	// register UserInfo handler - this prints information about logged in user
	http.HandleFunc(userInfoPath, handleUserInfo)

	lis := fmt.Sprintf("localhost:%s", port)
	logrus.Infof("listening on http://%s/", lis)
	logrus.Info("press ctrl+c to stop")
	logrus.Fatal(http.ListenAndServe(lis, nil))
}

// here we receive the auth tokens from IDP. We could also use the "info" param here, but instead we redirect the user to userinfo path
func handleAuthTokens(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty, info *oidc.UserInfo) {
	// store the received tokens for future use
	authTokens = tokens

	// and redirect to user info page
	http.Redirect(w, r, userInfoPath, http.StatusSeeOther)
}

// request a new access token from IDP by providing refresh token reveived on logging in
func handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("handling refresh token")

	newToken, err := RefreshToken(relyingParty, os.Getenv("CLIENT_SECRET"), authTokens.RefreshToken)
	if err != nil {
		logrus.Error("failed to refresh token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store the received tokens for future use
	authTokens.Token = newToken

	// and redirect to user info page
	http.Redirect(w, r, userInfoPath, http.StatusSeeOther)
}

func handleRevokeAccessRefreshToken(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("handling revoke token")

	var err error

	err = RevokeToken(relyingParty, authTokens.AccessToken, TokenHintAccess)
	if err != nil {
		logrus.Error("failed to revoke access token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = RevokeToken(relyingParty, authTokens.RefreshToken, TokenHintRefresh)
	if err != nil {
		logrus.Error("failed to revoke refresh token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("tokens revoked successfuly"))
}

// print UserInfo for the currently logged-in user
func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("handling user info")

	info, err := getUserInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// print UserInfo as a response
	data, err := yaml.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)

	// break line
	w.Write([]byte("\n\n"))

	// also print the tokens for introspection
	data, err = yaml.Marshal(*authTokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// generate some state (representing the state of the user in your application,
// e.g. the page where he was before sending him to login
func stateFunc() string {
	return uuid.New().String()
}

func makeRelyingParty() rp.RelyingParty {
	// prefer keys as more secure
	if os.Getenv("CLIENT_KEY") != "" && os.Getenv("CLIENT_KEY_ID") != "" {
		return makeRelyingPartyX509KeysAuth()
	}
	// use client secret as secondary option
	if os.Getenv("CLIENT_SECRET") != "" {
		return makeRelyingPartyClientSecretAuth()
	}
	logrus.Fatal("should export either CLIENT_KEY+CLIENT_KEY_ID or CLIENT_SECRET")
	return nil
}

// this function creates sort of OIDC client with ClientSecret auth
func makeRelyingPartyClientSecretAuth() rp.RelyingParty {
	logrus.Info("authenticating client in IDP using ClientSecret")

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}

	relyingParty, err := rp.NewRelyingPartyOIDC(issuer, clientID, clientSecret, redirectURI, scopes, options...)
	if err != nil {
		logrus.Fatalf("error creating provider: %s", err.Error())
	}
	return relyingParty
}

// this function creates sort of OIDC client with x.509 private PEM key auth
// to get the keys:
// - generate new key pair in JSON format: https://mkjwk.org/
// - convert public JSON to PEM: https://8gwifi.org/jwkconvertfunctions.jsp
func makeRelyingPartyX509KeysAuth() rp.RelyingParty {
	logrus.Info("authenticating client in IDP using user_assertion(PrivateKey)")

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
		rp.WithJWTProfile(rp.SignerFromKeyAndKeyID([]byte(clientKey), clientKeyID)),
		rp.WithHTTPClient(basicAuthRemovingClient()),
	}

	relyingParty, err := rp.NewRelyingPartyOIDC(issuer, clientID, "", redirectURI, scopes, options...)
	if err != nil {
		logrus.Fatalf("error creating provider: %s", err.Error())
	}

	return relyingParty
}

// getUserInfo returns data about the authenticated user.
// There are 2 sources of UserInfo data:
// 1. authentication token; the UserInfo data may be incomplete, but it works better in case of AzureAD
// 2. dedicated "UserInfo" endpoint in IDP server; supposed to always return full UserInfo, doesn't return user groups for AzureAD (thanks, Microsoft!)
func getUserInfo() (*oidc.UserInfo, error) {
	// get UserInfo from authentication token
	// return extractUserInfo(authTokens.IDTokenClaims) // doesnt return family_name nor given_name for Okta!

	// get UserInfo from dedicated endpoint in IDP
	return rp.Userinfo(authTokens.AccessToken, authTokens.TokenType, authTokens.IDTokenClaims.GetSubject(), relyingParty) // doesnt return groups for Azure!
}

// extractUserInfo makes UserInfo from the OIDC ID token
func extractUserInfo(claims *oidc.IDTokenClaims) (*oidc.UserInfo, error) {
	if claims.UserInfoEmail.Email == "" || len(claims.Claims) == 0 {
		return nil, fmt.Errorf("the OIDC ID token claims are missing Email and Claims")
	}
	return &oidc.UserInfo{
		Subject:         claims.Subject,
		UserInfoEmail:   claims.UserInfoEmail,
		UserInfoProfile: claims.UserInfoProfile,
		UserInfoPhone:   claims.UserInfoPhone,
		Claims:          claims.Claims,
		// TODO: copy the rest of data from 'claims' if necessary
	}, nil
}
