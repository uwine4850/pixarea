package hauth

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
)

func LogOutHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	auth := &http.Cookie{
		Name:     "AUTH",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, auth)
	authDate := &http.Cookie{
		Name:     "AUTH_DATE",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, authDate)
	return func() {
		http.Redirect(w, r, "login", http.StatusFound)
	}
}
