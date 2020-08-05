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
	Username string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required,is_valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"12345" validate:"required"`
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token"`
}

func (p *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, p)
}

func (p *AdminLoginInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

func (p *AdminLoginInput) OutputHandle(c *gin.Context, outIn interface{}) (out interface{}) {
	return outIn
}

func (p *AdminLoginInput) Exec(c *gin.Context) (out interface{}, err error) {
	adminInfo := &dao.Admin{
		Username: p.Username,
	}
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	adminInfo, err = adminInfo.FindOne(c, db)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	saltPd := public.GenSha256BySecret(p.Password, adminInfo.Salt)

	if saltPd != adminInfo.Password {
		return nil, errors.New("密码错误")
	}
	//set sessions.
	adminSession := &dao.AdminSessionInfo{
		ID:        adminInfo.ID,
		Username:  adminInfo.Username,
		LoginTime: time.Now(),
	}
	sess := sessions.Default(c)
	sess.Set(public.AdminSessionsKey, adminSession)
	err = sess.Save()
	if err != nil {
		return nil, err
	}

	out = &AdminLoginOutput{Token: adminInfo.Password}

	return
}
