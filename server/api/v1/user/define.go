package user

import "github.com/gin-gonic/gin"

type UserGroup struct{}

func (g *UserGroup) CreateUser(c *gin.Context) {}

func (g *UserGroup) DeleteUser(c *gin.Context) {}

func (g *UserGroup) UpdateUser(c *gin.Context) {}

func (g *UserGroup) ListUsers(c *gin.Context) {}

func (g *UserGroup) GetUser(c *gin.Context) {}
