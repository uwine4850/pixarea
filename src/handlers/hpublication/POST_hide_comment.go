package hpublication

import (
	"errors"
	"fmt"
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

func PublicationCommentHideHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	comm_publication_id := frm.Value("comm_publication_id")
	if comm_publication_id == "" {
		return utils.SuccessJsonError(w, errors.New("publication not found"))
	}
	comm_id := frm.Value("comm_id")
	if comm_id == "" {
		return utils.SuccessJsonError(w, errors.New("comment not found"))
	}

	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w)
		}
	}()
	if err := validatePublicationAuthor(r, manager, db, comm_publication_id, comm_id); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	_, err := db.SyncQ().Query(fmt.Sprintf("UPDATE %s SET is_hide = NOT is_hide WHERE id = %s;", pnames.PUBLICATION_COMMENTS_TABLE, comm_id))
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}
	return func() { router.SendJson(map[string]string{"success": "true"}, w) }
}

func validatePublicationAuthor(r *http.Request, manager interfaces.IManager, db *database.Database, publicationId string, commId string) error {
	auth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return err
	}
	res, err := db.SyncQ().Select([]string{"id"}, pnames.PUBLICATIONS_TABLE, dbutils.WHEquals(
		dbutils.WHValue{"id": publicationId, "author": auth.UID}, "AND",
	), 1)
	if err != nil {
		return err
	}
	if err := dbutils.DatabaseResultNotEmpty(res); err != nil {
		return errors.New("error validating comment hiding")
	}
	res1, err := db.SyncQ().Select([]string{"id"}, pnames.PUBLICATION_COMMENTS_TABLE, dbutils.WHEquals(
		dbutils.WHValue{"id": commId, "publication_id": publicationId}, "AND",
	), 1)
	if err != nil {
		return err
	}
	if err := dbutils.DatabaseResultNotEmpty(res1); err != nil {
		return errors.New("error validating comment hiding")
	}
	return nil
}
