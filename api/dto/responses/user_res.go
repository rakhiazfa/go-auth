package responses

import (
	"github.com/google/uuid"
	"time"
)

type UserRes struct {
	ID             uuid.UUID `json:"id"`
	ProfilePicture *string   `json:"profile_picture"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
