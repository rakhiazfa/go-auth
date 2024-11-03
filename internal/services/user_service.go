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

type UserService struct {
	db             *gorm.DB
	userRepository *repositories.UserRepository
}

func NewUserService(db *gorm.DB, userRepository *repositories.UserRepository) *UserService {
	return &UserService{db, userRepository}
}

func (s *UserService) GetAll(paginator *utils.Paginator) []entities.User {
	return s.userRepository.GetAll(paginator)
}

func (s *UserService) Create(req requests.CreateUserReq) entities.User {
	var user entities.User

	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := copier.Copy(&user, &req)
		if err != nil {
			return err
		}

		if s.userRepository.GetByUsernameUnscoped(req.Username).ID != uuid.Nil {
			return utils.NewHttpError(http.StatusConflict, "Username already exists", nil)
		}
		if s.userRepository.GetByEmailUnscoped(req.Email).ID != uuid.Nil {
			return utils.NewHttpError(http.StatusConflict, "Email already exists", nil)
		}

		return s.userRepository.Create(tx, &user)
	})
	utils.CatchError(err)

	return user
}

func (s *UserService) GetById(id uuid.UUID) entities.User {
	user := s.userRepository.GetById(id)

	if user.ID == uuid.Nil {
		utils.CatchError(utils.NewHttpError(http.StatusNotFound, "User not found", nil))
	}

	return user
}

func (s *UserService) Update(req requests.UpdateUserReq, id uuid.UUID) entities.User {
	var user entities.User

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if s.userRepository.GetByUsernameUnscoped(req.Username, uuid.UUIDs{id}).ID != uuid.Nil {
			return utils.NewHttpError(http.StatusConflict, "Username already exists", nil)
		}
		if s.userRepository.GetByEmailUnscoped(req.Email, uuid.UUIDs{id}).ID != uuid.Nil {
			return utils.NewHttpError(http.StatusConflict, "Email already exists", nil)
		}

		user = s.GetById(id)

		user.Avatar = req.Avatar
		user.Name = req.Name
		user.Username = req.Username
		user.Email = req.Email

		return s.userRepository.Update(tx, &user)
	})
	utils.CatchError(err)

	return user
}

func (s *UserService) Delete(id uuid.UUID) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.userRepository.Delete(tx, id)
	})
	utils.CatchError(err)
}

func (s *UserService) Restore(id uuid.UUID) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.userRepository.Restore(tx, id)
	})
	utils.CatchError(err)
}
