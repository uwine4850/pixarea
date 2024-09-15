package authapi

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/interfaces/itypeopr"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/messages"
	"github.com/uwine4850/pixarea/src/utils/formutils"
)

type RegisterForm struct {
	Name           []string `form:"name"`
	Username       []string `form:"username"`
	Password       []string `form:"password"`
	RepeatPassword []string `form:"repeat_password"`
}

func RegisterPostHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	response := messages.SingleErrorResponse{}
	registerFormPtr, err := parseRegisterForm(r)
	if err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	registerForm := registerFormPtr.Ptr().(*RegisterForm)

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
	if err := registerUser(w, db, manager, registerForm); err != nil {
		response.Error = err.Error()
		return func() { router.SendJson(response, w) }
	}
	return func() { router.SendJson(response, w) }
}

func parseRegisterForm(r *http.Request) (itypeopr.IPtr, error) {
	registerForm := &RegisterForm{}
	registerFormPtr := typeopr.Ptr{}.New(registerForm)
	requiredFields, err := formmapper.FieldsName(registerFormPtr, []string{})
	if err != nil {
		return nil, err
	}
	if err := formutils.ParseForm(r, registerFormPtr, []string{}, requiredFields); err != nil {
		return nil, err
	}
	if registerForm.Password[0] != registerForm.RepeatPassword[0] {
		return nil, err
	}
	return registerFormPtr, nil
}

func registerUser(w http.ResponseWriter, db *database.Database, manager interfaces.IManager, registerForm *RegisterForm) error {
	_auth := auth.NewAuth(db, w, manager)
	if err := _auth.RegisterUser(registerForm.Username[0], registerForm.Password[0]); err != nil {
		return err
	}
	registerUserDB, err := auth.UserByUsername(db, registerForm.Username[0])
	if err != nil {
		return err
	}
	if _, err := db.SyncQ().Insert("user", map[string]interface{}{"auth": registerUserDB["id"], "name": registerForm.Name[0]}); err != nil {
		if err := rollBackAuth(registerUserDB, db); err != nil {
			return err
		}
		return err
	}
	return nil
}

func rollBackAuth(registerUserDB map[string]interface{}, db *database.Database) error {
	_, err := db.SyncQ().Delete("auth", dbutils.WHEquals(map[string]interface{}{"id": registerUserDB["id"]}, ""))
	if err != nil {
		return err
	}
	return nil
}
