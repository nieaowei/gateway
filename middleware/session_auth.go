package middleware

import (
	"errors"
	"gateway/controller"
	"gateway/dao"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if adminInfo, ok := session.Get(public.AdminSessionsKey).(*dao.AdminSessionInfo); !ok || adminInfo == nil {
			controller.ResponseError(c, controller.InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
