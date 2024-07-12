package handlers

import (
	"net/http"
	"time"

	"github.com/hculpan/osetools/internal/auth"
)

// LoginPostHandler handles the login form submission
func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if !auth.ValidateUser(username, password) {
		http.Redirect(w, r, "/login?error=invalid_credentials", http.StatusSeeOther)
		return
	}

	tokenString, err := auth.GenerateJWT(username)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(30 * 24 * time.Hour), // 30 days
		HttpOnly: true,                                // Enhance security by limiting access to the cookie from client-side scripts
	})

	// Redirect to the welcome page upon successful login
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
