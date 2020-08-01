package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	db, err := lib.GetDefaultDB()
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
	err = adminInfo.UpdateByID(c, db)
	if err != nil {
		return
	}
	return
}
