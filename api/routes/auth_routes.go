package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
)

func setupAuthRoutes(r *gin.RouterGroup, handler *handlers.AuthHandler) {
	route := r.Group("/auth")

	route.POST("/sign-in", handler.SignIn)
	route.POST("/sign-up", handler.SignUp)
	route.POST("/sign-out", handler.SignOut)
	route.POST("/token", handler.RefreshToken)
}
