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

func (p *AdminLoginController) RouterRegister(group *gin.RouterGroup) {
	group.POST("/login", p.AdminLogin)
}

func (p *AdminLoginController) RouterGroupName() string {
	return "/admin"
}

func (p *AdminLoginController) Middlewares() (middlewares []gin.HandlerFunc) {
	conf := lib.GetDefaultConfRedis()
	store, err := sessions.NewRedisStore(
		conf.MaxIdle,
		"tcp",
		conf.ProxyList[0],
		"",
		[]byte("secret"),
	)
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
