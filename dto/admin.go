package dto

import (
	"gateway/dao"
	"gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminInfoOutput struct {
	*dao.AdminSessionInfo
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
	Roles        []string `json:"roles"`
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" validate:"required"`
}

func (p *ChangePwdInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *ChangePwdInput) ChangePwd(c *gin.Context) (err error) {
	// get session information.
	adminSession := sessions.Default(c).Get(public.AdminSessionsKey).(*dao.AdminSessionInfo)
	adminInfo := &dao.Admin{
		Username: adminSession.Username,
	}
	// get database.
	db, err := lib.GetGormPool("default")
	if err != nil {
		return
	}
	// get admin information by username.
	adminInfo, err = adminInfo.FindOne(c, db)
	if err != nil {
		return
	}
	// update password by id
	adminInfo = &dao.Admin{
		Model: gorm.Model{
			ID: adminInfo.ID,
		},
		Password: public.GenSha256BySecret(p.Password, adminInfo.Salt),
	}
	err = adminInfo.Updates(c, db)
	if err != nil {
		return
	}
	return
}
