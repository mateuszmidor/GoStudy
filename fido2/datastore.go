package main

import (
	"encoding/json"
	"fmt"

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
	return u.ID
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	credentialsCopy := make([]webauthn.Credential, len(u.Credentials))
	copy(credentialsCopy, u.Credentials)
	return credentialsCopy
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
	ret := *session
	return &ret
}

func (d *Datastore) SaveSession(s *webauthn.SessionData) {
	*session = *s
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
