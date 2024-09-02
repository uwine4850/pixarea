package tokenapi

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/pixarea/src/cnf/messages"
)

func CSRFToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	response := messages.CSRFTokenResponse{}
	token, err := r.Cookie(namelib.ROUTER.COOKIE_CSRF_TOKEN)
	if err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	response.Token = token.Value
	return func() { router.SendJson(response, w) }
}

func CSRFTokenOptions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	return func() {}
}
