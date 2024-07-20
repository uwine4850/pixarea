package hprofile

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router/cookies"
)

func GetCurrentAuth(r *http.Request, manager interfaces.IManager) (auth.AuthCookie, error) {
	hashKey := manager.Config().Get32BytesKey().HashKey()
	blockKey := manager.Config().Get32BytesKey().BlockKey()
	var cookieAuth auth.AuthCookie
	if err := cookies.ReadSecureCookieData([]byte(hashKey), []byte(blockKey), r, namelib.COOKIE_AUTH, &cookieAuth); err != nil {
		return auth.AuthCookie{}, err
	}
	return cookieAuth, nil
}
