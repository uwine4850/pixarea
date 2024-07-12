package usermddl

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/utils/fslice"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hauth"
)

func ParseUserCookies(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) {
	urlPattern, ok := manager.OneTimeData().GetUserContext(namelib.URL_PATTERN)
	if !ok {
		router.ServerError(w, "url pattern not exist", manager)
		return
	}
	if fslice.SliceContains([]string{"/login", "/login-post", "/register", "/register-post"}, urlPattern.(string)) {
		return
	}
	userUsername, err := r.Cookie(pnames.COOKIE_USER_USERNAME)
	if err != nil {
		hauth.LogOutHNDL(w, r, manager)()
		return
	}
	userAvatar, err := r.Cookie(pnames.COOKIE_USER_AVATAR)
	if err != nil {
		hauth.LogOutHNDL(w, r, manager)()
		return
	}
	manager.Render().SetContext(map[string]interface{}{"userUsername": userUsername.Value, "userAvatar": userAvatar.Value})
}
