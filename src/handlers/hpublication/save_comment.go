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
)

type CommentForm struct {
	ReplyId       []string `form:"reply_id"`
	PublicationId []string `form:"publication_id"`
	CommentText   []string `form:"comment_text"`
}

// type CommentDb struct {
// 	PublicationId string `db:"publication_id"`
// 	AuthorId      string `db:"author_id"`
// 	CommentText   string `db:"text"`
// }

func PublicationCommentHNDL(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DB_ARGS)
	if err := db.Connect(); err != nil {
		return func() { router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w) }
	}
	defer func() {
		if err := db.Close(); err != nil {
			router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w)
		}
	}()

	frm := form.NewForm(r)
	if err := frm.ValidateCsrfToken(); err != nil {
		return func() { router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w) }
	}
	comm, err := getComment(frm)
	if err != nil {
		return func() { router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w) }
	}

	user, err := getUser(r, manager, db)
	if err != nil {
		return func() { router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w) }
	}
	replyId := ""
	if len(comm.ReplyId) > 0 {
		replyId = comm.ReplyId[0]
	}
	commDb := Comment{
		ReplyId:       replyId,
		PublicationId: comm.PublicationId[0],
		AuthorId:      user.AuthId,
		Text:          comm.CommentText[0],
		IsHide:        "0",
	}
	commId, err := saveComment(&commDb, db)
	if err != nil {
		return func() { router.SendJson(map[string]string{"success": "false", "error": err.Error()}, w) }
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
	fillable := form.NewFillableFormStruct(&comm)
	if err := form.FillStructFromForm(frm, fillable, []string{}); err != nil {
		return CommentForm{}, err
	}
	fieldNames, err := form.FieldsName(fillable, []string{"ReplyId"})
	if err != nil {
		return CommentForm{}, err
	}
	if err := form.FieldsNotEmpty(fillable, fieldNames); err != nil {
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
	insertParams, err := dbutils.ParamsValueFromStruct(commentDb, []string{"id", "reply_id", "is_hide"})
	if err != nil {
		return nil, err
	}
	res, err := db.SyncQ().Insert(pnames.PUBLICATION_COMMENTS_TABLE, insertParams)
	if err != nil {
		return nil, err
	}
	return res["id"], nil
}
