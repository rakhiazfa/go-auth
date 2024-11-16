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

type RoleHandler struct {
	validator   *utils.Validator
	roleService *services.RoleService
}

func NewRoleHandler(validator *utils.Validator, roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{validator, roleService}
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	paginator := utils.NewPaginator(c)
	paginator.SetSortableFields([]string{"name"})

	roles := h.roleService.GetAll(&paginator)

	var res responses.PaginationRes[responses.RoleRes]
	utils.CatchError(copier.Copy(&res.Items, &roles))
	utils.CatchError(copier.Copy(&res.Meta, &paginator))

	c.JSON(http.StatusOK, res)
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req requests.CreateRoleReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	role := h.roleService.Create(req)

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Successfully created role",
		"createdId": role.ID,
	})
}

func (h *RoleHandler) GetById(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	role := h.roleService.GetById(id)

	var res responses.RoleRes
	utils.CatchError(copier.Copy(&res, &role))

	c.JSON(http.StatusOK, gin.H{
		"role": res,
	})
}

func (h *RoleHandler) Update(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	var req requests.UpdateRoleReq

	utils.CatchError(c.ShouldBind(&req))
	utils.CatchError(h.validator.Validate(req))

	h.roleService.Update(req, id)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Successfully updated role",
		"updatedId": id,
	})
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	h.roleService.Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Successfully deleted role",
		"updatedId": id,
	})
}

func (h *RoleHandler) Restore(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	h.roleService.Restore(id)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Successfully restored role",
		"restoredId": id,
	})
}
