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

// GetServiceDetail godoc
// @Summary 获服务详情
// @Description 获取单个服务想起
// @Tags 服务接口
// @ID /service/detail
// @Accept  json
// @Produce  json
// @Param query query dto.GetServiceDetailInput true "query"
// @Success 200 {object} dto.Response{data=dto.GetServiceDetailForHttpOutput} "success"
// @Router /service/detail [get]
func (p *ServiceController) GetServiceDetail(c *gin.Context) {
	exec := &dto.GetServiceDetailInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// AddHttpService godoc
// @Summary 增加http服务
// @Description 增加http服务
// @Tags 服务接口
// @ID /service/http/add
// @Accept  json
// @Produce  json
// @Param body body dto.AddHttpServiceInput true "body"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/http/add [post]
func (p *ServiceController) AddHttpService(c *gin.Context) {
	exec := &dto.AddHttpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(dto.IpListAndWeightListNumValid(exec.BindValidParam))))(c)
	return
}

// UpdateHttpService godoc
// @Summary 更新http服务
// @Description 更新http服务
// @Tags 服务接口
// @ID /service/http/update
// @Accept  json
// @Produce  json
// @Param body body dto.UpdateHttpServiceInput true "body"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/http/update [post]
func (p *ServiceController) UpdateHttpService(c *gin.Context) {
	exec := &dto.UpdateHttpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(dto.IpListAndWeightListNumValid(exec.BindValidParam))))(c)
	return
}

// GetServiceList godoc
// @Summary 获取服务列表
// @Description 获取服务列表
// @Tags 服务接口
// @ID /service/list
// @Accept  json
// @Produce  json
// @Param query query dto.GetServiceListInput true "query"
// @Success 200 {object} dto.Response{data=dto.GetServiceListOutput} "success"
// @Router /service/list [get]
func (p *ServiceController) GetServiceList(c *gin.Context) {
	exec := &dto.GetServiceListInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// DeleteService godoc
// @Summary 删除服务
// @Description 删除服务
// @Tags 服务接口
// @ID /service/del
// @Accept  json
// @Produce  json
// @Param query query dto.DeleteServiceInput true "query"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/del [get]
func (p *ServiceController) DeleteService(c *gin.Context) {
	exec := &dto.DeleteServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// AddTcpService godoc
// @Summary 增加tcp服务
// @Description 增加tcp服务
// @Tags 服务接口
// @ID /service/tcp/add
// @Accept  json
// @Produce  json
// @Param body body dto.AddTcpServiceInput true "body"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/tcp/add [post]
func (p *ServiceController) AddTcpService(c *gin.Context) {
	exec := &dto.AddTcpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(dto.IpListAndWeightListNumValid(exec.BindValidParam))))(c)
	return
}

// UpdateTcpService godoc
// @Summary 更新tcp服务
// @Description 更新tcp服务
// @Tags 服务接口
// @ID /service/tcp/update
// @Accept  json
// @Produce  json
// @Param body body dto.UpdateTcpServiceInput true "body"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/tcp/update [post]
func (p *ServiceController) UpdateTcpService(c *gin.Context) {
	exec := &dto.UpdateTcpServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(dto.IpListAndWeightListNumValid(exec.BindValidParam))))(c)
	return
}

// AddGrpcService godoc
// @Summary 增加grpc服务
// @Description 增加grpc服务
// @Tags 服务接口
// @ID /service/grpc/add
// @Accept  json
// @Produce  json
// @Param body body dto.AddGrpcServiceInput true "body"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/grpc/add [post]
func (p *ServiceController) AddGrpcService(c *gin.Context) {
	exec := &dto.AddGrpcServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(dto.IpListAndWeightListNumValid(exec.BindValidParam))))(c)
	return
}

// UpdateGrpcService godoc
// @Summary 更新grpc服务
// @Description 更新grpc服务
// @Tags 服务接口
// @ID /service/grpc/update
// @Accept  json
// @Produce  json
// @Param body body dto.UpdateGrpcServiceInput true "body"
// @Success 200 {object} dto.Response{} "success"
// @Router /service/grpc/update [post]
func (p *ServiceController) UpdateGrpcService(c *gin.Context) {
	exec := &dto.UpdateGrpcServiceInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(dto.IpListAndWeightListNumValid(exec.BindValidParam))))(c)
	return
}
