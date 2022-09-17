package user

import "github.com/gin-gonic/gin"

type Group struct{}

func (g *Group) CreateUser(c *gin.Context) {}

func (g *Group) DeleteUser(c *gin.Context) {}

func (g *Group) UpdateUser(c *gin.Context) {}

func (g *Group) ListUsers(c *gin.Context) {}

func (g *Group) GetUser(c *gin.Context) {}

var group = new(Group)
