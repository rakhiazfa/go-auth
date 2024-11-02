package repositories

import (
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/database/scopes"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll(paginator *utils.Paginator) (users []entities.User) {
	err := r.DB.Scopes(scopes.Paginate(&entities.User{}, r.DB, paginator)).Find(&users).Error
	utils.CatchError(err)

	return
}

func (r *UserRepository) Create(tx *gorm.DB, user *entities.User) error {
	return tx.Model(&entities.User{}).Create(user).Error
}

func (r *UserRepository) GetById(id uuid.UUID) (user entities.User) {
	r.DB.Model(&entities.User{}).First(&user, "id = ?", id)

	return
}

func (r *UserRepository) GetByUsernameUnscoped(username string, exclude ...uuid.UUIDs) (user entities.User) {
	q := r.DB.Model(&entities.User{}).Unscoped().Where("username = ?", username)

	if len(exclude) > 0 {
		q = q.Not("id IN ?", exclude[0])
	}

	q.First(&user)

	return
}

func (r *UserRepository) GetByEmailUnscoped(email string, exclude ...uuid.UUIDs) (user entities.User) {
	q := r.DB.Model(&entities.User{}).Unscoped().Where("email = ?", email)

	if len(exclude) > 0 {
		q = q.Not("id IN ?", exclude[0])
	}

	q.First(&user)

	return
}

func (r *UserRepository) GetByUsernameOrEmail(usernameOrEmail string) (user entities.User) {
	r.DB.Model(&entities.User{}).Preload("Roles").First(&user, "username = ? OR email = ?", usernameOrEmail, usernameOrEmail)

	return
}

func (r *UserRepository) Update(tx *gorm.DB, user *entities.User) error {
	return tx.Save(user).Error
}

func (r *UserRepository) Delete(tx *gorm.DB, id uuid.UUID) error {
	result := tx.Delete(&entities.User{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return utils.NewHttpError(http.StatusNotFound, "User not found", nil)
	}

	return nil
}

func (r *UserRepository) Restore(tx *gorm.DB, id uuid.UUID) error {
	var user entities.User

	tx.Model(&entities.User{}).Unscoped().First(&user, "id = ? AND deleted_at IS NOT NULL", id)

	if user.ID == uuid.Nil {
		return utils.NewHttpError(http.StatusNotFound, "User not found", nil)
	}

	user.DeletedAt = gorm.DeletedAt{}

	return tx.Save(&user).Error
}
