package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddTcpServiceInput struct {
	EditServiceInfo
	EditServiceLoadBalance
	EditServiceAccessControlRule
	EditServiceTCPRule
}

func (p *AddTcpServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *AddTcpServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		out, err := handle(c)
		if err != nil {
			ResponseError(c, 1002, err)
			return
		}
		ResponseSuccess(c, out)
		return
	}
}

func (p *AddTcpServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *AddTcpServiceInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*AddTcpServiceInput)
		db := lib.GetDefaultDB()
		// start
		err = db.Transaction(
			func(tx *gorm.DB) (err error) {
				// insert service info
				serviceInfo := &dao.ServiceInfo{
					LoadType:    dao.LoadType_HTTP,
					ServiceName: p.ServiceName,
					ServiceDesc: p.ServiceDesc,
				}
				err = serviceInfo.InsertAfterCheck(c, tx, true)
				if err != nil {
					return
				}
				//insert http rule
				serviceHTTPRule := &dao.ServiceTCPRule{
					ServiceID: serviceInfo.ID,
					Port:      p.Port,
				}
				err = serviceHTTPRule.InsertAfterCheck(c, tx, true)
				if err != nil {
					return
				}
				// insert accesscontrol
				serviceAccessControl := &dao.ServiceAccessControl{
					ServiceID:         serviceInfo.ID,
					OpenAuth:          p.OpenAuth,
					BlackList:         p.BlackList,
					WhiteList:         p.WeightList,
					WhiteHostName:     p.WhiteHostName,
					ClientipFlowLimit: p.ClientipFlowLimit,
					ServiceFlowLimit:  p.ServiceFlowLimit,
				}
				err = serviceAccessControl.InsertAfterCheck(c, tx, true)
				if err != nil {
					return
				}
				// insert loadbalance
				serviceLoadBalance := &dao.ServiceLoadBalance{
					ServiceID:              serviceInfo.ID,
					RoundType:              p.RoundType,
					IPList:                 p.IpList,
					WeightList:             p.WeightList,
					ForbidList:             p.ForbidLIst,
					UpstreamConnectTimeout: p.UpstreamConnectTimeout,
					UpstreamHeaderTimeout:  p.UpstreamHeaderTimeout,
					UpstreamIDleTimeout:    p.UpstreamIdleTimeout,
					UpstreamMaxIDle:        p.UpstreamMaxIdle,
				}
				err = serviceLoadBalance.InsertAfterCheck(c, tx, true)
				if err != nil {
					return
				}
				return
			})
		return
	}

}

type UpdateTcpServiceInput struct {
	ServiceID uint `json:"service_id"`
	AddTcpServiceInput
}

func (p *UpdateTcpServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *UpdateTcpServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
	return func(c *gin.Context) {
		out, err := handle(c)
		if err != nil {
			ResponseError(c, 1002, err)
			return
		}
		ResponseSuccess(c, out)
		return
	}
}

func (p *UpdateTcpServiceInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*UpdateTcpServiceInput)
		db := lib.GetDefaultDB()
		err = db.Transaction(
			func(tx *gorm.DB) (err error) {
				// insert service info
				serviceInfo := &dao.ServiceInfo{
					LoadType:    dao.LoadType_HTTP,
					ServiceName: p.ServiceName,
					ServiceDesc: p.ServiceDesc,
					ID:          p.ServiceID,
				}
				err = serviceInfo.UpdateAllByID(c, tx)
				if err != nil {
					return
				}
				//insert http rule
				serviceHTTPRule := &dao.ServiceTCPRule{
					ServiceID: serviceInfo.ID,
					Port:      p.Port,
				}
				err = serviceHTTPRule.UpdateAllByServiceID(c, tx)
				if err != nil {
					return
				}
				// insert accesscontrol
				serviceAccessControl := &dao.ServiceAccessControl{
					ServiceID:         serviceInfo.ID,
					OpenAuth:          p.OpenAuth,
					BlackList:         p.BlackList,
					WhiteList:         p.WhiteList,
					WhiteHostName:     p.WhiteHostName,
					ClientipFlowLimit: p.ClientipFlowLimit,
					ServiceFlowLimit:  p.ServiceFlowLimit,
				}
				err = serviceAccessControl.UpdateAllByServiceID(c, tx)
				if err != nil {
					return
				}
				// insert loadbalance
				serviceLoadBalance := &dao.ServiceLoadBalance{
					ServiceID:              serviceInfo.ID,
					RoundType:              p.RoundType,
					IPList:                 p.IpList,
					WeightList:             p.WeightList,
					ForbidList:             p.ForbidLIst,
					UpstreamConnectTimeout: p.UpstreamConnectTimeout,
					UpstreamHeaderTimeout:  p.UpstreamHeaderTimeout,
					UpstreamIDleTimeout:    p.UpstreamIdleTimeout,
					UpstreamMaxIDle:        p.UpstreamMaxIdle,
				}
				err = serviceLoadBalance.UpdateAllByServiceID(c, tx)
				if err != nil {
					return
				}
				return
			})
		return
	}

}

func (p *UpdateTcpServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}
