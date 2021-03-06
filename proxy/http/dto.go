package proxy_http

import (
	"gateway/dto"
	"gateway/proxy/manager"
	"gateway/public"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type TokenInput struct {
	GrantType string `json:"grant_type" form:"grant_type" comment:"授权类型" example:"client_credentials" validate:"required"`
	Scope     string `json:"scope" form:"scope" comment:"权限范围" example:"read_write" validate:"required"`
}

type TokenType string

const (
	TokenType_Bearer TokenType = "Bearer"
)

const (
	AuthHeaderKey = "Authorization"
)

type TokensOutput struct {
	AccessToken string    `json:"access_token" form:"access_token"` //access_token
	ExpiresIn   int       `json:"expires_in" form:"expires_in"`     //expires_in
	TokenType   TokenType `json:"token_type" form:"token_type"`     //token_type
	Scope       string    `json:"scope" form:"scope"`               //scope
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
		params := data.(TokenInput)
		secret := strings.Split(c.GetHeader(AuthHeaderKey), " ")
		if len(secret) != 2 {
			return nil, Error_AuthFormat
		}
		authInfo := strings.Split(secret[1], ":")
		if len(authInfo) != 2 {
			return nil, Error_AuthInfoFormat
		}
		appId, appSecert := authInfo[0], authInfo[1]
		inter, ok := manager.Default().APPMap.Get(appId)
		if !ok {
			return nil, Error_AppNotFound
		}
		appInfo := inter.(*manager.App)
		if appId != appInfo.AppId || appSecert != appInfo.Secret {
			return nil, Error_AppNotMatched
		}
		claims := jwt.StandardClaims{
			Issuer:    appInfo.AppId,
			ExpiresAt: time.Now().Add(time.Second * JwtExpireAt).In(manager.TimeLocation).Unix(),
		}
		token, err := JwtEncode(claims)
		if err != nil {
			return
		}

		o := TokensOutput{
			AccessToken: token,
			ExpiresIn:   JwtExpireAt,
			TokenType:   TokenType_Bearer,
			Scope:       params.GrantType,
		}

		return o, nil
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
		code := err.(*ProxyError)
		dto.ResponseError(c, code.Code, err)
		return
	}
}
