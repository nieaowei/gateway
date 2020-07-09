package dto

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
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
