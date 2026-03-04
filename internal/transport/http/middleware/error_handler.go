package middleware

import (
	"godest/internal/transport/http/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Written() {
			return
		}
		if len(c.Errors) == 0 {
			return
		}

		apiErr := response.NormalizeError(c.Errors.Last().Err)
		response.JSON(c, apiErr.HTTPStatus, apiErr.Code, apiErr.Message, nil)
	}
}
