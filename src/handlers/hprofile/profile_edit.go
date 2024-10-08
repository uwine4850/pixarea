package hprofile

import (
	"net/http"
	"os"
	"reflect"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/interfaces/itypeopr"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/cookies"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
)

type ProfileEditView struct {
	object.ObjView
}

func (v *ProfileEditView) Permissions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (bool, func()) {
	hashKey := manager.Config().Key().Get32BytesKey().HashKey()
	blockKey := manager.Config().Key().Get32BytesKey().BlockKey()
	var cookieAuth auth.AuthCookie
	if err := cookies.ReadSecureCookieData([]byte(hashKey), []byte(blockKey), r, namelib.AUTH.COOKIE_AUTH, &cookieAuth); err != nil {
		return false, func() {
			router.ServerForbidden(w, manager)
		}
	}
	context, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_CONTEXT)
	user, ok := context.(object.ObjectContext)["profile"].(User)
	if ok {
		if user.AuthId != cookieAuth.UID {
			return false, func() {
				router.ServerForbidden(w, manager)
			}
		}
	} else {
		return false, func() {
			router.ServerForbidden(w, manager)
		}
	}
	return true, func() {}
}

func (v *ProfileEditView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.ObjectContext, error) {
	context, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_CONTEXT)
	user, ok := context.(object.ObjectContext)["profile"].(User)
	if ok {
		db, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_DB)
		authDb, err := auth.UserByID(db.(*database.Database), user.AuthId)
		if err != nil {
			return nil, err
		}
		var auth auth.AuthItem
		if err := dbmapper.FillStructFromDb(authDb, typeopr.Ptr{}.New(&auth)); err != nil {
			return nil, err
		}
		user.Auth = auth
	}
	return object.ObjectContext{"profile": user}, nil
}

func (v *ProfileEditView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.ServerError(w, err.Error(), manager)
}

func ObjectProfileEditViewHNDL() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DB_ARGS)
	view := object.TemplateView{
		TemplatePath: "src/templates/profile/profile_edit.html",
		View: &ProfileEditView{
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

type ProfileEditForm struct {
	Avatar           []form.FormFile `form:"avatar"`
	DeleteAvatar     []string        `form:"delete_avatar"`
	Background       []form.FormFile `form:"background"`
	DeleteBackground []string        `form:"delete_background"`
	Name             []string        `form:"name"`
	Description      []string        `form:"description"`
}

func ProfileEditPostHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	var profileForm ProfileEditForm
	fillFormPtr := typeopr.Ptr{}.New(&profileForm)
	if err := formmapper.FillStructFromForm(frm, fillFormPtr, []string{"delete_avatar", "delete_background"}); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}

	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	defer func() {
		if err := db.Close(); err != nil {
			router.ServerError(w, err.Error(), manager)
		}
	}()

	authUID, err := getUIDFromCookie(r, manager)
	if err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}

	removedImages := []string{}
	createImages := []string{}
	profileDbData := map[string]any{}

	if err := handleImages(w, db, &authUID, fillFormPtr, &profileForm, &profileDbData, &removedImages, &createImages, manager); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	profileDbData["name"] = profileForm.Name[0]
	profileDbData["description"] = profileForm.Description[0]
	if _, err := db.SyncQ().Update("user", profileDbData, dbutils.WHEquals(dbutils.WHValue{"auth": authUID.UID}, "")); err != nil {
		if err := rollbackCreateImages(createImages); err != nil {
			return func() { router.ServerError(w, err.Error(), manager) }
		}
		cookies.SetStandartCookie(w, pnames.COOKIE_USER_AVATAR, "", "/", 0)
	}
	if err := removeImages(removedImages); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	return func() {
		http.Redirect(w, r, "/explore", http.StatusFound)
	}
}

func handleImages(w http.ResponseWriter, db *database.Database, authUID *auth.AuthCookie, fillForm itypeopr.IPtr, profileForm *ProfileEditForm,
	profileDbData *map[string]any, removedImages *[]string, createImages *[]string, manager interfaces.IManager) error {
	var profileFromDb User
	prifileEditForm := fillForm.Ptr().(*ProfileEditForm)
	if prifileEditForm.DeleteAvatar != nil || prifileEditForm.DeleteBackground != nil {
		user, err := db.SyncQ().Select([]string{"*"}, "user", dbutils.WHEquals(dbutils.WHValue{"auth": authUID.UID}, "AND"), 1)
		if err != nil {
			return err
		}
		if err := dbmapper.FillStructFromDb(user[0], typeopr.Ptr{}.New(&profileFromDb)); err != nil {
			return err
		}
	}

	if prifileEditForm.DeleteAvatar != nil {
		(*profileDbData)["avatar"] = ""
		if profileFromDb.Avatar != "" {
			*removedImages = append(*removedImages, profileFromDb.Avatar)
		}
		cookies.SetStandartCookie(w, pnames.COOKIE_USER_AVATAR, "", "/", 0)
	} else {
		if !reflect.ValueOf(prifileEditForm.Avatar).IsZero() {
			var avatarPath string
			if err := form.SaveFile(w, profileForm.Avatar[0].Header, "src/media/avatars", &avatarPath, manager); err != nil {
				return err
			}
			(*profileDbData)["avatar"] = avatarPath
			*createImages = append(*createImages, avatarPath)
			cookies.SetStandartCookie(w, pnames.COOKIE_USER_AVATAR, avatarPath, "/", 0)
		}
	}
	if prifileEditForm.DeleteBackground != nil {
		(*profileDbData)["bg_image"] = ""
		if profileFromDb.BgImage != "" {
			*removedImages = append(*removedImages, profileFromDb.BgImage)
		}
	} else {
		if !reflect.ValueOf(prifileEditForm.Background).IsZero() {
			var backgroundPath string
			if err := form.SaveFile(w, profileForm.Background[0].Header, "src/media/backgrounds", &backgroundPath, manager); err != nil {
				return err
			}
			(*profileDbData)["bg_image"] = backgroundPath
			*createImages = append(*createImages, backgroundPath)
		}
	}
	return nil
}

func getUIDFromCookie(r *http.Request, manager interfaces.IManager) (auth.AuthCookie, error) {
	hashKey := manager.Config().Key().Get32BytesKey().HashKey()
	blockKey := manager.Config().Key().Get32BytesKey().BlockKey()
	var cookieAuth auth.AuthCookie
	if err := cookies.ReadSecureCookieData([]byte(hashKey), []byte(blockKey), r, namelib.AUTH.COOKIE_AUTH, &cookieAuth); err != nil {
		return auth.AuthCookie{}, err
	}
	return cookieAuth, nil
}

func rollbackCreateImages(createFiles []string) error {
	for i := 0; i < len(createFiles); i++ {
		if err := os.Remove(createFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func removeImages(removedFiles []string) error {
	for i := 0; i < len(removedFiles); i++ {
		if err := os.Remove(removedFiles[i]); err != nil {
			return err
		}
	}
	return nil
}
