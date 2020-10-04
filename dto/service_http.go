package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddHttpServiceInput struct {
	//service info
	EditServiceInfo
	// http
	EditServiceHTTPRule
	// access
	EditServiceAccessControlRule
	// loadbalance
	EditServiceLoadBalance
}

func (p *AddHttpServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *AddHttpServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *AddHttpServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *AddHttpServiceInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*AddHttpServiceInput)
		db := lib.GetDefaultDB()
		// start
		err = db.Transaction(
			func(tx *gorm.DB) (err error) {
				// insert service info
				serviceInfo := &dao.ServiceInfo{
					LoadType:    dao.Load_HTTP,
					ServiceName: p.ServiceName,
					ServiceDesc: p.ServiceDesc,
				}
				err = serviceInfo.InsertAfterCheck(c, tx, true)
				if err != nil {
					return
				}
				//insert http rule
				serviceHTTPRule := &dao.ServiceHTTPRule{
					ServiceID:       serviceInfo.ID,
					RuleType:        p.RuleType,
					Rule:            p.Rule,
					NeedHTTPs:       p.NeedHttps,
					NeedStripURI:    p.NeedStripUri,
					NeedWebsocket:   p.NeedWebsocket,
					URLRewrite:      p.UrlRewrite,
					HeaderTransform: p.HeaderTransform,
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
					WhiteList:         p.WhiteList,
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
					ForbidList:             p.ForbidList,
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

type UpdateHttpServiceInput struct {
	ServiceID uint `json:"service_id"`
	AddHttpServiceInput
}

func (p *UpdateHttpServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *UpdateHttpServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *UpdateHttpServiceInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*UpdateHttpServiceInput)
		db := lib.GetDefaultDB()
		err = db.Transaction(
			func(tx *gorm.DB) (err error) {
				// insert service info
				serviceInfo := &dao.ServiceInfo{
					LoadType:    dao.Load_HTTP,
					ServiceName: p.ServiceName,
					ServiceDesc: p.ServiceDesc,
					ID:          p.ServiceID,
				}
				err = serviceInfo.UpdateAllByID(c, tx)
				if err != nil {
					return
				}
				//insert http rule
				serviceHTTPRule := &dao.ServiceHTTPRule{
					ServiceID:       serviceInfo.ID,
					RuleType:        p.RuleType,
					Rule:            p.Rule,
					NeedHTTPs:       p.NeedHttps,
					NeedStripURI:    p.NeedStripUri,
					NeedWebsocket:   p.NeedWebsocket,
					URLRewrite:      p.UrlRewrite,
					HeaderTransform: p.HeaderTransform,
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
					ForbidList:             p.ForbidList,
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

func (p *UpdateHttpServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}
