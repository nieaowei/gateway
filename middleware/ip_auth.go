package middleware

import (
	"errors"
	"fmt"
	"gateway/controller"
	"gateway/lib"
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range lib.GetDefaultConfBase().Http.AllowIP {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			controller.ResponseError(c, controller.InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
