package repositories

import (
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/database/scopes"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db}
}

func (r *PermissionRepository) GetAll(paginator *utils.Paginator) (permissions []entities.Permission) {
	err := r.db.Scopes(scopes.Paginate(&entities.Permission{}, r.db, paginator)).Find(&permissions).Error
	utils.CatchError(err)

	return
}

func (r *PermissionRepository) GetById(id uuid.UUID) (permission entities.Permission) {
	r.db.Model(&entities.Permission{}).First(&permission, "id = ?", id)

	return
}
