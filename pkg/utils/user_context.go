package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"gorm.io/gorm"
)

type UserContext struct {
	db *gorm.DB
}

func NewUserContext(db *gorm.DB) *UserContext {
	return &UserContext{db}
}

func (uc *UserContext) GetAuthUser(c *gin.Context) (user entities.User) {
	userId, exists := c.Get("userId")
	if !exists {
		CatchError(fmt.Errorf("failed to get user from request context"))
	}

	uc.db.Model(&entities.User{}).Preload("Roles").First(&user, "id = ?", userId)

	return user
}

func (uc *UserContext) GetAccessToken(c *gin.Context) string {
	accessToken, exists := c.Get("accessToken")
	if !exists {
		CatchError(fmt.Errorf("failed to get access token from request context"))
	}

	return accessToken.(string)
}
