package scopes

import (
	"fmt"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"gorm.io/gorm"
	"math"
	"strings"
)

func ApplySearchCondition(q *gorm.DB, search *string) *gorm.DB {
	if search != nil && *search != "" {
		return q.Where("LOWER(search_text) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(*search)))
	}

	return q
}

func ApplyDeletedCondition(q *gorm.DB, isDeleted *bool) *gorm.DB {
	q = q.Unscoped()

	if isDeleted != nil {
		if *isDeleted {
			q = q.Where("deleted_at IS NOT NULL")
		} else {
			q = q.Where("deleted_at IS NULL")
		}
	}

	return q
}

func Paginate(value interface{}, db *gorm.DB, paginator *utils.Paginator) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	q := db.Model(value)

	q = ApplySearchCondition(q, paginator.Search)
	q = ApplyDeletedCondition(q, paginator.IsDeleted)

	utils.CatchError(q.Count(&totalRows).Error)

	paginator.TotalRows = totalRows

	return func(db *gorm.DB) *gorm.DB {
		paginator.TotalPages = int64(math.Ceil(float64(paginator.TotalRows) / float64(paginator.Size)))

		q := db.Offset(paginator.GetOffset()).Limit(paginator.Size).Order(paginator.GetOrder())

		q = ApplySearchCondition(q, paginator.Search)
		q = ApplyDeletedCondition(q, paginator.IsDeleted)

		return q
	}
}
