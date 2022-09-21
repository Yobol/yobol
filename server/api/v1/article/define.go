package article

import "github.com/gin-gonic/gin"

type Group struct{}

func (g *Group) CreateArticle(c *gin.Context) {

}

func (g *Group) DeleteArticle(c *gin.Context) {

}

func (g *Group) UpdateArticle(c *gin.Context) {

}

func (g *Group) ListArticles(c *gin.Context) {

}

func (g *Group) GetArticle(c *gin.Context) {

}

var group = new(Group)
