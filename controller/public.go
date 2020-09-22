package controller

import (
	"gateway/dto"
	"gateway/lib"
	"gateway/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type PublicController struct {
}

func (p *PublicController) RouterRegister(group *gin.RouterGroup) {
	group.GET("/get/avatar", p.GetAvatar)
}

func (p *PublicController) RouterGroupName() (name string) {
	return "/public"
}

func (p *PublicController) Middlewares() (middlewares []gin.HandlerFunc) {
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

// GetAvatar godoc
// @Summary 获取头像
// @Description 获取头像
// @Tags 公共接口
// @ID /public/get/avatar
// @Accept  json
// @Produce  json
// @Param query query dto.GetAvatarInput false "用户名"
// @Success 200 {object} dto.Response{data=dto.GetAvatarOutput} "success"
// @Router /public/get/avatar [get]
func (p *PublicController) GetAvatar(c *gin.Context) {
	exec := &dto.GetAvatarInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}
