package hpublication

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
)

type PublicationDB struct {
	Author      string `db:"author"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Category1   any    `db:"category1"`
	Category2   any    `db:"category2"`
	Date        string `db:"date"`
}

type PublicationCategory struct {
	Name string `db:"name"`
}

type NewPublicationForm struct {
	Name        []string        `form:"name"`
	Description []string        `form:"description"`
	Categories  []string        `form:"catedories"`
	Images      []form.FormFile `form:"images"`
}

func PublicationViewHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	manager.Render().SetTemplatePath("src/templates/publication/publication_view.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	return func() {}
}

func NewPublicationPageHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	router.CatchRedirectError(r, manager)
	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	defer func() {
		if err := db.Close(); err != nil {
			router.ServerError(w, err.Error(), manager)
		}
	}()
	categoriesDB, err := db.SyncQ().QB().Select("*", pnames.CATEGORIES_TABLE).Ex()
	if err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	categories := []PublicationCategory{}
	for i := 0; i < len(categoriesDB); i++ {
		var category PublicationCategory
		if err := dbutils.FillStructFromDb(categoriesDB[i], &category); err != nil {
			return func() { router.ServerError(w, err.Error(), manager) }
		}
		categories = append(categories, category)
	}

	manager.Render().SetContext(map[string]interface{}{"categories": categories})
	manager.Render().SetTemplatePath("src/templates/publication/new_publication.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	return func() {}
}

type NewPublicationView struct {
	object.FormView
}

func (v *NewPublicationView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w)
}

func (v *NewPublicationView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.ObjectContext, error) {
	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return nil, err
	}
	defer func() {
		if err := db.Close(); err != nil {
			v.OnError(w, r, manager, err)
		}
	}()
	context, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_CONTEXT)
	pubForm := context.(object.ObjectContext)[namelib.OBJECT_CONTEXT_FORM].(NewPublicationForm)
	filledForm := form.NewFillableFormStruct(&pubForm)

	currentAuth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return nil, err
	}

	if len(pubForm.Categories) > 2 {
		return nil, errors.New("you can select up to 2 categories")
	}
	categories, err := getCategories(db, pubForm.Categories)
	if err != nil {
		return nil, err
	}
	db.BeginTransaction()
	publication := PublicationDB{
		Author:      currentAuth.UID,
		Name:        filledForm.GetOrDef("Name", 0).(string),
		Description: filledForm.GetOrDef("Description", 0).(string),
		Category1:   categories[0],
		Category2:   categories[1],
		Date:        time.Now().Format("2006-01-02 15:04:05"),
	}
	newPublicationID, err := createPublicationInTable(db, publication)
	if err != nil {
		if err := db.RollBackTransaction(); err != nil {
			return nil, err
		}
		return nil, err
	}
	createdImages, err := createImages(w, pubForm.Images, manager)
	if err != nil {
		return nil, err
	}
	if err := saveImages(db, newPublicationID, &createdImages); err != nil {
		return nil, err
	}
	if err := db.CommitTransaction(); err != nil {
		return nil, err
	}

	router.SendJson(map[string]string{"success": "true"}, w)
	return object.ObjectContext{}, nil
}

func createPublicationInTable(db *database.Database, publicationDb PublicationDB) (any, error) {
	publicationParams, err := dbutils.ParamsValueFromStruct(&publicationDb)
	if err != nil {
		return nil, err
	}
	if _, err := db.SyncQ().Insert(pnames.PUBLICATIONS_TABLE, publicationParams); err != nil {
		return nil, err
	}
	newPublicationID, err := db.SyncQ().Select([]string{"id"}, pnames.PUBLICATIONS_TABLE, dbutils.WHEquals(
		dbutils.WHValue{
			"author": publicationDb.Author,
			"date":   publicationDb.Date,
		}, "AND"), 0,
	)
	if err != nil {
		return nil, err
	}
	return newPublicationID[0]["id"], nil
}

func getCategories(db *database.Database, categoriesName []string) ([]any, error) {
	categories := make([]any, 2)
	for i := 0; i < len(categoriesName); i++ {
		categoryID, err := db.SyncQ().Select([]string{"id"}, pnames.CATEGORIES_TABLE,
			dbutils.WHEquals(dbutils.WHValue{"name": categoriesName[i]}, ""), 1)
		if err != nil {
			return nil, err
		}
		if len(categoryID) != 1 {
			return nil, nil
		}
		id, err := dbutils.ParseInt(categoryID[0]["id"])
		if err != nil {
			return nil, err
		}
		categories[i] = strconv.Itoa(id)
	}
	return categories, nil
}

func createImages(w http.ResponseWriter, images []form.FormFile, manager interfaces.IManager) ([]string, error) {
	imagesPath := []string{}
	for i := 0; i < len(images); i++ {
		var path string
		if err := form.SaveFile(w, images[i].Header, "src/media/publications", &path, manager); err != nil {
			return nil, err
		}
		imagesPath = append(imagesPath, path)
	}
	return imagesPath, nil
}

func saveImages(db *database.Database, publicationID any, imagesPath *[]string) error {
	asyncKeys := []string{}
	for i := 0; i < len(*imagesPath); i++ {
		key := "image" + strconv.Itoa(i)
		db.AsyncQ().AsyncInsert(key, pnames.PUBLCATION_IMAGES_TABLE, map[string]any{
			"publication": publicationID,
			"image_path":  (*imagesPath)[i],
		})
		asyncKeys = append(asyncKeys, key)
	}
	db.AsyncQ().Wait()
	if err := database.AsyncResError(asyncKeys, db); err != nil {
		return err
	}
	return nil
}

func NewPublicationHNDL() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	tv := object.TemplateView{
		TemplatePath: "",
		View: &NewPublicationView{
			object.FormView{
				FormStruct:       NewPublicationForm{},
				NotNilFormFields: []string{"Name", "Description", "Images"},
				NilIfNotExist:    []string{},
			},
		},
	}
	tv.SkipRender()
	return tv.Call
}
