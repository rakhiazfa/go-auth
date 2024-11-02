package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/vust-identity-service/api/dto/responses"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
)

type AccountHandler struct {
	userContext *utils.UserContext
}

func NewAccountHandler(userContext *utils.UserContext) *AccountHandler {
	return &AccountHandler{userContext}
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	user := h.userContext.GetAuthUser(c)

	var account responses.AccountRes
	utils.CatchError(copier.Copy(&account, &user))

	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}
