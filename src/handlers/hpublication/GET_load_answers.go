package hpublication

import (
	"net/http"
	"strconv"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
	"github.com/uwine4850/pixarea/src/utils"
)

func LoadAnswersHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	comm_id := r.URL.Query().Get("comm_id")
	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w)
		}
	}()
	comments, err := db.SyncQ().QB().Select("*", pnames.PUBLICATION_COMMENTS_TABLE).Where("reply_id", "=", comm_id, "ORDER BY id DESC").Ex()
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}
	var commentsDb []Comment
	mapper := dbmapper.NewMapper(comments, typeopr.Ptr{}.New(&commentsDb))
	if err := mapper.Fill(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	userDbKeys := []string{}
	for i := 0; i < len(commentsDb); i++ {
		key := "user" + strconv.Itoa(i)
		db.AsyncQ().QB(key).Select("*", pnames.USER_TABLE).Where("auth", "=", commentsDb[i].AuthorId).Ex()
		userDbKeys = append(userDbKeys, key)
	}
	db.AsyncQ().Wait()
	for i := 0; i < len(userDbKeys); i++ {
		res, _ := db.AsyncQ().LoadAsyncRes(userDbKeys[i])
		if res.Error != nil {
			return utils.SuccessJsonError(w, res.Error)
		}
		if err := dbutils.DatabaseResultNotEmpty(res.Res); err != nil {
			return utils.SuccessJsonError(w, err)
		}
		var user hprofile.User
		if err := dbmapper.FillStructFromDb(res.Res[0], typeopr.Ptr{}.New(&user)); err != nil {
			return utils.SuccessJsonError(w, err)
		}
		commentsDb[i].Author = user
	}

	return func() {
		router.SendJson(map[string]interface{}{"success": "true", "comments": commentsDb}, w)
	}
}
