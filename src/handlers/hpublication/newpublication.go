package hpublication

import (
	"fmt"
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/object"
)

func PublicationViewHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	manager.Render().SetTemplatePath("src/templates/publication/publication_view.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	return func() {}
}

func NewPublicationPageHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	router.CatchRedirectError(r, manager)
	manager.Render().SetTemplatePath("src/templates/publication/new_publication.html")
	if err := manager.Render().RenderTemplate(w, r); err != nil {
		return func() { router.ServerError(w, err.Error(), manager) }
	}
	return func() {}
}

type NewPublicationForm struct {
	Name         []string        `form:"name"`
	Description  []string        `form:"description"`
	CtgrPIXELART []string        `form:"pub_ctgr_PIXELART"`
	CtgrGAMEDEV  []string        `form:"pub_ctgr_GAMEDEV"`
	Images       []form.FormFile `form:"images"`
}

type NewPublicationView struct {
	object.FormView
}

func (v *NewPublicationView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w)
}

func (v *NewPublicationView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.ObjectContext, error) {
	context, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT_CONTEXT)
	form := context.(object.ObjectContext)[namelib.OBJECT_CONTEXT_FORM].(NewPublicationForm)
	for i := 0; i < len(form.Images); i++ {
		fmt.Println(form.Images[i].Header.Filename)
	}
	router.SendJson(map[string]string{"success": "true"}, w)
	return object.ObjectContext{}, nil
}

func NewPublicationHNDL() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	tv := object.TemplateView{
		TemplatePath: "",
		View: &NewPublicationView{
			object.FormView{
				FormStruct:       NewPublicationForm{},
				NotNilFormFields: []string{"Name", "Description", "Images"},
				NilIfNotExist:    []string{"pub_ctgr_PIXELART", "pub_ctgr_GAMEDEV"},
			},
		},
	}
	tv.SkipRender()
	return tv.Call
}
