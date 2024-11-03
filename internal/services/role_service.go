package services

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/vust-identity-service/api/dto/requests"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type RoleService struct {
	db             *gorm.DB
	roleRepository *repositories.RoleRepository
}

func NewRoleService(db *gorm.DB, roleRepository *repositories.RoleRepository) *RoleService {
	return &RoleService{db, roleRepository}
}

func (s *RoleService) GetAll(paginator *utils.Paginator) []entities.Role {
	return s.roleRepository.GetAll(paginator)
}

func (s *RoleService) Create(req requests.CreateRoleReq) entities.Role {
	var role entities.Role

	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := copier.Copy(&role, &req)
		if err != nil {
			return err
		}

		if s.roleRepository.GetByNameUnscoped(req.Name).ID != uuid.Nil {
			return utils.NewHttpError(http.StatusConflict, "Name already exists", nil)
		}

		return s.roleRepository.Create(tx, &role)
	})
	utils.CatchError(err)

	return role
}

func (s *RoleService) GetById(id uuid.UUID) entities.Role {
	role := s.roleRepository.GetById(id)

	if role.ID == uuid.Nil {
		utils.CatchError(utils.NewHttpError(http.StatusNotFound, "Role not found", nil))
	}

	return role
}

func (s *RoleService) Update(req requests.UpdateRoleReq, id uuid.UUID) entities.Role {
	var role entities.Role

	err := s.db.Transaction(func(tx *gorm.DB) error {
		role = s.GetById(id)

		if s.roleRepository.GetByNameUnscoped(req.Name, uuid.UUIDs{id}).ID != uuid.Nil {
			return utils.NewHttpError(http.StatusConflict, "Role name already exists", nil)
		}

		role.Name = req.Name

		return s.roleRepository.Update(tx, &role)
	})
	utils.CatchError(err)

	return role
}

func (s *RoleService) Delete(id uuid.UUID) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.roleRepository.Delete(tx, id)
	})
	utils.CatchError(err)
}

func (s *RoleService) Restore(id uuid.UUID) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.roleRepository.Restore(tx, id)
	})
	utils.CatchError(err)
}
