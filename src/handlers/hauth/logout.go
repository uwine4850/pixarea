package hauth

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
)

func LogOutHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	auth := &http.Cookie{
		Name:     namelib.AUTH_COOKIE,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, auth)
	authDate := &http.Cookie{
		Name:     namelib.AUTH_DATE_COOKIE,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, authDate)
	userUsername := &http.Cookie{
		Name:     pnames.COOKIE_USER_USERNAME,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, userUsername)
	userAvatar := &http.Cookie{
		Name:     pnames.COOKIE_USER_AVATAR,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, userAvatar)
	return func() {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
