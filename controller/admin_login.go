package controller

import (
	"gateway/dto"
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
		ResponseError(c, 1001, err)
		return
	}

	out, err := params.LoginCheck(c)
	if err != nil {
		ResponseError(c, 1002, err)
		return
	}

	ResponseSuccess(c, out)
	return
}
