package controller

import (
	"gateway/dto"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

func (p *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}

	out, err := params.LoginCheck(c)
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}

	middleware.ResponseSuccess(c, out)
	return
}
