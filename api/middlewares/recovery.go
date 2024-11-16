package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				switch err := r.(type) {
				case *utils.HttpError:
					handleHttpError(c, err)
					return
				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"message": "Internal Server Error",
						"error":   err.(error).Error(),
					})
				}
			}
		}()

		c.Next()
	}
}

func handleHttpError(c *gin.Context, err *utils.HttpError) {
	var validationErrors validator.ValidationErrors

	if errors.As(err.Reason, &validationErrors) {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"message": err.Message,
			"errors":  utils.FormatValidationErrors(validationErrors),
		})
	} else {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
	}
}
