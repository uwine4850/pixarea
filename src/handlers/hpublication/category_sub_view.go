package hpublication

import (
	"strconv"

	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	"github.com/uwine4850/foozy/pkg/typeopr"
	"github.com/uwine4850/pixarea/src/cnf/pnames"
)

type CategorySubView struct {
	DB *database.Database
}

// getPublicationCategories retrieving publication categories by their identifiers.
func (v *CategorySubView) CategoriesById(categoriesId []string) ([]PublicationCategory, error) {
	categories := []PublicationCategory{}
	for i := 0; i < len(categoriesId); i++ {
		// If the category is NULL.
		if categoriesId[i] == "" {
			continue
		}
		category, err := v.DB.SyncQ().QB().Select("*", pnames.CATEGORIES_TABLE).Where("id", "=", categoriesId[i], "LIMIT 1").Ex()
		if err != nil {
			return nil, err
		}
		if err := dbutils.DatabaseResultNotEmpty(category); err != nil {
			return nil, err
		}
		var categoryStruct PublicationCategory
		if err := dbmapper.FillStructFromDb(category[0], typeopr.Ptr{}.New(&categoryStruct)); err != nil {
			return nil, err
		}
		categories = append(categories, categoryStruct)
	}
	return categories, nil
}

func (v *CategorySubView) CategoriesByName(categoriesName []string) ([]string, error) {
	categories := make([]string, 2)
	for i := 0; i < len(categoriesName); i++ {
		categoryID, err := v.DB.SyncQ().Select([]string{"id"}, pnames.CATEGORIES_TABLE,
			dbutils.WHEquals(dbutils.WHValue{"name": categoriesName[i]}, ""), 1)
		if err != nil {
			return nil, err
		}
		if len(categoryID) != 1 {
			return nil, nil
		}
		id, err := dbutils.ParseInt(categoryID[0]["id"])
		if err != nil {
			return nil, err
		}
		categories[i] = strconv.Itoa(id)
	}
	return categories, nil
}
