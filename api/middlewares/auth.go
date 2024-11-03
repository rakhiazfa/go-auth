package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
	"strings"
)

func RequiresAuth(userContext *utils.UserContext, userSessionRepository *repositories.UserSessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil))
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil))
		}

		accessToken := tokenParts[1]

		claims, err := utils.VerifyAccessToken(accessToken)
		if err != nil {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err))
		}

		userSession := userSessionRepository.GetByJwtIDWithUser(utils.ParseUUID((*claims)["jti"].(string)))

		if userSession.ID == uuid.Nil {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil))
		}

		userContext.SetAccessToken(accessToken)
		userContext.SetAuthUser(userSession.User)

		c.Next()
	}
}
