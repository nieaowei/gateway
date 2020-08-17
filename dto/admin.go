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

type AdminChangePwdInput struct {
	Password string `json:"password" form:"password" validate:"required"`
}

func (p *AdminChangePwdInput) BindValidParam(cIn *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(cIn, p)
	params = p
	return
}

func (p *AdminChangePwdInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		out, err := handle(c)
		if err != nil {
			ResponseError(c, 1002, err)
			return
		}
		ResponseSuccess(c, out)
		return
	}
}

func (p *AdminChangePwdInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *AdminChangePwdInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*AdminChangePwdInput)
		// get session information.
		adminSession := sessions.Default(c).Get(public.AdminSessionsKey).(*dao.AdminSessionInfo)
		adminInfo := &dao.Admin{
			Username: adminSession.Username,
		}
		// get database.
		db := lib.GetDefaultDB()
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

}
