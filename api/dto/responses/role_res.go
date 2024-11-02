package responses

import (
	"github.com/google/uuid"
)

type RoleRes struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
