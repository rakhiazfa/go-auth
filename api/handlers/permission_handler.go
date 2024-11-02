package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/vust-identity-service/api/dto/responses"
	"github.com/rakhiazfa/vust-identity-service/internal/services"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
)

type PermissionHandler struct {
	permissionService *services.PermissionService
}

func NewPermissionHandler(permissionService *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService}
}

func (h *PermissionHandler) GetAll(c *gin.Context) {
	paginator := utils.NewPaginator(c)
	paginator.SetSortableFields([]string{"name", "service_key", "method", "path"})

	permissions := h.permissionService.GetAll(&paginator)

	var res responses.PaginationRes[responses.PermissionRes]
	utils.CatchError(copier.Copy(&res.Items, &permissions))
	utils.CatchError(copier.Copy(&res.Meta, &paginator))

	c.JSON(http.StatusOK, res)
}

func (h *PermissionHandler) GetById(c *gin.Context) {
	id := utils.ParseUUID(c.Param("id"))

	permission := h.permissionService.GetById(id)

	var res responses.PermissionRes
	utils.CatchError(copier.Copy(&res, &permission))

	c.JSON(http.StatusOK, gin.H{
		"permission": res,
	})
}
