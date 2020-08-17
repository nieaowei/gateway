package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type GetServiceListInput struct {
	Info     string `json:"info" form:"info"`
	PageNo   int    `json:"page_no" form:"page_no" validate:"required"`
	PageSize int    `json:"page_size" form:"page_size" validate:"required"`
}

type ServiceListItem struct {
	ID          uint   `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType    string `json:"load_type" form:"load_type"`
	Address     string `json:"address" form:"Address"`
	Qps         uint   `json:"qps" form:"qps"`
	Qpd         uint   `json:"qpd" form:"qpd"`
	TotalNode   uint   `json:"total_node" form:"total_node"`
}

type GetServiceListOutput struct {
	Total int64             `json:"total" form:"total" validate:""`
	List  []ServiceListItem `json:"list" form:"list" validate:""`
}

func (p *GetServiceListInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *GetServiceListInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *GetServiceListInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *GetServiceListInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*GetServiceListInput)
		serviceInfo := &dao.ServiceInfo{}
		db := lib.GetDefaultDB()
		serviceInfos, count, err := serviceInfo.PageList(c, db, &dao.PageSize{
			Size: p.PageSize,
			No:   p.PageNo,
			Info: p.Info,
		})
		if err != nil {
			return
		}
		outE := &GetServiceListOutput{
			Total: count,
			List:  []ServiceListItem{},
		}
		conf := lib.GetDefaultConfBase()
		clusterIP := conf.Cluster.Ip
		clusterPort := conf.Cluster.Port
		clusterSSLPort := conf.Cluster.SslPort
		for _, info := range serviceInfos {
			serviceDetail, err := info.FindOneServiceDetail(c, db)
			if err != nil {
				return nil, err
			}
			serviceAddr := ""
			loadType := ""
			switch serviceDetail.LoadType {
			case dao.LoadTypeHttp:
				{
					loadType = "HTTP"
					service := serviceDetail.ServiceHTTPRuleExceptModel
					if service.RuleType == dao.HttpRuleTypePrefixURL && service.NeedHTTPs == 0 {
						serviceAddr = clusterIP + ":" + clusterPort + service.Rule

					}
					if service.RuleType == dao.HttpRuleTypePrefixURL && service.NeedHTTPs == 1 {
						serviceAddr = clusterIP + ":" + clusterSSLPort + service.Rule
					}
					if service.RuleType == dao.HttpRuleTypeDomain {
						serviceAddr = service.Rule
					}
					break
				}
			case dao.LoadTypeTcp:
				{
					loadType = "TCP"
					service := serviceDetail.ServiceTCPRuleExceptModel

					serviceAddr = clusterIP + ":" + strconv.Itoa(service.Port)
					break
				}
			case dao.LoadTypeGrpc:
				{
					loadType = "GRPC"
					service := serviceDetail.ServiceGrpcRuleExceptModel

					serviceAddr = clusterIP + ":" + strconv.Itoa(service.Port)
					break
				}
			}
			ipList := serviceDetail.ServiceLoadBalanceExceptModel.GetIPListByModel()

			item := ServiceListItem{
				ID:          info.ID,
				ServiceName: info.ServiceName,
				ServiceDesc: info.ServiceDesc,
				LoadType:    loadType,
				Address:     serviceAddr,
				Qps:         0,
				Qpd:         0,
				TotalNode:   uint(len(ipList)),
			}
			outE.List = append(outE.List, item)
		}
		return
	}

}

type DeleteServiceInput struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

func (p *DeleteServiceInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *DeleteServiceInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *DeleteServiceInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *DeleteServiceInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*DeleteServiceInput)
		db := lib.GetDefaultDB()
		serviceInfo := &dao.ServiceInfo{
			ID: p.ID,
		}

		err = db.Transaction(
			func(tx *gorm.DB) (err error) {
				serviceInfo, err = serviceInfo.FindOne(c, tx)
				if err != nil {
					return
				}

				switch serviceInfo.LoadType {
				case dao.LoadTypeHttp:
					{
						http := dao.ServiceHTTPRule{
							ServiceID: serviceInfo.ID,
						}
						err = http.DeleteByID(c, tx)
						break
					}
				case dao.LoadTypeGrpc:
					{
						grpc := dao.ServiceGrpcRule{
							ServiceID: serviceInfo.ID,
						}
						err = grpc.DeleteByID(c, tx)
						break
					}
				case dao.LoadTypeTcp:
					{
						tcp := dao.ServiceTCPRule{
							ServiceID: serviceInfo.ID,
						}
						err = tcp.DeleteByID(c, tx)
						break
					}
				}
				if err != nil {
					return
				}
				slb := &dao.ServiceLoadBalance{
					ServiceID: serviceInfo.ID,
				}
				err = slb.DeleteByID(c, tx)
				if err != nil {
					return
				}

				sac := &dao.ServiceAccessControl{
					ServiceID: serviceInfo.ID,
				}
				err = sac.DeleteByID(c, tx)
				if err != nil {
					return
				}

				err = serviceInfo.DeleteByID(c, tx)
				if err != nil {
					return
				}
				return
			})
		return
	}

}

type GetServiceDetailInput struct {
	ServiceID uint `json:"service_id" form:"service_id"`
}

type GetServiceDetailForHttpOutput struct {
	UpdateHttpServiceInput
}

type GetServiceDetailForTcpOutput struct {
	UpdateTcpServiceInput
}

type GetServiceDetailForGrpcOutput struct {
	UpdateGrpcServiceInput
}

func (p *GetServiceDetailInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = p
	return
}

func (p *GetServiceDetailInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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

func (p *GetServiceDetailInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	//o, ok := outIn.(*dao.ServiceDetail)
	//if ok {
	//	esi := EditServiceInfo{
	//		ServiceName: o.ServiceName,
	//		ServiceDesc: o.ServiceDesc,
	//	}
	//	eac := EditServiceAccessControlRule{
	//		OpenAuth:          o.OpenAuth,
	//		BlackList:         o.BlackList,
	//		WhiteList:         o.WhiteList,
	//		WhiteHostName:     o.WhiteHostName,
	//		ClientipFlowLimit: o.ClientipFlowLimit,
	//		ServiceFlowLimit:  o.ServiceFlowLimit,
	//	}
	//	elb := EditServiceLoadBalance{
	//		CheckMethod:            o.CheckMethod,
	//		CheckTimeout:           o.CheckTimeout,
	//		CheckInterval:          o.CheckInterval,
	//		RoundType:              o.RoundType,
	//		IpList:                 o.IPList,
	//		WeightList:             o.WeightList,
	//		ForbidLIst:             o.ForbidList,
	//		UpstreamConnectTimeout: o.UpstreamConnectTimeout,
	//		UpstreamHeaderTimeout:  o.UpstreamHeaderTimeout,
	//		UpstreamIdleTimeout:    o.UpstreamIDleTimeout,
	//		UpstreamMaxIdle:        o.UpstreamMaxIDle,
	//	}
	//	switch o.LoadType {
	//	case dao.LoadTypeHttp:
	//		{
	//			eh := EditServiceHTTPRule{
	//				RuleType:       o.RuleType,
	//				Rule:           o.Rule,
	//				NeedHttps:      o.NeedHTTPs,
	//				NeedStripUri:   o.NeedStripURI,
	//				NeedWebSocket:  o.NeedWebsocket,
	//				UrlRewrite:     o.URLRewrite,
	//				HeaderTransfor: o.ServiceHTTPRuleExceptModel.HeaderTransfor,
	//			}
	//			return GetServiceDetailForHttpOutput{
	//				UpdateHttpServiceInput{
	//					ServiceID: o.ID,
	//					AddHttpServiceInput: AddHttpServiceInput{
	//						EditServiceInfo:              esi,
	//						EditServiceHTTPRule:          eh,
	//						EditServiceAccessControlRule: eac,
	//						EditServiceLoadBalance:       elb,
	//					},
	//				},
	//			}
	//		}
	//	case dao.LoadTypeTcp:
	//		{
	//			et := EditServiceTCPRule{
	//				Port: o.ServiceTCPRuleExceptModel.Port,
	//			}
	//			return GetServiceDetailForHttpOutput{
	//				UpdateHttpServiceInput{
	//					ServiceID: o.ID,
	//					AddHttpServiceInput: AddHttpServiceInput{
	//						EditServiceInfo:              esi,
	//						EditServiceHTTPRule:          eh,
	//						EditServiceAccessControlRule: eac,
	//						EditServiceLoadBalance:       elb,
	//					},
	//				},
	//			}
	//		}
	//	case dao.LoadTypeGrpc:
	//		{
	//			eg := EditServiceGRPCRule{
	//				Port:           o.ServiceGrpcRuleExceptModel.Port,
	//				HeaderTransfor: o.ServiceGrpcRuleExceptModel.HeaderTransfor,
	//			}
	//			return GetServiceDetailForGrpcOutput{
	//				UpdateGrpcServiceInput{
	//					ServiceID: o.ID,
	//					AddHttpServiceInput: AddHttpServiceInput{
	//						EditServiceInfo:              esi,
	//						EditServiceHTTPRule:          eh,
	//						EditServiceAccessControlRule: eac,
	//						EditServiceLoadBalance:       elb,
	//					},
	//				},
	//			}
	//		}
	//	}
	//}
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (p *GetServiceDetailInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		params, err := handle(c)
		if err != nil {
			return
		}
		p = params.(*GetServiceDetailInput)
		db := lib.GetDefaultDB()
		info := &dao.ServiceInfo{
			ID: p.ServiceID,
		}
		out, err = info.FindOneServiceDetail(c, db)
		return
	}
}
