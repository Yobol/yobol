package auth

import "github.com/gin-gonic/gin"

func MountApiGroup(router *gin.RouterGroup) error {
	router = router.Group("auth")
	router.POST("login", group.Login)   // 用户登录
	router.POST("logout", group.Logout) // 用户登出
	router.GET("whoami", group.WhoAmI)  // 当前用户身份获取
	return nil
}
