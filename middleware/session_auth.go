package middleware

import (
	"errors"
	"gateway/dao"
	"gateway/dto"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if adminInfo, ok := session.Get(public.AdminSessionsKey).(*dao.AdminSessionInfo); !ok || adminInfo == nil {
			dto.ResponseError(c, dto.InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
