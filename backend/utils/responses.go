/*
This is where every function would use as a response lies
*/
package utils

import "gorm.io/gorm"

type (
	Response struct {
		StatusCode int         `json:"status"` // make sure u use the const from the fiber for the status code (fiber already provided it, so no need to think about it)
		Message    string      `json:"message"`
		Data       any         `json:"data,omitempty"`
		Error      any         `json:"error,omitempty"`
		Meta       *Pagination `json:"meta,omitempty"`
	}

	Pagination struct {
		Total     *int
		Page      *int
		Limit     *int
		TotalPage *int
	}
)

func NewPagination(Total, Page, Limit, TotalPage *int) *Pagination {
	return &Pagination{Total, Page, Limit, TotalPage}
}

func (p *Pagination) PaginationQuery(db *gorm.DB) *gorm.DB {
	offset := (*p.Page - 1) * *p.Limit
	return db.Offset(offset).Limit(*p.Limit)
}

func GetTotalData(tableName string, db *gorm.DB, filter *func(*gorm.DB) *gorm.DB) int64 {
	var count int64

	query := db.Table(tableName)

	if filter != nil {
		if errGetTotalData := query.Scopes(*filter).Count(&count).Error; errGetTotalData != nil {
			return 0
		}
	}

	if errGetTotalData := query.Count(&count).Error; errGetTotalData != nil {
		return 0
	}
	return count
}
