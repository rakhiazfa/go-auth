package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
)

func setupPermissionRoutes(r *gin.RouterGroup, handler *handlers.PermissionHandler) {
	route := r.Group("/permissions")

	route.GET("", handler.GetAll)
	route.GET("/:id", handler.GetById)
}
