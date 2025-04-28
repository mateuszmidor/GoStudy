package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	webAuthn  *webauthn.WebAuthn
	err       error
	datastore = &Datastore{}
)

func main() {
	wconfig := &webauthn.Config{
		RPDisplayName: "gostudy-fido2",                    // Display Name for your site
		RPID:          "localhost",                        // Generally the FQDN for your site
		RPOrigins:     []string{"https://localhost:8888"}, // The origin URLs allowed for WebAuthn requests
	}

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
		http.ServeFile(w, r, "create.html")
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var requestData map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestData); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Print the received JSON in a formatted, indented way
		prettyJSON, err := json.MarshalIndent(requestData, "", "    ")
		if err != nil {
			http.Error(w, "Error formatting JSON", http.StatusInternalServerError)
			return
		}

		fmt.Println("Received JSON:")
		fmt.Println(string(prettyJSON))
	})

	fmt.Println("Starting server at port 8888")
	if err := http.ListenAndServeTLS(":8888", "./localhost/cert.pem", "./localhost/key.pem", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}

	// u := webauthn.User{}
}

func BeginRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BeginRegistration")

	user := datastore.GetUser() // Find or create the new user
	options, session, err := webAuthn.BeginRegistration(user)

	// handle errors if present
	if err != nil {
		JSONResponse(w, err, http.StatusInternalServerError)
		return
	}

	// store the sessionData values
	datastore.SaveSession(session)
	fmt.Printf("Session Data: %+v\n", session)

	JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options

	fmt.Println("BeginRegistration success")
}

func FinishRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FinishRegistration")

	user := datastore.GetUser() // Get the user
	fmt.Printf("User: %+v\n", user)

	// Get the session data stored from the function above
	session := datastore.GetSession()
	fmt.Printf("Session: %+v\n", session)

	credential, err := webAuthn.FinishRegistration(user, *session, r)

	// Handle Error and return.
	if err != nil {
		fmt.Printf("%+v\n", err)
		JSONResponse(w, err, http.StatusInternalServerError)
		return
	}

	// If creation was successful, store the credential object
	// Pseudocode to add the user credential.
	user.AddCredential(*credential)
	datastore.SaveUser(user)

	JSONResponse(w, "Registration Success", http.StatusOK) // Handle next steps
	fmt.Println("FinishRegistration success")
}

func BeginLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BeginLogin")

	user := datastore.GetUser() // Find the user

	options, session, err := webAuthn.BeginLogin(user)
	if err != nil {
		// Handle Error and return.

		return
	}

	// store the session values
	datastore.SaveSession(session)

	JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options
	fmt.Println("BeginLogin success")
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FinishLogin")

	user := datastore.GetUser() // Get the user

	// Get the session data stored from the function above
	session := datastore.GetSession()

	credential, err := webAuthn.FinishLogin(user, *session, r)
	if err != nil {
		// Handle Error and return.

		return
	}

	// Handle credential.Authenticator.CloneWarning

	// If login was successful, update the credential object
	// Pseudocode to update the user credential.
	user.UpdateCredential(*credential)
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
