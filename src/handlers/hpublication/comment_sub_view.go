package hpublication

import (
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
	"github.com/uwine4850/pixarea/src/handlers/hprofile"
)

type Comment struct {
	Id            string `db:"id"`
	ReplyId       string `db:"reply_id"`
	PublicationId string `db:"publication_id"`
	AuthorId      string `db:"author_id"`
	Text          string `db:"text"`
	IsHide        string `db:"is_hide"`
	Author        hprofile.User
	ReplyCount    int
}

type CommentSubView struct {
	PublicationId string
	DB            *database.Database
}

func (v *CommentSubView) CommentReplyCount(id any) (int, error) {
	count, err := v.DB.SyncQ().Count([]string{"*"}, pnames.PUBLICATION_COMMENTS_TABLE, dbutils.WHEquals(
		dbutils.WHValue{"reply_id": id}, "",
	), 0)
	if err != nil {
		return -1, err
	}
	if err := dbutils.DatabaseResultNotEmpty(count); err != nil {
		return -1, err
	}
	intCount, err := dbutils.ParseInt(count[0]["count"])
	if err != nil {
		return -1, err
	}
	return intCount, nil
}

func (v *CommentSubView) GetComments() ([]Comment, error) {
	commentsList := []Comment{}
	comments, err := v.DB.SyncQ().QB().Select("*", pnames.PUBLICATION_COMMENTS_TABLE).
		Where("publication_id", "=", v.PublicationId, "AND", "reply_id IS NULL", "ORDER BY id DESC").Ex()
	if err != nil {
		return nil, err
	}
	if err := dbutils.DatabaseResultNotEmpty(comments); err != nil {
		return []Comment{}, nil
	}
	for i := 0; i < len(comments); i++ {
		var comm Comment
		if err := dbmapper.FillStructFromDb(comments[i], typeopr.Ptr{}.New(&comm)); err != nil {
			return nil, err
		}
		user, err := hprofile.GetUserByAuthId(v.DB, comm.AuthorId)
		if err != nil {
			return nil, err
		}
		comm.Author = user
		commentReplyCount, err := v.CommentReplyCount(comm.Id)
		if err != nil {
			return nil, err
		}
		comm.ReplyCount = commentReplyCount
		commentsList = append(commentsList, comm)
	}
	return commentsList, nil
}
