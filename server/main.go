package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/yobol/yobol/config"
)

const (
	AuthCookieName = "YAC"
)

// General API Info
// Format Reference: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

// @title       API Docs for Yobol
// @version     1.0.0
// @descrption  API Docs for Yobol

// @contact.name   Yobol
// @contact.email  yobol1024@gmail.com
// @contact.url    yobol.me (unavailble)

// @host      host:8080
// @BasePath  /apidocs/#/

func init() {
	config.Init()
}

func main() {
	route := gin.Default()
	route.GET("/home", CheckCookie(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello",
		})
	})

	apiV1 := route.Group("/api/v1")
	apiV1.POST("/auth/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "admin" || password != "123456" { // fake
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.SetCookie(AuthCookieName, "ok", 3*24*60*60, "/", "", false, true) // YAS indicates 'Yobol Authentication Cookie'
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Success!",
		})
	})
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.Run() // listen and serve on 0.0.0.0:8080
}

func CheckCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie(AuthCookieName); err != nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
