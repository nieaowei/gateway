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
	group.POST("/tcp/add", p.AddTcpService)
	group.POST("/tcp/update", p.UpdateTcpService)
	group.POST("/grpc/add", p.AddGrpcService)
	group.POST("/grpc/update", p.UpdateGrpcService)
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
	exec := &dto.GetServiceDetailInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) AddHttpService(c *gin.Context) {
	exec := &dto.AddHttpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) UpdateHttpService(c *gin.Context) {
	exec := &dto.UpdateHttpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) GetServiceList(c *gin.Context) {
	exec := &dto.GetServiceListInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) DeleteService(c *gin.Context) {
	exec := &dto.DeleteServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) AddTcpService(c *gin.Context) {
	exec := &dto.AddTcpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) UpdateTcpService(c *gin.Context) {
	exec := &dto.UpdateTcpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) AddGrpcService(c *gin.Context) {
	exec := &dto.AddGrpcServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

func (p *ServiceController) UpdateGrpcService(c *gin.Context) {
	exec := &dto.UpdateGrpcServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}
