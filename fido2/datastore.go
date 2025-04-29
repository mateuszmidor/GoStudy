package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
)

// User implements webauthn.User interface.
type User struct {
	ID          []byte
	Name        string
	DisplayName string
	Credentials []webauthn.Credential
}

type Datastore struct {
}

var lonelyUser = &User{Name: "lonely-user", DisplayName: "lonely-user@example.com", ID: []byte(string("ab"))}
var session = &webauthn.SessionData{}

func (u *User) AddCredential(credential webauthn.Credential) {
	printAsJSON("AddCredential:", credential)
	u.Credentials = append(u.Credentials, credential)
}

func (u *User) UpdateCredential(credential webauthn.Credential) {
	printAsJSON("UpdateCredential:", credential)
	for i, c := range u.Credentials {
		if string(c.ID) == string(credential.ID) {
			u.Credentials[i] = credential
			return
		}
	}
	// If credential not found, print a warning and append it as a new credential
	fmt.Printf("Warning: Credential not found, appending as new credential [id=%v]\n", string(credential.ID))
	u.AddCredential(credential)
}

func (u *User) WebAuthnID() []byte {
	// result := make([]byte, base64.RawStdEncoding.EncodedLen(len(u.ID)))
	// base64.RawStdEncoding.Encode(result, u.ID)
	// return result
	return u.ID
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (d *Datastore) GetUser() *User {
	ret := *lonelyUser
	return &ret
}

func (d *Datastore) SaveUser(user *User) {
	printAsJSON("SaveUser:", user)
	*lonelyUser = *user
}

func (d *Datastore) GetSession() *webauthn.SessionData {
	return session
}

func (d *Datastore) SaveSession(s *webauthn.SessionData) {
	// challenge returned from browser is base64 encoded without trailing '=' padding chars:
	// https://developer.mozilla.org/en-US/docs/Web/API/AuthenticatorResponse/clientDataJSON#challenge
	encodedChallenge := base64.URLEncoding.EncodeToString([]byte(s.Challenge))
	encodedChallenge = strings.TrimRight(encodedChallenge, "=")
	s.Challenge = encodedChallenge
	session = s
	printAsJSON("SaveSession:", s)
}

func printAsJSON(header string, v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
		return
	}
	fmt.Println(header)
	fmt.Println(string(jsonData))
}
