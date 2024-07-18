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
	Auth        auth.AuthItem
}

type ProfileView struct {
	object.ObjView
}

func (v *ProfileView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.ObjectContext, error) {
	context, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_CONTEXT)
	user, ok := context.(object.ObjectContext)["profile"].(User)
	if ok {
		db, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_DB)
		authDb, err := auth.UserByID(db.(*database.Database), user.AuthId)
		if err != nil {
			return nil, err
		}
		var auth auth.AuthItem
		if err := dbutils.FillStructFromDb(authDb, &auth); err != nil {
			return nil, err
		}
		user.Auth = auth
	}
	return object.ObjectContext{"profile": user}, nil
}

func (v *ProfileView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.ServerError(w, err.Error(), manager)
}

func ObjectProfileViewHNDL() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DB_ARGS)
	view := object.TemplateView{
		TemplatePath: "src/templates/profile/profile.html",
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
