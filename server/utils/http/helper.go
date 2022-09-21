package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ClaimsKey = "claims"
)

type Empty struct{}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func FailureResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Error{
		Code:    code,
		Message: message,
	})
}
