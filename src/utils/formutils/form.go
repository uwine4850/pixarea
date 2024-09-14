package formutils

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces/itypeopr"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/secure"
)

func ParseForm(r *http.Request, fill itypeopr.IPtr, nilIfNotExist []string, requiredFields []string) error {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return err
	}
	if err := secure.ValidateFormCsrfToken(r, frm); err != nil {
		return err
	}
	if err := formmapper.FillStructFromForm(frm, fill, nilIfNotExist); err != nil {
		return err
	}
	if err := formmapper.FieldsNotEmpty(fill, requiredFields); err != nil {
		return err
	}
	return nil
}
