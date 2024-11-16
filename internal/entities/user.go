package entities

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseEntityWithSoftDelete
	ProfilePicture *uuid.UUID    `gorm:"type:varchar(255)"`
	Name           string        `gorm:"type:varchar(255)"`
	Username       string        `gorm:"type:varchar(255);unique"`
	Email          string        `gorm:"type:varchar(255);unique"`
	Password       string        `gorm:"type:varchar(255)"`
	Roles          []Role        `gorm:"many2many:user_roles"`
	Sessions       []UserSession `gorm:"foreignKey:UserId;references:ID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := u.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash
	}

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") && u.Password != "" {
		hash, err := u.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash
	}

	return
}

func (u *User) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
