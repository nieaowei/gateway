package controller

import (
	"gateway/dto"
	"gateway/lib"
	"gateway/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type AdminLoginController struct {
}

func (p *AdminLoginController) Register(group *gin.RouterGroup) {
	group.POST("/login", p.AdminLogin)
}

func (p *AdminLoginController) GroupName() string {
	return "/admin"
}

func (p *AdminLoginController) Middleware() []gin.HandlerFunc {
	store, err := sessions.NewRedisStore(lib.GetDefaultConfRedis().MaxIdle, "tcp", lib.GetDefaultConfRedis().ProxyList[0], "", []byte("secret"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	return []gin.HandlerFunc{
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware(),
	}
}

func (p *AdminLoginController) AdminLogin(c *gin.Context) {
	dto.Exec(&dto.AdminLoginInput{}, c)
	return
}
