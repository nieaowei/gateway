package controller

import (
	"gateway/dao"
	"gateway/dto"
	"gateway/lib"
	"gateway/middleware"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type AdminController struct {
}

func (p *AdminController) RouterGroupName() string {
	return "/admin"
}

func (p *AdminController) Middleware() []gin.HandlerFunc {
	store, err := sessions.NewRedisStore(lib.GetDefaultConfRedis().MaxIdle, "tcp", lib.GetDefaultConfRedis().ProxyList[0], "", []byte("secret"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	return []gin.HandlerFunc{
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware(),
	}
}

//AdminRegister is used for router registration.
func (p *AdminController) RouterRegister(group *gin.RouterGroup) {
	group.GET("/admin_info", p.AdminInfo)
	group.GET("/logout", p.AdminLogout)
	group.POST("/change/password", p.AdminChangePwd)
}

//AdminLogout is the administrator login interface.
func (p *AdminController) AdminLogout(c *gin.Context) {
	sessions.Default(c).Delete(public.AdminSessionsKey)
	err := sessions.Default(c).Save()
	if err != nil {
		dto.ResponseError(c, 10000, err)
		return
	}
	dto.ResponseSuccess(c, nil)
	return
}

//AdminInfo is an interface for obtaining user information.
func (p *AdminController) AdminInfo(c *gin.Context) {
	adminSession := sessions.Default(c).Get(public.AdminSessionsKey).(*dao.AdminSessionInfo)
	adminInfo := dto.AdminInfoOutput{
		AdminSessionInfo: adminSession,
		Avatar:           "",
		Introduction:     "",
		Roles:            nil,
	}
	dto.ResponseSuccess(c, adminInfo)
	return
}

//AdminChangePwd
func (p *AdminController) AdminChangePwd(c *gin.Context) {
	dto.Exec(&dto.AdminChangePwdInput{}, c)
	return
}
