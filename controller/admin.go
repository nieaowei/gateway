package controller

import (
	"gateway/dao"
	"gateway/dto"
	"gateway/middleware"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

/******sql******
CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='管理员表'
******sql******/
// Admin 管理员表
type Admin struct {
	gorm.Model
	Username string // 用户名
	Salt     string // 盐
	Password string // 密码
}
