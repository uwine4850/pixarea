package authmddl

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/utils/fslice"
	"github.com/uwine4850/pixarea/src/handlers/hauth"
)

func AuthPermissions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) {
	urlPattern, ok := manager.OneTimeData().GetUserContext(namelib.URL_PATTERN)
	if !ok {
		router.ServerError(w, "url pattern not exist", manager.Config())
		return
	}
	if fslice.SliceContains([]string{"/login", "/login-post", "/register", "/register-post"}, urlPattern.(string)) {
		return
	}
	for i := 0; i < len(r.Cookies()); i++ {
		cookie := r.Cookies()[i]
		if cookie.Name == namelib.AUTH_COOKIE {
			return
		}
	}
	hauth.LogOutHNDL(w, r, manager)()
}
