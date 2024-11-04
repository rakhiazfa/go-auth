package services

import (
	"github.com/rakhiazfa/vust-identity-service/api/dto/requests"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AccountService struct {
	db             *gorm.DB
	fileService    *FileService
	userRepository *repositories.UserRepository
}

func NewAccountService(
	db *gorm.DB,
	fileService *FileService,
	userRepository *repositories.UserRepository,
) *AccountService {
	return &AccountService{db, fileService, userRepository}
}

func (s *AccountService) SetProfilePicture(accessToken string, user entities.User, req requests.SetProfilePictureReq) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		file, err := s.fileService.UploadFile(requests.UploadFileReq{
			AccessToken: accessToken,
			ServiceKey:  viper.GetString("application.key"),
			BucketName:  "vust",
			Directory:   "/users/" + user.ID.String() + "/profile-pictures",
			File:        req.ProfilePicture,
		})
		if err != nil {
			return err
		}

		user.ProfilePicture = &file.ID

		err = s.userRepository.Update(tx, &user)

		return err
	})
	utils.CatchError(err)
}
