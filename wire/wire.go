//go:build wireinject

// + build:wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/rakhiazfa/vust-identity-service/api/handlers"
	"github.com/rakhiazfa/vust-identity-service/api/routes"
	"github.com/rakhiazfa/vust-identity-service/internal/database"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/internal/services"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
)

var permissionModule = wire.NewSet(
	repositories.NewPermissionRepository,
	services.NewPermissionService,
	handlers.NewPermissionHandler,
)

var roleModule = wire.NewSet(
	repositories.NewRoleRepository,
	services.NewRoleService,
	handlers.NewRoleHandler,
)

var userModule = wire.NewSet(
	repositories.NewUserRepository,
	services.NewUserService,
	handlers.NewUserHandler,
)

var accountModule = wire.NewSet(
	handlers.NewAccountHandler,
)

var authModule = wire.NewSet(
	repositories.NewUserSessionRepository,
	services.NewAuthService,
	handlers.NewAuthHandler,
)

func SetupProviders() *gin.Engine {
	wire.Build(
		database.NewPostgresConnection,
		utils.NewUserContext,
		utils.NewValidator,
		permissionModule,
		roleModule,
		userModule,
		accountModule,
		authModule,
		routes.SetupRoutes,
	)

	return nil
}
