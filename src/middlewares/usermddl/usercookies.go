package usermddl

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hauth"
)

func ParseUserCookies(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) {
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
