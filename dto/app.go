package dto

import (
	"errors"
	"gateway/dao"
	"gateway/lib"
	"gateway/proxy/manager"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type GetAppListInput struct {
	Info     string `json:"info" form:"info" example:"http"`
	PageNo   int    `json:"page_no" form:"page_no" example:"2" validate:"required,min=1"`
	PageSize int    `json:"page_size" form:"page_size" example:"10" validate:"required,min=1"`
}

type AppListItem struct {
	ID     uint   `json:"id"`
	AppID  string `json:"app_id"` // 租户id
	Name   string `json:"name"`   // 租户名称
	Secret string `json:"secret"` // 密钥
	//WhiteIPs string `json:"white_ips"` // ip白名单，支持前缀匹配
	Qpd int64 `json:"qpd"` // 日请求量限制
	QPS int64 `json:"qps"`
}

type GetAppListOutput struct {
	List  []AppListItem `json:"list"`
	Total int64         `json:"total"`
}

func (g *GetAppListInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, g)
	params = g
	return
}

func (g *GetAppListInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		data, err := handle(c)
		if err != nil {
			return
		}
		params := data.(*GetAppListInput)
		db := lib.GetDefaultDB()

		app := &dao.App{}
		list, total, err := app.PageListIdDesc(c, db, &dao.PageSize{
			Size: params.PageSize,
			No:   params.PageNo,
			Info: params.Info,
		})
		if err != nil {
			return
		}
		outData := &GetAppListOutput{
			Total: total,
		}
		for _, item := range list {
			counter, ok := manager.Default().GetRedisService(manager.RedisAppPrefix + item.AppID)
			var qpd, qps int64
			if ok {
				impl := counter.(*manager.RedisFlowCountService)
				qpd, _ = impl.GetDayData(time.Now())
				qps = impl.QPS
			}
			outData.List = append(outData.List, AppListItem{
				ID:     item.ID,
				AppID:  item.AppID,
				Name:   item.Name,
				Secret: item.Secret,
				Qpd:    qpd,
				QPS:    qps,
			})
		}
		return outData, nil
	}
}

func (g *GetAppListInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (g *GetAppListInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err == nil {
			ResponseSuccess(c, data)
			return
		}
		ResponseError(c, 2002, err)
		return
	}
}

type GetAppDetailInput struct {
	ID uint `json:"id" form:"id" validate:"required" example:"4"`
}

type GetAppDetailOutput struct {
	AppID string `json:"app_id"`
	Name  string `json:"name"`
	Qps   int64  `json:"qps"`
	Qpd   int64  `json:"qpd"`
}

func (g *GetAppDetailInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, g)
	params = g
	return
}

func (g *GetAppDetailInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		var data interface{}
		data, err = handle(c)
		params := data.(*GetAppDetailInput)
		db := lib.GetDefaultDB()

		app := &dao.App{
			Model: gorm.Model{
				ID: params.ID,
			},
		}
		getAppDetailOutput := &GetAppDetailOutput{}
		err = app.FindOneScan(c, db, getAppDetailOutput)
		out = getAppDetailOutput
		return
	}
}

func (g *GetAppDetailInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (g *GetAppDetailInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		out, err := handle(c)
		if err == nil {
			ResponseSuccess(c, out)
			return
		}
		ResponseError(c, 2002, err)
		return
	}
}

type AddAppInput struct {
	AppID string `json:"app_id" validate:"required" example:"app_1234567"`
	Name  string `json:"name" validate:"required" example:"sevice_1234567"`
	Qps   int64  `json:"qps"`
	Qpd   int64  `json:"qpd"`
}

type AddAppOutput struct {
}

func (a *AddAppInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, a)
	params = a
	return
}

func (a *AddAppInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		var data interface{}
		data, err = handle(c)
		if err != nil {
			return
		}
		db := lib.GetDefaultDB()
		params := data.(*AddAppInput)
		app := &dao.App{
			AppID:  params.AppID,
			Name:   params.Name,
			Secret: public.Md5(params.AppID),
			Qpd:    params.Qpd,
			QPS:    params.Qps,
		}
		err = app.InsertAfterCheck(c, db, true)
		return
	}
}

