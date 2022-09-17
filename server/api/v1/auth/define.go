package auth

import "github.com/gin-gonic/gin"

type Group struct{}

func (g *Group) Login(c *gin.Context) {

}

func (g *Group) Logout(c *gin.Context) {

}

var group = new(Group)
