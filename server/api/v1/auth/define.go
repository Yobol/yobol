package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/yobol/yobol/dao/model"
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
// @Description login with username & password
// @Tags        auth
// @Router      /api/v1/auth/login [post]
// @Param       request body LoginReq true "request boy for login"
// @Accept      application/json
// @Produce     application/json
// @Header      all {string} YAT "yobol auth token"
// @Success     200 {object} httputils.Empty
// @Failure     401 {object} httputils.Error
// @Failure     500 {object} httputils.Error
func (g *Group) Login(c *gin.Context) {

}

// @Id          v1-auth-logout
// @Summary     Auth Logout
// @Description logout to clear auth cookie YAT
// @Tags        auth
// @Router      /api/v1/auth/logout [post]
// @Param       YAT      header string          false "yobol auth token"
// @Param       request  body   httputils.Empty true  "empty request body"
// @Accept      application/json
// @Produce     application/json
// @Success     200 {object} httputils.Empty
// @Failure     500 {object} httputils.Error
func (g *Group) Logout(c *gin.Context) {

}

// @Id          v1-auth-whoami
// @Summary     Who Am I
// @Description get the profile of the currently logged in user
// @Tags        auth
// @Router      /api/v1/auth/whoami [get]
// @Param       YAT header string false "yobol auth token"
// @Produce     application/json
// @Success     200 {object} model.User
// @Failure     401 {object} httputils.Error
// @Failure     500 {object} httputils.Error
func (g *Group) WhoAmI(c *gin.Context) {

}
