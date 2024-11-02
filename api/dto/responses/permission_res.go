package responses

import (
	"github.com/google/uuid"
	"time"
)

type PermissionRes struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	ServiceKey string    `json:"service_key"`
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
