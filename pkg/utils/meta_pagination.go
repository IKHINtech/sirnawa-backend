package utils

import (
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty"`
	Page       int    `json:"page,omitempty"`
	Sort       string `json:"sort,omitempty"`
	SortBy     string `json:"sort_by,omitempty"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	switch p.Sort {
	case "ASC":
		p.Sort = p.SortBy + " ASC"
	case "DESC":
		p.Sort = p.SortBy + " DESC"
	}
	return p.Sort
}

func (p *Pagination) GetSortBy() string {
	if p.SortBy == "" {
		p.SortBy = "created_at"
	}
	return p.SortBy
}

func Paginate(value any, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func GetPaginationParams(pageStr, perPageStr string) (int, int) {
	const defaultPage = 1
	const defaultPerPage = 10

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		perPage = defaultPerPage
	}

	return page, perPage
}
