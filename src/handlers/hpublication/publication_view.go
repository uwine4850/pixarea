package hpublication

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
)

type PublicationView struct {
	object.ObjView
}

func (v *PublicationView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.ServerError(w, err.Error(), manager)
}

func (v *PublicationView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.ObjectContext, error) {
	dbInterface, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_DB)
	db := dbInterface.(*database.Database)
	publicationContextInterface, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_CONTEXT)
	publicationContext := publicationContextInterface.(object.ObjectContext)["publication"].(PublicationDB)
	author, err := getParentUser(db, publicationContext.Author)
	if err != nil {
		return nil, err
	}
	categories, err := getPublicationCategories(db, []string{publicationContext.Category1, publicationContext.Category2})
	if err != nil {
		return nil, err
	}
	likes, err := getLikeCount(db, publicationContext.Id)
	if err != nil {
		return nil, err
	}
	currentAuth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return nil, err
	}
	isLike, err := LikeExist(db, publicationContext.Id, currentAuth.UID)
	if err != nil {
		return nil, err
	}
	return object.ObjectContext{"categories": categories, "author": author, "likes": likes, "isLike": isLike}, nil
}

// getParentUser In the database, the publication is bound to the auth table.
// This function gets information about the user based on his auth ID.
func getParentUser(db *database.Database, authId any) (hprofile.User, error) {
	authUser, err := auth.UserByID(db, authId)
	if err != nil {
		return hprofile.User{}, err
	}
	if authUser == nil {
		return hprofile.User{}, fmt.Errorf("auth by id %s not found", authId)
	}
	var auth auth.AuthItem
	if err := dbutils.FillStructFromDb(authUser, &auth); err != nil {
		return hprofile.User{}, err
	}

	user, err := db.SyncQ().QB().Select("*", "user").Where("auth", "=", authId, "LIMIT 1").Ex()
	if err != nil {
		return hprofile.User{}, err
	}
	if len(user) != 1 {
		return hprofile.User{}, fmt.Errorf("user by auth id %s not found", authId)
	}
	var userStruct hprofile.User
	if err := dbutils.FillStructFromDb(user[0], &userStruct); err != nil {
		return hprofile.User{}, err
	}
	userStruct.Auth = auth
	return userStruct, nil
}

// getPublicationCategories retrieving publication categories by their identifiers.
func getPublicationCategories(db *database.Database, categoriesId []string) ([]PublicationCategory, error) {
	categories := make([]PublicationCategory, len(categoriesId))
	for i := 0; i < len(categoriesId); i++ {
		category, err := db.SyncQ().QB().Select("*", pnames.CATEGORIES_TABLE).Where("id", "=", categoriesId[i], "LIMIT 1").Ex()
		if err != nil {
			return nil, err
		}
		if len(category) != 1 {
			return nil, fmt.Errorf("category by id %s not found", categoriesId[i])
		}
		var categoryStruct PublicationCategory
		if err := dbutils.FillStructFromDb(category[0], &categoryStruct); err != nil {
			return nil, err
		}
		categories[i] = categoryStruct
	}
	return categories, nil
}

// getLikeCount getting the number of likes on a post.
func getLikeCount(db *database.Database, publicationId string) (int, error) {
	count, err := db.SyncQ().QB().Select("*", pnames.PUBLICATION_LIKES_TABLE).Where("publication", "=", publicationId).Count().Ex()
	if err != nil {
		return 0, err
	}
	if len(count) != 1 {
		return 0, errors.New("error getting number of likes")
	}
	likes, err := dbutils.ParseInt(count[0]["COUNT(*)"])
	if err != nil {
		return 0, err
	}
	return likes, nil
}

func PublicationViewHNDL() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	tv := object.TemplateView{
		TemplatePath: "src/templates/publication/publication_view.html",
		View: &PublicationView{
			object.ObjView{
				Name:       "publication",
				DB:         database.NewDatabase(cnf.DB_ARGS),
				TableName:  "publication",
				FillStruct: PublicationDB{},
				Slug:       "id",
			},
		},
	}
	return tv.Call
}
