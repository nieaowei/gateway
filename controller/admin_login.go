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
	dto.Exec(&dto.AdminLoginInput{}, c)
	return
}
