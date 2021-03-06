package proxy_http

import (
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

type OauthController struct {
}

func (o *OauthController) RouterRegister(group *gin.RouterGroup) {
	group.POST("/tokens", o.Tokens)
}

func (o *OauthController) RouterGroupName() (name string) {
	return "/oauth"
}

func (o *OauthController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{
		middleware.TranslationMiddleware(),
	}
}

func (o *OauthController) Tokens(c *gin.Context) {
	exec := &TokenInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
}
