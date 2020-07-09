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
	group.POST("/login",adminLogin.AdminLogin)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (p *AdminLoginController) AdminLogin(c *gin.Context) {
	params := & dto.AdminLoginInput{}
	if err := params.BindValidParam(c);err!=nil {
		middleware.ResponseError(c,1001,err)
		return
	}
	out := dto.AdminLoginOutput{Token: params.Username}
	middleware.ResponseSuccess(c,out)
}
