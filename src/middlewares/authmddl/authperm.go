package authmddl

import (
	"fmt"
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
	"github.com/uwine4850/foozy/pkg/utils/fslice"
	"github.com/uwine4850/pixarea/src/cnf/messages"
	"github.com/uwine4850/pixarea/src/handlers/projcookies"
)

func AuthPermissions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) {
	singleErrorResponse := messages.SingleErrorResponse{}
	urlPattern, ok := manager.OneTimeData().GetUserContext(namelib.ROUTER.URL_PATTERN)
	if !ok {
		singleErrorResponse.Error = "url pattern not exist"
		sendError(&singleErrorResponse, w, manager.OneTimeData())
		fmt.Println("1")
		return
	}
	if fslice.SliceContains([]string{"/login", "/login-post", "/register", "/register-post", "/api/login"}, urlPattern.(string)) {
		return
	}
	for i := 0; i < len(r.Cookies()); i++ {
		cookie := r.Cookies()[i]
		if cookie.Name == namelib.AUTH.COOKIE_AUTH {
			return
		}
	}
	singleErrorResponse.Redirect = "/login"
	projcookies.ClearAuthCookies(w)
	sendError(&singleErrorResponse, w, manager.OneTimeData())
	fmt.Println("2")
}

func sendError(singleErrorResponse *messages.SingleErrorResponse, w http.ResponseWriter, manager interfaces.IManagerOneTimeData) {
	router.SendJson(singleErrorResponse, w)
	middlewares.SkipNextPage(manager)
}
