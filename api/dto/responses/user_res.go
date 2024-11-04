package responses

import (
	"time"

	"github.com/google/uuid"
)

type UserRes struct {
	ID             uuid.UUID  `json:"id"`
	ProfilePicture *uuid.UUID `json:"profilePicture"`
	Name           string     `json:"name"`
	Username       string     `json:"username"`
	Email          string     `json:"email"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}
