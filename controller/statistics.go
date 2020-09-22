package controller

import (
	"gateway/dto"
	"gateway/lib"
	"gateway/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type StatisticsController struct {
}

func (s *StatisticsController) RouterRegister(group *gin.RouterGroup) {
	group.GET("/service/amount", s.GetServiceAmount)
	group.GET("/total", s.GetTotal)
}

func (s *StatisticsController) RouterGroupName() (name string) {
	return "/statistics"
}

func (s *StatisticsController) Middlewares() (middlewares []gin.HandlerFunc) {
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
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware(),
	}
}

// GetServiceAmount godoc
// @Summary 获服务统计
// @Description 获服务统计
// @Tags 统计接口
// @ID /statistics/service/amount
// @Accept  json
// @Produce  json
// @Param query query dto.GetServiceAmountInput true "query"
// @Success 200 {object} dto.Response{data=dto.GetServiceAmountOutput} "success"
// @Router /statistics/service/amount [get]
func (s *StatisticsController) GetServiceAmount(c *gin.Context) {
	exec := dto.GetServiceAmountInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
}

// GetTotal godoc
// @Summary 获取仪表盘统计数据
// @Description 获取仪表盘统计数据
// @Tags 统计接口
// @ID /statistics/total
// @Accept  json
// @Produce  json
// @Param query query dto.GetTotalInput true "query"
// @Success 200 {object} dto.Response{data=dto.GetTotalOutput} "success"
// @Router  /statistics/total [get]
func (s *StatisticsController) GetTotal(c *gin.Context) {
	exec := &dto.GetTotalInput{}
	exec.ErrorHandle(exec.OutputHandle(exec.ExecHandle(exec.BindValidParam)))(c)
	return
}
