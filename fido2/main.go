package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	webAuthn  *webauthn.WebAuthn
	datastore = &Datastore{}
)

func main() {
	wconfig := &webauthn.Config{
		RPDisplayName: "gostudy-fido2",                    // Display Name for your site
		RPID:          "localhost",                        // Generally the FQDN for your site
		RPOrigins:     []string{"https://localhost:8888"}, // The origin URLs allowed for WebAuthn requests
	}

	var err error
	if webAuthn, err = webauthn.New(wconfig); err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/webauthn/begin-registration", BeginRegistration)
	http.HandleFunc("/webauthn/finish-registration", FinishRegistration)
	http.HandleFunc("/webauthn/begin-login", BeginLogin)
	http.HandleFunc("/webauthn/finish-login", FinishLogin)

	// my code
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// server must be TLS for webauthn to work at all, otherwise in JS navigator.credentials is undefined
	fmt.Println("Starting server at port 8888")
	if err := http.ListenAndServeTLS(":8888", "./localhost/cert.pem", "./localhost/key.pem", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func BeginRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BeginRegistration")

	user := datastore.GetUser() // Find or create the new user
	options, session, err := webAuthn.BeginRegistration(user)

	// handle errors if present
	if err != nil {
		fmt.Printf("Failed to BeginRegistration%+v\n", err)
		JSONResponse(w, err, http.StatusInternalServerError)
		return
	}

	// store the sessionData to be retrieved in FinishRegistration
	datastore.SaveSession(session)

	// return the options to frontend to be used when registering user with authenticator (e.g. TouchID, samsung pass)
	JSONResponse(w, options, http.StatusOK) // options.publicKey contain our registration options

	fmt.Println("BeginRegistration success")
}

func FinishRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FinishRegistration")

	user := datastore.GetUser() // Get the user

	// Get the session data stored from BeginRegistration
	session := datastore.GetSession()

	// challenge returned from browser is url base64 encoded without trailing '=' padding chars, so need to base64encode the session challenge to make them comparable in webAuthn.FinishRegistration
	// https://developer.mozilla.org/en-US/docs/Web/API/AuthenticatorResponse/clientDataJSON#challenge
	session.Challenge = base64encodeString(session.Challenge)
	credential, err := webAuthn.FinishRegistration(user, *session, r)

	// Handle errors
	if err != nil {
		fmt.Printf("Failed to FinishRegistration: %+v\n", err)
		JSONResponse(w, err, http.StatusInternalServerError)
		return
	}

	// If creation was successful, store the credential object with the user for later logging in
	user.AddCredential(*credential)
	datastore.SaveUser(user)

	JSONResponse(w, "Registration Success", http.StatusOK)

	fmt.Println("FinishRegistration success")
}

func BeginLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BeginLogin")

	user := datastore.GetUser() // Find the user
	options, session, err := webAuthn.BeginLogin(user)
	if err != nil {
		fmt.Printf("Failed to BeginLogin: %+v\n", err)
		JSONResponse(w, err, http.StatusInternalServerError)
		return
	}

	// store the session to be used in FinishLogin
	datastore.SaveSession(session)

	// return the options to frontend to be used when loging in with authenticator
	JSONResponse(w, options, http.StatusOK) // options.publicKey contain our login options

	fmt.Println("BeginLogin success")
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FinishLogin")

	user := datastore.GetUser()          // Get the user
	originalUserID := user.ID            // remember the ID
	user.ID = base64encodeBytes(user.ID) // webAuthn compares the user.ID with base64-encoded userID received from browser, so base64encode to make them comparable in webAuthn.FinishLogin

	// Get the session data stored from BeginLogin
	session := datastore.GetSession()
	session.UserID = base64encodeBytes(session.UserID)        // must match user.ID so need to base64encode
	session.Challenge = base64encodeString(session.Challenge) // webAuthn.FinishLogin compares the base64encodedchallenge, so base64encode here

	credential, err := webAuthn.FinishLogin(user, *session, r)
	if err != nil {
		fmt.Printf("Failed to FinishLogin: %+v\n", err)
		JSONResponse(w, err, http.StatusInternalServerError)
		return
	}

	if credential.Authenticator.CloneWarning {
		fmt.Println("WARNING: the authenticator has been cloned, so more than single copy of the private key exists in the world. Continuing...")
	}

	// If login was successful, update the credential object
	user.UpdateCredential(*credential)
	user.ID = originalUserID // restore the non-base64encoded ID
	datastore.SaveUser(user)

	JSONResponse(w, "Login Success", http.StatusOK)

	fmt.Println("FinishLogin success")
}

func JSONResponse(w http.ResponseWriter, val interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(val); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func base64encodeBytes(input []byte) []byte {
	output := make([]byte, base64.RawStdEncoding.EncodedLen(len(input)))
	base64.RawStdEncoding.Encode(output, input)
	fmt.Println("encoding", string(input), "->", string(output))
	return output
}

func base64encodeString(input string) string {
	// output := base64.RawURLEncoding.EncodeToString([]byte(input))
	// fmt.Println("encoding", input, "->", output)
	// return output
	return input
}
