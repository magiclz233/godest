package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeOK           = "OK"
	CodeBadRequest   = "BAD_REQUEST"
	CodeUnauthorized = "UNAUTHORIZED"
	CodeConflict     = "CONFLICT"
	CodeInternal     = "INTERNAL_ERROR"
)

type Body struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func JSON(c *gin.Context, status int, code, message string, data any) {
	c.JSON(status, Body{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Success(c *gin.Context, data any) {
	JSON(c, http.StatusOK, CodeOK, "success", data)
}

func Created(c *gin.Context, data any) {
	JSON(c, http.StatusCreated, CodeOK, "success", data)
}
