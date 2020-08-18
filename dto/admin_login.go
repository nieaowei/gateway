package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

type AdminLoginInput struct {
	//用户名
	Username string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required,is_valid_username"`
	//密码
	Password string `json:"password" form:"password" comment:"密码" example:"12345" validate:"required"`
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token"`
}

func (p *AdminLoginInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *AdminLoginInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *AdminLoginInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *AdminLoginInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*AdminLoginInput)
		adminInfo := &dao.Admin{
			Username: p.Username,
		}
		db := lib.GetDefaultDB()
		adminInfo, err = adminInfo.FindOne(c, db)
		if err != nil {
			return out, errors.New("用户不存在")
		}
		saltPd := public.GenSha256BySecret(p.Password, adminInfo.Salt)

		if saltPd != adminInfo.Password {
			return out, errors.New("密码错误")
		}
		//set sessions.
		adminSession := &dao.AdminSessionInfo{
			ID:        adminInfo.ID,
			Username:  adminInfo.Username,
			LoginTime: time.Now(),
			Avatar:    adminInfo.Avatar,
		}
		sess := sessions.Default(c)
		sess.Set(public.AdminSessionsKey, adminSession)
		err = sess.Save()
		if err != nil {
			return out, err
		}

		out = &AdminLoginOutput{Token: adminInfo.Password}

		return
	}

}
