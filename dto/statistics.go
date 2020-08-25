package dto

import (
	"gateway/dao"
	"gateway/lib"
	"github.com/gin-gonic/gin"
)

type GetTotalInput struct {
}

type GetTotalOutput struct {
	ServiceAmount int64 `json:"service_amount"`
	QPS           int64 `json:"qps"`
	QPD           int64 `json:"qpd"`
	TenantAmount  int64 `json:"tenant_amount"`
}

func (g *GetTotalInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	return
}

func (g *GetTotalInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		db := lib.GetDefaultDB()
		data := GetTotalOutput{}
		err = db.Model(&dao.ServiceInfo{}).Count(&data.ServiceAmount).Error
		if err != nil {
			return nil, err
		}
		err = db.Model(&dao.App{}).Count(&data.TenantAmount).Error
		return data, err
	}
}

func (g *GetTotalInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (g *GetTotalInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

type GetServiceAmountInput struct {
}

type GetServiceAmountOutput struct {
	TCP  int64 `json:"TCP"`
	HTTP int64 `json:"HTTP"`
	GRPC int64 `json:"GRPC"`
}

func (g *GetServiceAmountInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	return
}

func (g *GetServiceAmountInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		_, err = handle(c)
		if err != nil {
			return
		}
		db := lib.GetDefaultDB()
		data := &GetServiceAmountOutput{}
		err = db.Model(&dao.ServiceHTTPRule{}).Count(&data.HTTP).Error
		if err != nil {
			return nil, err
		}
		err = db.Model(&dao.ServiceTCPRule{}).Count(&data.TCP).Error
		if err != nil {
			return nil, err
		}
		err = db.Model(&dao.ServiceGrpcRule{}).Count(&data.GRPC).Error
		if err != nil {
			return nil, err
		}
		return data, err
	}
}

func (g *GetServiceAmountInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (g *GetServiceAmountInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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
