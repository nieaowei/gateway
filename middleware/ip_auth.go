package middleware

import (
	"errors"
	"fmt"
	"gateway/dto"
	"gateway/lib"
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	conf := lib.GetDefaultConfBase()
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range conf.Http.AllowIP {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			dto.ResponseError(c, dto.InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
