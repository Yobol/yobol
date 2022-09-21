package article

import "github.com/gin-gonic/gin"

func MountApiGroup(router *gin.RouterGroup) error {
	router = router.Group("articles")
	return nil
}
