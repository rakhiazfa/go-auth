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

type AccountHandler struct {
	userContext    *utils.UserContext
	accountService *services.AccountService
	validator      *utils.Validator
}

func NewAccountHandler(
	userContext *utils.UserContext,
	accountService *services.AccountService,
	validator *utils.Validator,
) *AccountHandler {
	return &AccountHandler{userContext, accountService, validator}
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	user := h.userContext.GetAuthUser(c)

	var account responses.AccountRes
	utils.CatchError(copier.Copy(&account, &user))

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func (h *AccountHandler) SetProfilePicture(c *gin.Context) {
	user := h.userContext.GetAuthUser(c)
	accessToken := h.userContext.GetAccessToken(c)

	var req requests.SetProfilePictureReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	h.accountService.SetProfilePicture(accessToken, user, req)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated profile picture",
	})
}
