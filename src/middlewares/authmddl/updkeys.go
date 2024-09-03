package authmddl

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/builtin_mddl"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
	"github.com/uwine4850/pixarea/src/cnf/messages"
	"github.com/uwine4850/pixarea/src/handlers/projcookies"
)

func UpdKeys(db *database.Database) middlewares.MddlFunc {
	return builtin_mddl.Auth("/api/login", db, func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
		if err != http.ErrNoCookie {
			projcookies.ClearAuthCookies(w)
			response := messages.SingleErrorResponse{Error: "Error update keys.", Redirect: "/login"}
			router.SendJson(response, w)
			middlewares.SkipNextPage(manager.OneTimeData())
		}
	})
}
