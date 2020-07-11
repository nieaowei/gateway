package dto

import (
	"gateway/dao"
	"gateway/public"
	"github.com/e421083458/golang_common/lib"
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
	adminSession := sessions.Default(c).Get(public.AdminSessionsKey).(*dao.AdminSessionInfo)
	adminInfo := &dao.Admin{
		Username: adminSession.Username,
	}
	db, err := lib.GetGormPool("default")
	if err != nil {
		return
	}
	adminInfo, err = adminInfo.FindOne(c, db)
	if err != nil {
		return
	}
	adminInfo.Password = public.GenSha256BySecret(p.Password, adminInfo.Salt)
	err = adminInfo.Save(c, db)
	if err != nil {
		return
	}
	return
}
