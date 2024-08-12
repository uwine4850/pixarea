package hpublication

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/foozy/pkg/utils/fstring"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
)

const MAX_IMAGE_BYTES_SIZE = 10_485_760 // 10MB
const MAX_SELECT_CATEGORIES = 2
const PATH_TO_PUBLICATION_DIRECTORY = "src/media/publications"

type PublicationDB struct {
	Id          string `db:"id"`
	Author      string `db:"author"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Category1   string `db:"category1"`
	Category2   string `db:"category2"`
	Date        string `db:"date"`
}

type PublicationCategory struct {
	Name string `db:"name"`
}

type NewPublicationForm struct {
	Name        []string        `form:"name"`
	Description []string        `form:"description"`
	Categories  []string        `form:"catedories"`
	Images      []form.FormFile `form:"images" ext:".png .jpg .jpeg"`
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
		if err := dbmapper.FillStructFromDb(categoriesDB[i], typeopr.Ptr{}.New(&category)); err != nil {
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
	// Getting the NewPublicationForm form.
	formInterface, err := v.FormInterface(manager.OneTimeData())
	if err != nil {
		return nil, err
	}
	publicationForm := formInterface.(NewPublicationForm)

	// Validation of form data.
	if err := imageSizeValidation(publicationForm.Images); err != nil {
		return nil, err
	}
	if err := selectCategoriesValidation(publicationForm.Categories); err != nil {
		return nil, err
	}

	// Obtaining IDs of selected categories from the database.
	// Get the current authentication ID.
	categories, err := getCategories(db, publicationForm.Categories)
	if err != nil {
		return nil, err
	}
	currentAuth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return nil, err
	}

	//Start of transaction.
	// Creating a publication in a table. Then creating images in a directory and saving them in a database.
	// If something goes wrong, roll back the changes.
	db.BeginTransaction()
	publication := PublicationDB{
		Author:      currentAuth.UID,
		Name:        publicationForm.Name[0],
		Description: publicationForm.Description[0],
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
	createdImages, err := createImages(w, publicationForm.Images, manager)
	if err != nil {
		return nil, err
	}
	if err := saveImages(db, newPublicationID, &createdImages); err != nil {
		rollBackCreateImages(createdImages)
		return nil, err
	}
	if err := db.CommitTransaction(); err != nil {
		rollBackCreateImages(createdImages)
		return nil, err
	}

	router.SendJson(map[string]string{"success": "true"}, w)
	return object.ObjectContext{}, nil
}

func imageSizeValidation(images []form.FormFile) error {
	for i := 0; i < len(images); i++ {
		if images[i].Header.Size > MAX_IMAGE_BYTES_SIZE {
			return errors.New("image size is more than 10MB")
		}
	}
	return nil
}

func selectCategoriesValidation(categories []string) error {
	if len(categories) > MAX_SELECT_CATEGORIES {
		return fmt.Errorf("you can select up to %s categories", strconv.Itoa(MAX_SELECT_CATEGORIES))
	}
	return nil
}

func createPublicationInTable(db *database.Database, publicationDb PublicationDB) (any, error) {
	publicationParams, err := dbmapper.ParamsValueFromStruct(typeopr.Ptr{}.New(&publicationDb), []string{"id", "category1", "category2"})
	if err != nil {
		return nil, err
	}
	info, err := db.SyncQ().Insert(pnames.PUBLICATIONS_TABLE, publicationParams)
	if err != nil {
		return nil, err
	}

	return info["id"], nil
}

func getCategories(db *database.Database, categoriesName []string) ([]string, error) {
	categories := make([]string, 2)
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
		if err := form.SaveFile(w, images[i].Header, PATH_TO_PUBLICATION_DIRECTORY, &path, manager); err != nil {
			return nil, err
		}
		imagesPath = append(imagesPath, path)
	}
	return imagesPath, nil
}

func rollBackCreateImages(paths []string) {
	for i := 0; i < len(paths); i++ {
		if fstring.PathExist(paths[i]) {
			os.Remove(paths[i])
		}
	}
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
				NotNilFormFields: []string{"*"},
				NilIfNotExist:    []string{},
			},
		},
	}
	tv.SkipRender()
	return tv.Call
}
