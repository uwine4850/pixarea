package handlers

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
)

func LoginHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	manager.Render().SetTemplatePath("src/templates/auth/login.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager.Config()) }
	}
	return func() {}
}

func RegisterHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	manager.Render().SetTemplatePath("src/templates/auth/register.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager.Config()) }
	}
	return func() {}
}
