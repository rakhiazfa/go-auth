package repositories

import (
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/database/scopes"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) GetAll(paginator *utils.Paginator) (roles []entities.Role) {
	err := r.db.Scopes(scopes.Paginate(&entities.Role{}, r.db, paginator)).Find(&roles).Error
	utils.CatchError(err)

	return
}

func (r *RoleRepository) Create(tx *gorm.DB, role *entities.Role) error {
	return tx.Model(&entities.Role{}).Create(role).Error
}

func (r *RoleRepository) GetById(id uuid.UUID) (role entities.Role) {
	r.db.Model(&entities.Role{}).First(&role, "id = ?", id)

	return
}

func (r *RoleRepository) GetByNameUnscoped(name string, exclude ...uuid.UUIDs) (role entities.Role) {
	q := r.db.Model(&entities.Role{}).Unscoped().Where("name = ?", name)

	if len(exclude) > 0 {
		q = q.Not("id IN ?", exclude[0])
	}

	q.First(&role)

	return
}

func (r *RoleRepository) GetByName(name string) (role entities.Role) {
	r.db.Model(&entities.Role{}).First(&role, "name = ?", name)

	return
}

func (r *RoleRepository) Update(tx *gorm.DB, role *entities.Role) error {
	return tx.Save(role).Error
}

func (r *RoleRepository) Delete(tx *gorm.DB, id uuid.UUID) error {
	result := tx.Delete(&entities.Role{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return utils.NewHttpError(http.StatusNotFound, "Role not found", nil)
	}

	return nil
}

func (r *RoleRepository) Restore(tx *gorm.DB, id uuid.UUID) error {
	var role entities.Role

	tx.Model(&entities.Role{}).Unscoped().First(&role, "id = ? AND deleted_at IS NOT NULL", id)

	if role.ID == uuid.Nil {
		return utils.NewHttpError(http.StatusNotFound, "Role not found", nil)
	}

	role.DeletedAt = gorm.DeletedAt{}

	return tx.Save(&role).Error
}
