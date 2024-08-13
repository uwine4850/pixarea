package hpublication

import (
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
)

type LikeSubView struct {
	PublicationId any
	DB            *database.Database
}

func (v *LikeSubView) GetLikeCount() (int, error) {
	count, err := v.DB.SyncQ().QB().Select("*", pnames.PUBLICATION_LIKES_TABLE).Where("publication", "=", v.PublicationId).Count().Ex()
	if err != nil {
		return 0, err
	}
	if err := dbutils.DatabaseResultNotEmpty(count); err != nil {
		return 0, err
	}
	likes, err := dbutils.ParseInt(count[0]["count"])
	if err != nil {
		return 0, err
	}
	return likes, nil
}

func (v *LikeSubView) LikeExist(authId string) (bool, error) {
	like, err := v.DB.SyncQ().Select([]string{"*"}, pnames.PUBLICATION_LIKES_TABLE, dbutils.WHEquals(
		dbutils.WHValue{"publication": v.PublicationId, "auth_id": authId}, "AND",
	), 1)
	if err != nil {
		return false, err
	}
	if len(like) == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
