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

type UserHandler struct {
	validator   *utils.Validator
	userService *services.UserService
}

func NewUserHandler(validator *utils.Validator, userService *services.UserService) *UserHandler {
	return &UserHandler{validator, userService}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	paginator := utils.NewPaginator(c)
	paginator.SetSortableFields([]string{"name", "username", "email"})

	users := h.userService.GetAll(&paginator)

	var res responses.PaginationRes[responses.UserRes]
	utils.CatchError(copier.Copy(&res.Items, &users))
	utils.CatchError(copier.Copy(&res.Meta, &paginator))

	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req requests.CreateUserReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	user := h.userService.Create(req)

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Successfully created user",
		"createdId": user.ID,
	})
}

func (h *UserHandler) GetById(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	user := h.userService.GetById(id)

	var res responses.UserRes
	utils.CatchError(copier.Copy(&res, &user))

	c.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	var req requests.UpdateUserReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	h.userService.Update(req, id)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Successfully updated user",
		"updatedId": id,
	})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	h.userService.Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Successfully deleted user",
		"updatedId": id,
	})
}

func (h *UserHandler) Restore(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	h.userService.Restore(id)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Successfully restored user",
		"updatedId": id,
	})
}
