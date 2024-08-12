package utils

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/router"
)

func SuccessJsonError(w http.ResponseWriter, err error) func() {
	return func() { router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w) }
}
