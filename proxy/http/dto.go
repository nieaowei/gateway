package proxy_http

import (
	"gateway/dto"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

type TokenInput struct {
	GrantType string `json:"grant_type" form:"grant_type" comment:"授权类型" example:"client_credentials" validate:"required"`
	Scope     string `json:"scope" form:"scope" comment:"权限范围" example:"read_write" validate:"required"`
}

type TokensOutput struct {
	AccessToken string `json:"access_token" form:"access_token"` //access_token
	ExpiresIn   int    `json:"expires_in" form:"expires_in"`     //expires_in
	TokenType   string `json:"token_type" form:"token_type"`     //token_type
	Scope       string `json:"scope" form:"scope"`               //scope
}

func (t *TokenInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, t)
	params = *t
	return
}

func (t *TokenInput) ExecHandle(handle dto.FunctionalHandle) dto.FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		data, err := handle(c)
		if err != nil {
			return
		}
		_ = data.(TokenInput)
		secret := strings.Split(c.GetHeader("Authorization"), " ")
		if len(secret) != 2 {
			return nil, errors.New("Authorization error")
		}
		return
	}
}

func (t *TokenInput) OutputHandle(handle dto.FunctionalHandle) dto.FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (t *TokenInput) ErrorHandle(handle dto.FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err == nil {
			dto.ResponseSuccess(c, data)
			return
		}
		dto.ResponseError(c, 2002, err)
		return
	}
}
