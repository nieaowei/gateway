package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddGrpcServiceInput struct {
	EditServiceInfo
	EditServiceLoadBalance
	EditServiceAccessControlRule
	EditServiceGRPCRule
}

func (p *AddGrpcServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *AddGrpcServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *AddGrpcServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *AddGrpcServiceInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*AddGrpcServiceInput)
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
				serviceGRPCRule := &dao.ServiceGrpcRule{
					ServiceID:       serviceInfo.ID,
					Port:            p.Port,
					HeaderTransform: p.HeaderTransform,
				}
				err = serviceGRPCRule.InsertAfterCheck(c, tx, true)
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

type UpdateGrpcServiceInput struct {
	ServiceID uint `json:"service_id"`
	AddGrpcServiceInput
}

func (p *UpdateGrpcServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *UpdateGrpcServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *UpdateGrpcServiceInput) Exec(params interface{}, cIn *gin.Context, errIn error) (out interface{},
	cOut *gin.Context, err error) {
	if errIn != nil {
		return nil, cIn, errIn
	}
	p = params.(*UpdateGrpcServiceInput)
	cOut = cIn
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
			err = serviceInfo.UpdateAllByID(cIn, tx)
			if err != nil {
				return
			}
			//insert http rule
			serviceHTTPRule := &dao.ServiceGrpcRule{
				ServiceID:       serviceInfo.ID,
				Port:            p.Port,
				HeaderTransform: p.HeaderTransform,
			}
			err = serviceHTTPRule.UpdateAllByServiceID(cIn, tx)
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
			err = serviceAccessControl.UpdateAllByServiceID(cIn, tx)
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
			err = serviceLoadBalance.UpdateAllByServiceID(cIn, tx)
			if err != nil {
				return
			}
			return
		})
	return
}

func (p *UpdateGrpcServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}
