package hpublication

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
)

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
