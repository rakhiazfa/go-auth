package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/vust-identity-service/api/dto/requests"
	"github.com/rakhiazfa/vust-identity-service/internal/entities"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type AuthService struct {
	userRepository        *repositories.UserRepository
	userSessionRepository *repositories.UserSessionRepository
	roleRepository        *repositories.RoleRepository
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	userSessionRepository *repositories.UserSessionRepository,
	roleRepository *repositories.RoleRepository,
) *AuthService {
	return &AuthService{userRepository, userSessionRepository, roleRepository}
}

func (s *AuthService) SignUp(req requests.SignUpReq) entities.User {
	var user entities.User

	err := s.userRepository.DB.Transaction(func(tx *gorm.DB) error {
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

		customerRole := s.roleRepository.GetByName("Customer")

		if customerRole.ID == uuid.Nil {
			return utils.NewHttpError(http.StatusNotFound, "Role 'Customer' not found", nil)
		}

		user.Roles = append(user.Roles, customerRole)

		return s.userRepository.Create(tx, &user)
	})
	utils.CatchError(err)

	return user
}

func (s *AuthService) SignIn(req requests.SignInReq) (refreshToken string, accessToken string, user entities.User) {
	err := s.userSessionRepository.DB.Transaction(func(tx *gorm.DB) error {
		user = s.userRepository.GetByUsernameOrEmail(req.Username)

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); user.ID == uuid.Nil || err != nil {
			return utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
		}

		jti := uuid.New()

		refreshToken = utils.CreateRefreshToken(jwt.MapClaims{
			"sub": user.ID,
			"jti": jti,
		})
		accessToken = utils.CreateAccessToken(jwt.MapClaims{
			"sub": user.ID,
			"jti": jti,
		})

		return s.userSessionRepository.Create(tx, &entities.UserSession{
			UserId:    user.ID,
			JTI:       jti,
			IpAddress: &req.IpAddress,
			UserAgent: &req.UserAgent,
			ExpAt:     utils.ExtractJwtExp(refreshToken),
		})
	})
	utils.CatchError(err)

	return
}

func (s *AuthService) SignOut(req requests.SignOutReq) {
	err := s.userSessionRepository.DB.Transaction(func(tx *gorm.DB) error {
		claims, err := utils.VerifyRefreshToken(req.RefreshToken)
		if err != nil {
			return utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err)
		}

		jti := (*claims)["jti"].(string)

		userSession := s.userSessionRepository.GetByJwtID(utils.ParseUUID(jti))

		if userSession.ID == uuid.Nil {
			return utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
		}

		return s.userSessionRepository.Revoke(tx, &userSession)
	})
	utils.CatchError(err)
}

func (s *AuthService) RefreshToken(req requests.RefreshTokenReq) (refreshToken string, accessToken string) {
	err := s.userSessionRepository.DB.Transaction(func(tx *gorm.DB) error {
		claims, err := utils.VerifyRefreshToken(req.RefreshToken)
		if err != nil {
			return utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err)
		}

		jti := utils.ParseUUID((*claims)["jti"].(string))

		prevUserSession := s.userSessionRepository.GetByJwtID(jti)

		if prevUserSession.ID == uuid.Nil {
			return utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
		}

		err = s.userSessionRepository.Delete(tx, prevUserSession.ID)
		if err != nil {
			return utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err)
		}

		jti = uuid.New()

		refreshToken = utils.CreateRefreshToken(jwt.MapClaims{
			"sub": prevUserSession.UserId,
			"jti": jti,
		})
		accessToken = utils.CreateAccessToken(jwt.MapClaims{
			"sub": prevUserSession.UserId,
			"jti": jti,
		})

		return s.userSessionRepository.Create(tx, &entities.UserSession{
			UserId:    prevUserSession.UserId,
			JTI:       jti,
			IpAddress: &req.IpAddress,
			UserAgent: &req.UserAgent,
			ExpAt:     utils.ExtractJwtExp(refreshToken),
		})
	})
	utils.CatchError(err)

	return
}
