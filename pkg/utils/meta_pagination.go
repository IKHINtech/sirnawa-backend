package utils

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty" query:"limit"`
	Page       int    `json:"page,omitempty" query:"page"`
	Sort       string `json:"sort,omitempty" query:"sort"`
	SortBy     string `json:"sort_by,omitempty" query:"sort_by"`
	TotalRows  int64  `json:"total_rows" query:"total_rows"`
	TotalPages int    `json:"total_pages" query:"total_pages"`
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

func GetPaginationParams(ctx *fiber.Ctx) Pagination {
	orderBy := ctx.Query("order_by", "created_at")

	// Mengambil nilai parameter order dari query
	order := ctx.Query("order", "DESC")

	page := ctx.QueryInt("page", 1)
	pageSize := ctx.QueryInt("page_size", 10)

	paginate := Pagination{
		Limit:  pageSize,
		Page:   page,
		Sort:   order,
		SortBy: orderBy,
	}
	return paginate
}
