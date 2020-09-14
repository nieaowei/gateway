package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type GetAppListInput struct {
	Info     string `json:"info" form:"info"`
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
			outData.List = append(outData.List, AppListItem{
				ID:     item.ID,
				AppID:  item.AppID,
				Name:   item.Name,
				Secret: item.Secret,
				Qpd:    item.Qpd,
				QPS:    item.QPS,
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
	ID uint `json:"id" form:"id" validate:"required"`
}

type GetAppDetailOutput struct {
}

func (g *GetAppDetailInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	panic("implement me")
}

func (g *GetAppDetailInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	panic("implement me")
}

func (g *GetAppDetailInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	panic("implement me")
}

func (g *GetAppDetailInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	panic("implement me")
}

type AddAppInput struct {
	AppID string `json:"app_id"`
	Name  string `json:"name"`
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
	panic("implement me")
}

func (a *AddAppInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	panic("implement me")
}

type UpdateAppInput struct {
}

type UpdateAppOutput struct {
}

func (u *UpdateAppInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	panic("implement me")
}

func (u *UpdateAppInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	panic("implement me")
}

func (u *UpdateAppInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	panic("implement me")
}

func (u *UpdateAppInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	panic("implement me")
}

type DeleteAppInput struct {
}

type DeleteAppOutput struct {
}

func (d *DeleteAppInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	panic("implement me")
}

func (d *DeleteAppInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	panic("implement me")
}

func (d *DeleteAppInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	panic("implement me")
}

func (d *DeleteAppInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	panic("implement me")
}
