package responses

import (
	"github.com/google/uuid"
)

type AccountRes struct {
	ID       uuid.UUID `json:"id"`
	Avatar   *string   `json:"avatar"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Roles    []RoleRes `json:"roles"`
}
