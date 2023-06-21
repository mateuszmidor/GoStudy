package main

import (
	"context"
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

var (
	loginPath         = "/"
	loginCallbackPath = "/auth/callback"
	refreshTokenPath  = "/auth/refresh"
	userInfoPath      = "/auth/userinfo"
)

var authTokens *oidc.Tokens[*oidc.IDTokenClaims] // holds logged-in user tokens

func main() {
	logrus.SetReportCaller(true)

	// register the AuthURLHandler at your preferred path.
	// the AuthURLHandler creates the auth request and redirects the user to the auth server.
	// including state handling with secure cookie and the possibility to use PKCE.
	// Prompts can optionally be set to inform the server of
	// any messages that need to be prompted back to the user.
	http.Handle(loginPath, rp.AuthURLHandler(stateFunc, makeRelyingParty(), rp.WithPromptURLParam("Welcome back!")))

	// register the CodeExchangeHandler at the callbackPath
	// the CodeExchangeHandler handles the auth response, creates the token request and calls the callback function
	// with the returned tokens from the token endpoint
	// in this example the callback function itself is wrapped by the UserinfoCallback which
	// will call the Userinfo endpoint, check the sub and pass the info into the callback function
	http.Handle(loginCallbackPath, rp.CodeExchangeHandler(rp.UserinfoCallback(handleAuthTokens), makeRelyingParty()))

	// register the OAuth2 RefreshToken handler (this is optional, only if you plan to use token refresh flow)
	http.HandleFunc(refreshTokenPath, handleRefreshToken)

	// register UserInfo handler - this prints information about logged in user
	http.HandleFunc(userInfoPath, handleUserInfo)

	port := os.Getenv("PORT")
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
	// request new access token
	tmpToken := *authTokens.Token
	tmpToken.Expiry = time.Now() // force token expiration so the tokenSource always returns new token
	tokenSource := makeRelyingParty().OAuthConfig().TokenSource(context.TODO(), &tmpToken)
	newToken, err := tokenSource.Token()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store the received tokens for future use
	authTokens.Token = newToken

	// and redirect to user info page
	http.Redirect(w, r, userInfoPath, http.StatusSeeOther)
}

// print UserInfo for the currently logged-in user
func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	// request UserInfo from IDP using the previously obtained access token
	info, err := rp.Userinfo(authTokens.AccessToken, authTokens.TokenType, authTokens.IDTokenClaims.GetSubject(), makeRelyingParty())
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

// this function creates sort of OIDC client
func makeRelyingParty() rp.RelyingParty {
	secretKey := []byte("test1234test1234")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	keyPath := os.Getenv("KEY_PATH")
	issuer := os.Getenv("ISSUER")
	port := os.Getenv("PORT")
	scopes := strings.Split(os.Getenv("SCOPES"), " ")

	redirectURI := fmt.Sprintf("http://localhost:%v%v", port, loginCallbackPath)
	cookieHandler := httphelper.NewCookieHandler(secretKey, secretKey, httphelper.WithUnsecure())

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	if keyPath != "" {
		options = append(options, rp.WithJWTProfile(rp.SignerFromKeyPath(keyPath)))
	}

	relyingParty, err := rp.NewRelyingPartyOIDC(issuer, clientID, clientSecret, redirectURI, scopes, options...)
	if err != nil {
		logrus.Fatalf("error creating provider %s", err.Error())
	}
	return relyingParty
}
