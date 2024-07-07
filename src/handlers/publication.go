package handlers

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
)

func PublicationViewHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	manager.Render().SetTemplatePath("src/templates/publication/publication_view.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager.Config()) }
	}
	return func() {}
}

func NewPublicationHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	manager.Render().SetTemplatePath("src/templates/publication/new_publication.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager.Config()) }
	}
	return func() {}
}
