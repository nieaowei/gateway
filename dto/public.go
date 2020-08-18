package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type GetAvatarInput struct {
	Username string `json:"username" form:"username" validate:"required"`
}

type GetAvatarOutput struct {
	Avatar string `json:"avatar"`
}

func (g *GetAvatarInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, g)
	params = g
	return
}

func (g *GetAvatarInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p := params.(*GetAvatarInput)
		db := lib.GetDefaultDB()
		admin := &dao.Admin{
			Username: p.Username,
		}
		out = &GetAvatarOutput{}
		err = admin.FindOneScan(c, db, out)
		return
	}
}

func (g *GetAvatarInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		out, err = handle(c)
		if err != nil {
			return
		}
		return
	}
}

func (g *GetAvatarInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		out, err := handle(c)
		if err != nil {
			ResponseError(c, 2001, err)
			return
		}
		ResponseSuccess(c, out)
		return
	}
}
