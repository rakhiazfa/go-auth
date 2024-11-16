package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-"`
}

type BaseEntityWithSoftDelete struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-"`
	DeletedAt gorm.DeletedAt
}
