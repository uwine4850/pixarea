package authapi

import (
	"net/http"
	"strconv"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/interfaces/itypeopr"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/cookies"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/messages"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/utils/formutils"
)

type LoginForm struct {
	Username []string `form:"username"`
	Password []string `form:"password"`
}

func LoginPostHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	response := messages.SingleErrorResponse{}
	fillFormPtr, err := parseForm(r)
	if err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	loginForm := fillFormPtr.Ptr().(*LoginForm)

	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	defer func() {
		if err := db.Close(); err != nil {
			response.Error = err.Error()
			router.SendJson(response, w)
		}
	}()
	_auth := auth.NewAuth(db, w, manager)
	if err := loginUser(_auth, loginForm); err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	if err := userCookies(w, loginForm.Username[0], db); err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	return func() { router.SendJson(response, w) }
}

func parseForm(r *http.Request) (itypeopr.IPtr, error) {
	loginForm := &LoginForm{}
	loginFormPtr := typeopr.Ptr{}.New(loginForm)
	requiredFields, err := formmapper.FieldsName(loginFormPtr, []string{})
	if err != nil {
		return nil, err
	}
	if err := formutils.ParseForm(r, loginFormPtr, []string{}, requiredFields); err != nil {
		return nil, err
	}
	return loginFormPtr, nil
}

func loginUser(_auth *auth.Auth, loginForm *LoginForm) error {
	if _, err := _auth.LoginUser(loginForm.Username[0], loginForm.Password[0]); err != nil {
		return err
	}
	return nil
}

func userCookies(w http.ResponseWriter, username string, db *database.Database) error {
	authUser, err := auth.UserByUsername(db, username)
	if err != nil {
		return err
	}
	user, err := db.SyncQ().Select([]string{"id", "avatar"}, "user", dbutils.WHEquals(map[string]interface{}{"auth": authUser["id"]}, ""), 1)
	if err != nil {
		return err
	}
	id, err := dbutils.ParseInt(user[0]["id"])
	if err != nil {
		return err
	}
	cookies.SetStandartCookie(w, pnames.COOKIE_USER_USERNAME, username, "/", 0)
	cookies.SetStandartCookie(w, pnames.COOKIE_USER_AVATAR, dbutils.ParseString(user[0]["avatar"]), "/", 0)
	cookies.SetStandartCookie(w, pnames.COOKIE_USER_ID, strconv.Itoa(id), "/", 0)
	return nil
}
