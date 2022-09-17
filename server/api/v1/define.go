package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yobol/yobol/api/v1/auth"
	"github.com/yobol/yobol/api/v1/user"
)

func MountApiGroups(router *gin.Engine) error {
	group := router.Group("/api/v1")
	if err := auth.MountApiGroup(group); err != nil {
		return err
	}
	if err := user.MountApiGroup(group); err != nil {
		return err
	}
	return nil
}