func (a *AddAppInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (a *AddAppInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		out, err := handle(c)
		if err == nil {
			ResponseSuccess(c, out)
			return
		}
		ResponseError(c, 2002, err)
		return
	}
}

type UpdateAppInput struct {
	ID    uint   `json:"id" validate:"required"`
	AppID string `json:"app_id" validate:"required" example:"app_127"`
	Name  string `json:"name" validate:"required" example:"sevice_1567"`
	Qps   int64  `json:"qps"`
	Qpd   int64  `json:"qpd"`
}

type UpdateAppOutput struct {
}

func (u *UpdateAppInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, u)
	params = u
	return
}

func (u *UpdateAppInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		data, err := handle(c)
		if err != nil {
			return
		}
		params := data.(*UpdateAppInput)
		db := lib.GetDefaultDB()
		app := &dao.App{
			Model: gorm.Model{
				ID: params.ID,
			},
			AppID:  params.AppID,
			Name:   params.Name,
			Secret: public.Md5(params.AppID),
			Qpd:    params.Qpd,
			QPS:    params.Qps,
		}
		err = app.UpdateAllById(c, db)
		return
	}
}

func (u *UpdateAppInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (u *UpdateAppInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err == nil {
			ResponseSuccess(c, data)
			return
		}
		ResponseError(c, 2002, err)
		return
	}
}

type DeleteAppInput struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

type DeleteAppOutput struct {
}

func (d *DeleteAppInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, d)
	params = d
	return
}

func (d *DeleteAppInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		data, err := handle(c)
		if err != nil {
			return
		}
		db := lib.GetDefaultDB()
		params := data.(*DeleteAppInput)
		app := &dao.App{
			Model: gorm.Model{
				ID: params.ID,
			},
		}
		err = app.DeleteByID(c, db)
		return
	}
}

func (d *DeleteAppInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (d *DeleteAppInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err == nil {
			ResponseSuccess(c, data)
			return
		}
		ResponseError(c, 2002, err)
		return
	}
}

type GetAppStatInput struct {
	ID uint `json:"id" form:"id" example:"3" validate:"required"`
}

type GetAppStatOutput struct {
	TodayList     []int `json:"today_list"`
	YesterdayList []int `json:"yesterday_list"`
}

func (g *GetAppStatInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, g)
	params = g
	return
}

func (g *GetAppStatInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		inter, err := handle(c)
		if err != nil {
			return
		}
		params := inter.(*GetAppStatInput)
		db := lib.GetDefaultDB()
		type Select struct {
			AppID string
		}
		s := &Select{}
		err = (&dao.App{
			Model: gorm.Model{ID: params.ID},
		}).FindOneScan(c, db, s)

		if err != nil {
			err = errors.New("不存在该租户")
			return
		}

		data := &GetServiceStatOutput{}
		redisService, ok := manager.Default().GetRedisService(manager.RedisAppPrefix + s.AppID)
		if !ok {
			err = errors.New("没有可利用的Redis服务")
			return
		}
		totalService := redisService.(*manager.RedisFlowCountService)
		currentTime := time.Now().In(manager.TimeLocation)
		for i := 0; i <= currentTime.Hour(); i++ {
			dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, manager.TimeLocation)
			hourData, _ := totalService.GetHourData(dateTime)
			data.TodayList = append(data.TodayList, hourData)
		}

		yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
		for i := 0; i <= 23; i++ {
			dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, manager.TimeLocation)
			hourData, _ := totalService.GetHourData(dateTime)
			data.YesterdayList = append(data.YesterdayList, hourData)
		}

		//data.TodayList = append(data.TodayList, []int{1, 32, 54, 212, 432, 453, 123, 312}...)
		//data.YesterdayList = append(data.YesterdayList, []int{32, 3, 23, 43, 43, 123, 121, 44}...)
		return data, nil
	}
}

func (g *GetAppStatInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (g *GetAppStatInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err == nil {
			ResponseSuccess(c, data)
			return
		}
		ResponseError(c, 1001, err)
		return
	}
}
