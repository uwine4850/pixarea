package hpublication

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
	"github.com/uwine4850/pixarea/src/utils"
)

func PublicationLikeHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w)
		}
	}()
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	publicationId := frm.GetMultipartForm().Value["publication-id"][0]
	currentAuth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}
	likeSubView := LikeSubView{
		PublicationId: publicationId,
		DB:            db,
	}
	isLike, err := likeSubView.LikeExist(currentAuth.UID)
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}
	if isLike {
		if err := removeLike(db, publicationId, currentAuth.UID); err != nil {
			return utils.SuccessJsonError(w, err)
		}
		return func() { router.SendJson(map[string]string{"success": "true", "addLike": "false"}, w) }
	} else {
		if err := addLike(db, publicationId, currentAuth.UID); err != nil {
			return utils.SuccessJsonError(w, err)
		}
		return func() { router.SendJson(map[string]string{"success": "true", "addLike": "true"}, w) }
	}
}

func addLike(db *database.Database, publicationId string, authId string) error {
	_, err := db.SyncQ().Insert(pnames.PUBLICATION_LIKES_TABLE, map[string]any{"publication": publicationId, "auth_id": authId})
	if err != nil {
		return err
	}
	return nil
}

func removeLike(db *database.Database, publicationId string, authId string) error {
	_, err := db.SyncQ().Delete(pnames.PUBLICATION_LIKES_TABLE, dbutils.WHEquals(
		dbutils.WHValue{"publication": publicationId, "auth_id": authId}, "AND",
	))
	if err != nil {
		return err
	}
	return nil
}
