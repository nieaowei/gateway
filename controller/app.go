package controller

import (
	"gateway/dto"
	"gateway/lib"
	"gateway/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type AppController struct {
}

func (a *AppController) RouterRegister(group *gin.RouterGroup) {
	group.GET("/list", a.GetAppList)
	group.GET("/add", a.AddApp)
	group.GET("/update", a.UpdateApp)
	group.GET("/detail", a.GetAppDetail)
	group.GET("/del", a.DeleteApp)
}

func (a *AppController) RouterGroupName() (name string) {
	return "/app"
}

func (a *AppController) Middlewares() (middlewares []gin.HandlerFunc) {
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

// GetAppList godoc
// @Summary 获取租户列表
// @Description 获取租户列表
// @Tags 租户接口
// @ID /app/list
// @Accept  json
// @Produce  json
// @Param query query dto.GetAppListInput true "query"
// @Success 200 {object} dto.Response{data=dto.GetAppListOutput} "success"
// @Router /app/list [get]
func (a *AppController) GetAppList(c *gin.Context) {
	exec := &dto.GetAppListInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// GetAppDetail godoc
// @Summary 获取租户详细信息
// @Description 获取租户详细信息
// @Tags 租户接口
// @ID /app/detail
// @Accept  json
// @Produce  json
// @Param query query dto.GetAppDetailInput true "query"
// @Success 200 {object} dto.Response{data=dto.GetAppDetailOutput} "success"
// @Router /app/detail [get]
func (a *AppController) GetAppDetail(c *gin.Context) {
	exec := &dto.GetAppDetailInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// AddApp godoc
// @Summary 增加租户
// @Description 增加租户
// @Tags 租户接口
// @ID /app/add
// @Accept  json
// @Produce  json
// @Param body body dto.AddAppInput true "body"
// @Success 200 {object} dto.Response{data=dto.AddAppOutput} "success"
// @Router /app/add [post]
func (a *AppController) AddApp(c *gin.Context) {
	exec := &dto.AddAppInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// UpdateApp godoc
// @Summary 更新租户
// @Description 更新租户
// @Tags 租户接口
// @ID /app/update
// @Accept  json
// @Produce  json
// @Param body body dto.UpdateAppInput true "body"
// @Success 200 {object} dto.Response{data=dto.UpdateAppOutput} "success"
// @Router /app/update [post]
func (a *AppController) UpdateApp(c *gin.Context) {
	exec := &dto.UpdateAppInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}

// DeleteApp godoc
// @Summary 删除租户
// @Description 删除租户
// @Tags 租户接口
// @ID /app/update
// @Accept  json
// @Produce  json
// @Param body body dto.DeleteAppInput true "body"
// @Success 200 {object} dto.Response{data=dto.DeleteAppOutput} "success"
// @Router /app/update [post]
func (a *AppController) DeleteApp(c *gin.Context) {
	exec := &dto.DeleteAppInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}
