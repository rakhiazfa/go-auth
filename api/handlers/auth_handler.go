package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/vust-identity-service/api/dto/requests"
	"github.com/rakhiazfa/vust-identity-service/api/dto/responses"
	"github.com/rakhiazfa/vust-identity-service/internal/services"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
)

type AuthHandler struct {
	validator   *utils.Validator
	authService *services.AuthService
}

func NewAuthHandler(validator *utils.Validator, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{validator, authService}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req requests.SignUpReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	h.authService.SignUp(req)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created an account",
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req requests.SignInReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	req.IpAddress = c.ClientIP()
	req.UserAgent = c.GetHeader("User-Agent")

	refreshToken, accessToken, user := h.authService.SignIn(req)

	var account responses.AccountRes
	utils.CatchError(copier.Copy(&account, &user))

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
		"account":      account,
	})
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	var req requests.SignOutReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	h.authService.SignOut(req)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed out",
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req requests.RefreshTokenReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	req.IpAddress = c.ClientIP()
	req.UserAgent = c.GetHeader("User-Agent")

	refreshToken, accessToken := h.authService.RefreshToken(req)

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}
