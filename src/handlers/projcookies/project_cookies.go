package projcookies

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router/cookies"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
)

func ClearAuthCookies(w http.ResponseWriter) {
	cookies.SetStandartCookie(w, namelib.AUTH.COOKIE_AUTH, "", "/", -1)
	cookies.SetStandartCookie(w, namelib.AUTH.COOKIE_AUTH_DATE, "", "/", -1)
	cookies.SetStandartCookie(w, pnames.COOKIE_USER_USERNAME, "", "/", -1)
	cookies.SetStandartCookie(w, pnames.COOKIE_USER_AVATAR, "", "/", -1)
}
