package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
	"github.com/rakhiazfa/vust-identity-service/api/middlewares"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"github.com/spf13/viper"
	"net/http"
)

func SetupRoutes(
	userContext *utils.UserContext,
	userSessionRepository *repositories.UserSessionRepository,
	permissionHandler *handlers.PermissionHandler,
	roleHandler *handlers.RoleHandler,
	userHandler *handlers.UserHandler,
	accountHandler *handlers.AccountHandler,
	authHandler *handlers.AuthHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Recovery())

	publicApi := r.Group("/api")
	protectedApi := r.Group("/api")

	protectedApi.Use(middlewares.RequiresAuth(userContext, userSessionRepository))

	publicApi.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": gin.H{
				"name":    viper.GetString("application.name"),
				"version": viper.GetString("application.version"),
			},
		})
	})

	setupPermissionRoutes(protectedApi, permissionHandler)
	setupRoleRoutes(protectedApi, roleHandler)
	setupUserRoutes(protectedApi, userHandler)
	setupAccountRoutes(protectedApi, accountHandler)
	setupAuthRoutes(publicApi, authHandler)

	return r
}
