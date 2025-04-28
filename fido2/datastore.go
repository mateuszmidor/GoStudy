package main

import (
	"encoding/base64"
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

var lonelyUser = &User{Name: "lonely-user", DisplayName: "lonely-user@example.com", ID: []byte(string(rune(1)))}
var session = &webauthn.SessionData{}

func (u *User) AddCredential(credential webauthn.Credential) {
	u.Credentials = append(u.Credentials, credential)
}

func (u *User) UpdateCredential(updatedCredential webauthn.Credential) {
	for i, credential := range u.Credentials {
		if string(credential.ID) == string(updatedCredential.ID) {
			u.Credentials[i] = updatedCredential
			return
		}
	}
	// If credential not found, print a warning and append it as a new credential
	fmt.Printf("Warning: Credential not found, appending as new credential [id=%v]\n", string(updatedCredential.ID))
	u.AddCredential(updatedCredential)
}

func (u *User) WebAuthnID() []byte {
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
	return lonelyUser
}

func (d *Datastore) SaveUser(user *User) {
	lonelyUser = user
}

func (d *Datastore) GetSession() *webauthn.SessionData {
	return session
}

func (d *Datastore) SaveSession(s *webauthn.SessionData) {
	encodedChallenge := base64.StdEncoding.EncodeToString([]byte(s.Challenge)) // challenge returned from browser is base64 encoded and compared with session as such
	encodedChallenge = strings.TrimRight(encodedChallenge, "=")                // remove the padding "=" characters to match what is returned from the browser
	s.Challenge = encodedChallenge
	session = s
}
