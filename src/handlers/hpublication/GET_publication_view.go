package hpublication

import (
	"net/http"

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
	currentAuth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return nil, err
	}

	// Handle publication info.
	categorySubView := CategorySubView{
		DB: db,
	}
	categories, err := categorySubView.CategoriesById([]string{publicationContext.Category1, publicationContext.Category2})
	if err != nil {
		return nil, err
	}
	images, err := getPublicationImages(db, publicationContext.Id)
	if err != nil {
		return nil, err
	}

	// Handle likes.
	likeSubView := LikeSubView{
		PublicationId: publicationContext.Id,
		DB:            db,
	}
	likes, err := likeSubView.GetLikeCount()
	if err != nil {
		return nil, err
	}
	isLike, err := likeSubView.LikeExist(currentAuth.UID)
	if err != nil {
		return nil, err
	}

	// Handle comments.
	commentSubView := CommentSubView{
		PublicationId: publicationContext.Id,
		DB:            db,
	}
	comments, err := commentSubView.GetComments()
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
