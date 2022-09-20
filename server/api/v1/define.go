package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yobol/yobol/api/v1/article"
	"github.com/yobol/yobol/api/v1/auth"
	"github.com/yobol/yobol/api/v1/user"
)

var handlers = []gin.HandlerFunc{
	AuthHandler(),
}

func MountApiGroups(router *gin.Engine) error {
	group := router.Group("/api/v1")
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
	}
}
