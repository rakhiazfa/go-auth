package repositories

import (
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type UserSessionRepository struct {
	DB *gorm.DB
}

func NewUserSessionRepository(db *gorm.DB) *UserSessionRepository {
	return &UserSessionRepository{db}
}

func (r *UserSessionRepository) Create(tx *gorm.DB, userSession *entities.UserSession) error {
	return tx.Model(&entities.UserSession{}).Create(userSession).Error
}

func (r *UserSessionRepository) GetByJwtID(jwtId uuid.UUID) (userSession entities.UserSession) {
	r.DB.Model(&entities.UserSession{}).Select(
		"id", "user_id", "jti", "revoked",
	).First(&userSession, "revoked IS NOT TRUE AND jti = ?", jwtId)

	return
}

func (r *UserSessionRepository) Revoke(tx *gorm.DB, userSession *entities.UserSession) error {
	return tx.Model(userSession).Update("revoked", true).Error
}

func (r *UserSessionRepository) Delete(tx *gorm.DB, id uuid.UUID) error {
	result := r.DB.Delete(&entities.UserSession{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return utils.NewHttpError(http.StatusNotFound, "User session not found", nil)
	}

	return nil
}
