package hprofile

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/pixarea/src/cnf"
)

type User struct {
	Id          string `db:"id"`
	AuthId      string `db:"auth"`
	Name        string `db:"name"`
	Avatar      string `db:"avatar"`
	BgImage     string `db:"bg_image"`
	Description string `db:"description"`
	Auth        auth.User
}

type ProfileView struct {
	object.ObjView
}

func (v *ProfileView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) object.ObjectContext {
	context, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_CONTEXT)
	user, ok := context.(object.ObjectContext)["profile"].(User)
	if ok {
		db := database.NewDatabase(cnf.DB_ARGS)
		if err := db.Connect(); err != nil {
			v.OnError(w, r, manager, err)
			return nil
		}
		defer func() {
			if err := db.Close(); err != nil {
				v.OnError(w, r, manager, err)
			}
		}()
		authDb, err := db.SyncQ().Select([]string{"*"}, "auth", dbutils.WHEquals(map[string]interface{}{"id": user.AuthId}, ""), 1)
		if err != nil {
			v.OnError(w, r, manager, err)
			return nil
		}
		var auth auth.User
		if err := dbutils.FillStructFromDb(authDb[0], &auth); err != nil {
			v.OnError(w, r, manager, err)
			return nil
		}
		user.Auth = auth
	}
	return object.ObjectContext{"profile": user}
}

func (v *ProfileView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.ServerError(w, err.Error(), manager.Config())
}

func ObjectProfileViewHNDL() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DB_ARGS)
	view := object.TemplateView{
		TemplatePath: "src/templates/profile.html",
		View: &ProfileView{
			object.ObjView{
				Name:       "profile",
				DB:         db,
				TableName:  "user",
				FillStruct: User{},
				Slug:       "id",
			},
		},
	}
	return view.Call
}
