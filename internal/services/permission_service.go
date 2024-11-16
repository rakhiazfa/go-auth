package services

import (
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
)

type PermissionService struct {
	permissionRepository *repositories.PermissionRepository
}

func NewPermissionService(permissionRepository *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{permissionRepository}
}

func (s *PermissionService) GetAll(paginator *utils.Paginator) []entities.Permission {
	return s.permissionRepository.GetAll(paginator)
}

func (s *PermissionService) GetById(id uuid.UUID) entities.Permission {
	permission := s.permissionRepository.GetById(id)

	if permission.ID == uuid.Nil {
		utils.CatchError(utils.NewHttpError(http.StatusNotFound, "Permission not found", nil))
	}

	return permission
}
