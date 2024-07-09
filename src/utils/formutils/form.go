package formutils

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/router/form"
)

func ParseForm(r *http.Request, fill *form.FillableFormStruct, nilIfNotExist []string, requiredFields []string) error {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return err
	}
	if err := frm.ValidateCsrfToken(); err != nil {
		return err
	}
	if err := form.FillStructFromForm(frm, fill, nilIfNotExist); err != nil {
		return err
	}
	if err := form.FieldsNotEmpty(fill, requiredFields); err != nil {
		return err
	}
	return nil
}
