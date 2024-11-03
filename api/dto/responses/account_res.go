package responses

import (
	"github.com/google/uuid"
)

type AccountRes struct {
	ID             uuid.UUID `json:"id"`
	ProfilePicture *string   `json:"profile_picture"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Roles          []RoleRes `json:"roles"`
}
