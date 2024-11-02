package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
)

func setupRoleRoutes(r *gin.RouterGroup, handler *handlers.RoleHandler) {
	route := r.Group("/roles")

	route.GET("", handler.GetAll)
	route.POST("", handler.Create)
	route.GET("/:id", handler.GetById)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
	route.POST("/:id/restore", handler.Restore)
}
