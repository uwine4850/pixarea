package securitymddll

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Cors(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
