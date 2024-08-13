package hpublication

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
	"github.com/uwine4850/pixarea/src/utils"
)

type CommentForm struct {
	ReplyId       []string `form:"reply_id"`
	PublicationId []string `form:"publication_id"`
	CommentText   []string `form:"comment_text"`
}

func PublicationCommentHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
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
	if err := frm.ValidateCsrfToken(); err != nil {
		return utils.SuccessJsonError(w, err)
	}
	comm, err := getComment(frm)
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}

	user, err := getUser(r, manager, db)
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}
	commDb := Comment{
		ReplyId:       comm.ReplyId[0],
		PublicationId: comm.PublicationId[0],
		AuthorId:      user.AuthId,
		Text:          comm.CommentText[0],
		IsHide:        "0",
	}
	commId, err := saveComment(&commDb, db)
	if err != nil {
		return utils.SuccessJsonError(w, err)
	}

	return func() {
		router.SendJson(map[string]interface{}{
			"success": "true", "text": commDb.Text, "name": user.Name, "avatar": user.Avatar, "comm_id": commId,
		}, w)
	}
}

func getComment(frm *form.Form) (CommentForm, error) {
	if err := frm.Parse(); err != nil {
		return CommentForm{}, err
	}
	var comm CommentForm
	commPtr := typeopr.Ptr{}.New(&comm)
	if err := formmapper.FillStructFromForm(frm, commPtr, []string{}); err != nil {
		return CommentForm{}, err
	}
	fieldNames, err := formmapper.FieldsName(commPtr, []string{"ReplyId"})
	if err != nil {
		return CommentForm{}, err
	}
	if err := formmapper.FieldsNotEmpty(commPtr, fieldNames); err != nil {
		return CommentForm{}, err
	}
	return comm, nil
}

func getUser(r *http.Request, manager interfaces.IManager, db *database.Database) (hprofile.User, error) {
	currentAuth, err := hprofile.GetCurrentAuth(r, manager)
	if err != nil {
		return hprofile.User{}, err
	}
	user, err := hprofile.GetUserByAuthId(db, currentAuth.UID)
	if err != nil {
		return hprofile.User{}, err
	}
	return user, nil
}

func saveComment(commentDb *Comment, db *database.Database) (any, error) {
	insertParams, err := dbmapper.ParamsValueFromStruct(typeopr.Ptr{}.New(commentDb), []string{"id", "reply_id", "is_hide"})
	if err != nil {
		return nil, err
	}
	res, err := db.SyncQ().Insert(pnames.PUBLICATION_COMMENTS_TABLE, insertParams)
	if err != nil {
		return nil, err
	}
	return res["id"], nil
}
