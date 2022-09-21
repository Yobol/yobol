package httputils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yobol/yobol/config"
)

func SetAuthCookie(c *gin.Context, token string) {
	cc := config.C
	fmt.Println(cc)

	c.SetCookie(config.C.Server.Auth.TokenName, token,
		int(config.C.Server.Auth.TokenMaxAge), "/", "", false, true)
}
