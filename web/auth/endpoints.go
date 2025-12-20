package main

import (
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// set auth cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authtoken",
		Value:    createJWT(email),
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	// Respond with HTTP 200
	w.WriteHeader(http.StatusOK)
	log.Printf("login, email=%s, password=%s", email, password)
}

func logout(w http.ResponseWriter, r *http.Request) {
	// retrieve user email
	cookie, err := r.Cookie("authtoken")
	if err != nil {
		panic(err)
	}
	email := decodeJWT(cookie.Value)

	// Remove the auth cookie by setting MaxAge < 0
	http.SetCookie(w, &http.Cookie{
		Name:     "authtoken",
		Value:    "",
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	// Respond with HTTP 200
	w.WriteHeader(http.StatusOK)
	log.Printf("logout, email=%s", email)
}
