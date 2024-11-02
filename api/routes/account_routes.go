package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
)

func setupAccountRoutes(r *gin.RouterGroup, handler *handlers.AccountHandler) {
	route := r.Group("/account")

	route.GET("", handler.GetAccount)
}
