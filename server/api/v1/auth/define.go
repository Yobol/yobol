package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/yobol/yobol/utils/http"
)

var group = new(Group)

type Group struct{}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Id          v1-auth-login
// @Summary     Auth Login
// @Description login using username & password
// @Tags        auth
// @Router      /api/v1/auth/login [post]
// @Param       request body LoginReq true "request boy for login"
// @Accept      application/json
// @Produce     application/json
// @Header      all {string} YAT "yobol auth token"
// @Success     200 {object} httputils.EmptyResponse
// @Failure     401 {object} httputils.Error
// @Failure     500 {object} httputils.Error
func (g *Group) Login(c *gin.Context) {

}

func (g *Group) Logout(c *gin.Context) {

}
