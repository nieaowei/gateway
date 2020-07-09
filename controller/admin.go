package controller

import (
	"gateway/dao"
	"gateway/dto"
	"gateway/middleware"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admin := &AdminController{}
	group.GET("/admin_info", admin.AdminInfo)
	group.GET("/logout", admin.AdminLogout)
}

func (p *AdminController) AdminLogout(c *gin.Context) {
	sessions.Default(c).Delete(public.AdminSessionsKey)
	err := sessions.Default(c).Save()
	if err != nil {
		middleware.ResponseError(c, 10000, err)
		return
	}
	middleware.ResponseSuccess(c, nil)
	return
}

func (p *AdminController) AdminInfo(c *gin.Context) {
	adminSession := sessions.Default(c).Get(public.AdminSessionsKey).(*dao.AdminSessionInfo)
	adminInfo := dto.AdminInfoOutput{
		AdminSessionInfo: adminSession,
		Avatar:           "",
		Introduction:     "",
		Roles:            nil,
	}
	middleware.ResponseSuccess(c, adminInfo)
	return
}
