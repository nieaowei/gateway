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

// AdminLogin godoc
// @Summary 测试数据绑定
// @Description 测试数据绑定
// @Tags 用户
// @ID /admin/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin/login [post]
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
