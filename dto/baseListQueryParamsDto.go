package dto

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseListQueryParamsDto struct {
	SortOrder *string
	SortBy    *string
	Search    *string
	Page      *int
	PerPage   *int
}

func (d *BaseListQueryParamsDto) GetSortOrderBool() bool {
	if d.SortOrder == nil {
		return true
	}
	return *d.SortOrder == "desc"
}

func (d *BaseListQueryParamsDto) GetOffset() *int {
	if d.Page == nil || d.PerPage == nil {
		return nil
	}
	offset := (*d.Page - 1) * *d.PerPage

	return &offset
}

func (d *BaseListQueryParamsDto) GetSortByWithDefaultId(sortBy *string) string {
	defaultSortBy := "id"

	if sortBy != nil {
		defaultSortBy = *sortBy
	}

	return defaultSortBy
}

// Turn Query params url to query for gorm
// This usualy use to query implement paginate
func (d *BaseListQueryParamsDto) GetQueryParamsToDbQuery(query *gorm.DB, totalRecord *int64) *gorm.DB {
	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: d.GetSortByWithDefaultId(d.SortBy),
		},
		Desc: d.GetSortOrderBool(),
	})

	query.Count(totalRecord)

	if offset := d.GetOffset(); offset != nil {
		query = query.Limit(*d.PerPage).Offset(*offset)
	}

	return query
}
