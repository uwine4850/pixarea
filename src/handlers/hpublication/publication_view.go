package hpublication

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
)

type Comment struct {
	Id            string `db:"id"`
	ReplyId       string `db:"reply_id"`
	PublicationId string `db:"publication_id"`
	AuthorId      string `db:"author_id"`
	Text          string `db:"text"`
	IsHide        string `db:"is_hide"`
	Author        hprofile.User
}

type PublicationView struct {
	object.ObjView
}

func (v *PublicationView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.ServerError(w, err.Error(), manager)
}

func (v *PublicationView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.ObjectContext, error) {
	dbInterface, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_DB)
	db := dbInterface.(*database.Database)
	publicationContextInterface, err := object.GetObjectContext(manager)
	if err != nil {
		return nil, err
	}
	publicationContext := publicationContextInterface["publication"].(PublicationDB)

	author, err := hprofile.GetUserByAuthId(db, publicationContext.Author)
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
	images, err := getPublicationImages(db, publicationContext.Id)
	if err != nil {
		return nil, err
	}
	isLike, err := LikeExist(db, publicationContext.Id, currentAuth.UID)
	if err != nil {
		return nil, err
	}
	comments, err := getComments(db, publicationContext.Id)
	if err != nil {
		return nil, err
	}
	return object.ObjectContext{
		"categories": categories,
		"author":     author,
		"likes":      likes,
		"isLike":     isLike,
		"images":     images,
		"comments":   comments,
	}, nil
}

// getPublicationCategories retrieving publication categories by their identifiers.
func getPublicationCategories(db *database.Database, categoriesId []string) ([]PublicationCategory, error) {
	categories := []PublicationCategory{}
	for i := 0; i < len(categoriesId); i++ {
		// If the category is NULL.
		if categoriesId[i] == "" {
			continue
		}
		category, err := db.SyncQ().QB().Select("*", pnames.CATEGORIES_TABLE).Where("id", "=", categoriesId[i], "LIMIT 1").Ex()
		if err != nil {
			return nil, err
		}
		if err := dbutils.DatabaseResultNotEmpty(category); err != nil {
			return nil, err
		}
		var categoryStruct PublicationCategory
		if err := dbmapper.FillStructFromDb(category[0], typeopr.Ptr{}.New(&categoryStruct)); err != nil {
			return nil, err
		}
		categories = append(categories, categoryStruct)
	}
	return categories, nil
}

// getLikeCount getting the number of likes on a post.
func getLikeCount(db *database.Database, publicationId string) (int, error) {
	count, err := db.SyncQ().QB().Select("*", pnames.PUBLICATION_LIKES_TABLE).Where("publication", "=", publicationId).Count().Ex()
	if err != nil {
		return 0, err
	}
	if err := dbutils.DatabaseResultNotEmpty(count); err != nil {
		return 0, err
	}
	likes, err := dbutils.ParseInt(count[0]["count"])
	if err != nil {
		return 0, err
	}
	return likes, nil
}

func getPublicationImages(db *database.Database, publicationId string) ([]string, error) {
	imagesPaths := []string{}
	images, err := db.SyncQ().QB().Select("*", pnames.PUBLCATION_IMAGES_TABLE).Where("publication", "=", publicationId).Ex()
	if err != nil {
		return nil, err
	}
	if err := dbutils.DatabaseResultNotEmpty(images); err != nil {
		return nil, err
	}
	for i := 0; i < len(images); i++ {
		imagesPaths = append(imagesPaths, dbutils.ParseString(images[i]["image_path"]))
	}
	return imagesPaths, nil
}

func getComments(db *database.Database, publicationId string) ([]Comment, error) {
	commentsList := []Comment{}
	comments, err := db.SyncQ().QB().Select("*", pnames.PUBLICATION_COMMENTS_TABLE).
		Where("publication_id", "=", publicationId, "AND", "reply_id IS NULL", "ORDER BY id DESC").Ex()
	if err != nil {
		return nil, err
	}
	if err := dbutils.DatabaseResultNotEmpty(comments); err != nil {
		return []Comment{}, nil
	}
	for i := 0; i < len(comments); i++ {
		var comm Comment
		if err := dbmapper.FillStructFromDb(comments[i], typeopr.Ptr{}.New(&comm)); err != nil {
			return nil, err
		}
		user, err := hprofile.GetUserByAuthId(db, comm.AuthorId)
		if err != nil {
			return nil, err
		}
		comm.Author = user
		commentsList = append(commentsList, comm)
	}
	return commentsList, nil
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
