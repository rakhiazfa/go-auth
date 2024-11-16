package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/internal/repositories"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
	"strings"
)

func RequiresAuth(userSessionRepository *repositories.UserSessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil))
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil))
		}

		accessToken := tokenParts[1]

		claims, err := utils.VerifyAccessToken(accessToken)
		if err != nil {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", err))
		}

		userSession := userSessionRepository.GetByJwtID(utils.ParseUUID((*claims)["jti"].(string)))

		if userSession.ID == uuid.Nil {
			utils.CatchError(utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil))
		}

		c.Set("accessToken", accessToken)
		c.Set("userId", userSession.UserId)

		c.Next()
	}
}
