package controller

import (
	"gateway/dto"
	"gateway/lib"
	"gateway/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type ServiceController struct {
}

func (p *ServiceController) RouterRegister(group *gin.RouterGroup) {
	group.GET("/list", p.GetServiceList)
	group.GET("/del", p.DeleteService)
	group.POST("/http/add", p.AddHttpService)
	group.POST("/http/update", p.UpdateHttpService)
	group.GET("/detail", p.GetServiceDetail)
}

func (p *ServiceController) RouterGroupName() string {
	return "/service"
}

func (p *ServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
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
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware(),
	}
}

func (p *ServiceController) GetServiceDetail(c *gin.Context) {
	dto.Exec(&dto.GetServiceDetailInput{}, c)
	return
}

func (p *ServiceController) AddHttpService(c *gin.Context) {
	dto.Exec(&dto.AddHttpServiceInput{}, c)
	return
}

func (p *ServiceController) UpdateHttpService(c *gin.Context) {
	dto.Exec(&dto.UpdateHttpServiceInput{}, c)
	return
}

func (p *ServiceController) GetServiceList(c *gin.Context) {
	dto.Exec(&dto.GetServiceListInput{}, c)
	return
}

func (p *ServiceController) DeleteService(c *gin.Context) {
	dto.Exec(&dto.DeleteServiceInput{}, c)
	return
}
