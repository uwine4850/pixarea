package hprofile

import (
	"fmt"
	"net/http"

	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
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

// GetUserByAuthId In the database, the publication is bound to the auth table.
// This function gets information about the user based on his auth ID.
func GetUserByAuthId(db *database.Database, authId any) (User, error) {
	authUser, err := auth.UserByID(db, authId)
	if err != nil {
		return User{}, err
	}
	if authUser == nil {
		return User{}, fmt.Errorf("auth by id %s not found", authId)
	}
	var auth auth.AuthItem
	if err := dbutils.FillStructFromDb(authUser, &auth); err != nil {
		return User{}, err
	}

	user, err := db.SyncQ().QB().Select("*", "user").Where("auth", "=", authId, "LIMIT 1").Ex()
	if err != nil {
		return User{}, err
	}
	if len(user) != 1 {
		return User{}, fmt.Errorf("user by auth id %s not found", authId)
	}
	var userStruct User
	if err := dbutils.FillStructFromDb(user[0], &userStruct); err != nil {
		return User{}, err
	}
	userStruct.Auth = auth
	return userStruct, nil
}
