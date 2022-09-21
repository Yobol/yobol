package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yobol/yobol/api/v1/article"
	"github.com/yobol/yobol/api/v1/auth"
	"github.com/yobol/yobol/api/v1/user"
	"github.com/yobol/yobol/config"
	httputils "github.com/yobol/yobol/utils/http"
)

const (
	ApiRoot = "/api/v1"
)

var (
	ErrLackAuthToken = fmt.Errorf("lack of auth token")

	unauthorizedUris = []string{
		ApiRoot + "/auth/login",
		ApiRoot + "/auth/logout",
	}

	handlers = []gin.HandlerFunc{
		AuthHandler(),
	}
)

func MountApiGroups(router *gin.Engine) error {
	group := router.Group(ApiRoot)
	for _, handler := range handlers {
		group.Use(handler)
	}

	if err := auth.MountApiGroup(group); err != nil {
		return err
	}
	if err := user.MountApiGroup(group); err != nil {
		return err
	}
	if err := article.MountApiGroup(group); err != nil {
		return err
	}
	return nil
}

// AuthHandler is a gin middleware for checking authority.
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// pre-handle request
		for _, unauthorizedUri := range unauthorizedUris {
			if strings.HasPrefix(c.Request.URL.Path, unauthorizedUri) {
				c.Next()
				return
			}
		}

		token, err := c.Cookie(config.C.Server.Auth.TokenName)
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httputils.Error{
				Code:    http.StatusUnauthorized,
				Message: ErrLackAuthToken.Error(),
			})
			return
		}

		claims, err := httputils.NewJWT().ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httputils.Error{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}
		c.Set(httputils.ClaimsKey, claims)

		// in-handle request
		c.Next()

		// post-handle request
	}
}
