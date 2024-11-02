package entities

import (
	"github.com/google/uuid"
	"time"
)

type UserSession struct {
	BaseEntity
	UserId    uuid.UUID
	User      *User     `gorm:"foreignKey:UserId;references:ID"`
	JTI       uuid.UUID `gorm:"type:uuid;unique"`
	IpAddress *string   `gorm:"type:varchar(50)"`
	UserAgent *string   `gorm:"type:varchar(255)"`
	Revoked   bool      `gorm:"type:boolean;default:false"`
	ExpAt     time.Time
}
