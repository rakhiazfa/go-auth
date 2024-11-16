package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
)

func setupUserRoutes(r *gin.RouterGroup, handler *handlers.UserHandler) {
	route := r.Group("/users")

	route.GET("", handler.GetAll)
	route.POST("", handler.Create)
	route.GET("/:id", handler.GetById)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
	route.POST("/:id/restore", handler.Restore)
}
