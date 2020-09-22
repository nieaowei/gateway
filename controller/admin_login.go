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
	//store, err := sessions.NewRedisStore(
	//	conf.MaxIdle,
	//	"tcp",
	//	conf.ProxyList[0],
	//	"1234",
	//	[]byte("secret"),
	//)
	store, err := lib.NewRedisStoreClusterCli(lib.NewRedisClusterCli(conf), []byte("secret"))

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

// AdminLogin godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} dto.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin/login [post]
func (p *AdminLoginController) AdminLogin(c *gin.Context) {
	exec := dto.AdminLoginInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}
