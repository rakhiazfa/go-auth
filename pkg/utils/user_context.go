package utils

import (
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"gorm.io/gorm"
)

type UserContext struct {
	db          *gorm.DB
	accessToken *string
	authUser    *entities.User
}

func NewUserContext(db *gorm.DB) *UserContext {
	return &UserContext{db, nil, nil}
}

func (uc *UserContext) GetAccessToken() string {
	if uc.accessToken != nil {
		return *uc.accessToken
	}

	return ""
}

func (uc *UserContext) SetAccessToken(accessToken string) {
	uc.accessToken = &accessToken
}

func (uc *UserContext) GetAuthUser() entities.User {
	if uc.authUser != nil {
		return *uc.authUser
	}

	return entities.User{}
}

func (uc *UserContext) SetAuthUser(user *entities.User) {
	uc.authUser = user
}
