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

//AdminRegister is used for router registration.
func AdminRegister(group *gin.RouterGroup) {
	admin := &AdminController{}
	group.GET("/admin_info", admin.AdminInfo)
	group.GET("/logout", admin.AdminLogout)
	group.POST("/change/password", admin.ChangePassword)
}

//AdminLogout is the administrator login interface.
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

//AdminInfo is an interface for obtaining user information.
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

//ChangePassword
func (p *AdminController) ChangePassword(c *gin.Context) {
	//get parameters and validate it.
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	//pass
	err := params.ChangePwd(c)
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
