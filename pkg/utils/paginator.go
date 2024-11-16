package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"slices"
	"strings"
)

type Paginator struct {
	Page           int     `json:"page" form:"page"`
	Size           int     `json:"size" form:"size"`
	Sort           *string `json:"-" form:"sort"`
	Search         *string `json:"-" form:"search"`
	IsDeleted      *bool   `json:"-" form:"isDeleted"`
	TotalRows      int64   `json:"totalRows"`
	TotalPages     int64   `json:"totalPages"`
	sortableFields []string
}

func NewPaginator(c *gin.Context) Paginator {
	var paginator Paginator

	if err := c.ShouldBindQuery(&paginator); err != nil {
		CatchError(err)
	}

	if paginator.Page <= 0 {
		paginator.Page = 1
	}
	if paginator.Size <= 0 {
		paginator.Size = 10
	}

	paginator.sortableFields = append(paginator.sortableFields, "created_at", "updated_at")

	return paginator
}

func (p *Paginator) GetOffset() int {
	return (p.Page - 1) * p.Size
}

func (p *Paginator) SetSortableFields(fields []string) {
	p.sortableFields = append(p.sortableFields, fields...)
}

func (p *Paginator) GetOrder() string {
	orderBy := "created_at desc"

	if p.Sort != nil {
		sort := strings.Split(*p.Sort, ":")

		if len(sort) == 2 {
			field, direction := sort[0], strings.ToLower(sort[1])

			if slices.Contains(p.sortableFields, field) && (direction == "asc" || direction == "desc") {
				orderBy = fmt.Sprintf("%s %s", field, direction)
			}
		}
	}

	return orderBy
}
