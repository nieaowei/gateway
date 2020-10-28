package dto

import (
	"errors"
	"gateway/dao"
	"gateway/lib"
	"gateway/proxy/manager"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type GetServiceListInput struct {
	Info     string `json:"info" form:"info"`
	PageNo   int    `json:"page_no" form:"page_no" example:"2" validate:"required,min=1"`
	PageSize int    `json:"page_size" form:"page_size" example:"10" validate:"required,min=1"`
}

type ServiceListItem struct {
	ID          uint   `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType    string `json:"load_type" form:"load_type"`
	Address     string `json:"address" form:"Address"`
	Qps         int64  `json:"qps" form:"qps"`
	Qpd         int64  `json:"qpd" form:"qpd"`
	TotalNode   uint   `json:"total_node" form:"total_node"`
}

type GetServiceListOutput struct {
	Total int64             `json:"total" form:"total" validate:""`
	List  []ServiceListItem `json:"list" form:"list" validate:""`
}

func (p *GetServiceListInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, p)
	params = *p
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
		data, err := handle(c)
		if err != nil {
			return
		}
		params := data.(GetServiceListInput)
		serviceInfo := &dao.ServiceInfo{}
		db := lib.GetDefaultDB()
		serviceInfos, count, err := serviceInfo.PageListIdDesc(c, db, &dao.PageSize{
			Size: params.PageSize,
			No:   params.PageNo,
			Info: params.Info,
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
			//serviceDetail, err := info.FindOneServiceDetail(c, db)
			//if err != nil {
			//	return nil, err
			//}
			serviceAddr := ""
			loadType := ""

			type Port struct {
				Port int
			}

			type ServiceTypeInfo struct {
				RuleType  dao.HttpRuleType
				Rule      string
				NeedHTTPs int8
			}

			switch info.LoadType {
			case dao.Load_HTTP:
				{
					loadType = "HTTP"
					service := &dao.ServiceHTTPRule{
						ServiceID: info.ID,
					}
					res := &ServiceTypeInfo{}
					err = service.FindOneScan(c, db, res)
					if err != nil {
						return nil, err
					}
					//service := serviceDetail.ServiceHTTPRuleExceptModel
					if res.RuleType == dao.HttpRule_PrefixURL && res.NeedHTTPs == 0 {
						serviceAddr = clusterIP + ":" + clusterPort + res.Rule

					}
					if res.RuleType == dao.HttpRule_PrefixURL && res.NeedHTTPs == 1 {
						serviceAddr = clusterIP + ":" + clusterSSLPort + res.Rule
					}
					if res.RuleType == dao.HttpRule_Domain {
						serviceAddr = service.Rule
					}
					break
				}
			case dao.Load_TCP:
				{
					loadType = "TCP"
					//service := serviceDetail.ServiceTCPRuleExceptModel
					service := &dao.ServiceTCPRule{
						ServiceID: info.ID,
					}
					res := &Port{}
					err = service.FindOneScan(c, db, res)
					if err != nil {
						return nil, err
					}
					serviceAddr = clusterIP + ":" + strconv.Itoa(res.Port)
					break
				}
			case dao.Load_GRPC:
				{

					loadType = "GRPC"
					//service := serviceDetail.ServiceGRPCRuleExceptModel
					service := &dao.ServiceGrpcRule{
						ServiceID: info.ID,
					}
					res := &Port{}
					err = service.FindOneScan(c, db, res)
					if err != nil {
						return nil, err
					}
					serviceAddr = clusterIP + ":" + strconv.Itoa(res.Port)
					break
				}
			}
			type IpListRes struct {
				IpList dao.IpListType
			}
			//ipList := serviceDetail.ServiceLoadBalanceExceptModel.GetIPListByModel()
			res := &IpListRes{}
			serviceLoadBalance := &dao.ServiceLoadBalance{
				ServiceID: info.ID,
			}
			err := serviceLoadBalance.FindOneScan(c, db, res)
			if err != nil {
				return nil, err
			}
			counter, ok := manager.Default().GetRedisService(manager.RedisServicePrefix + info.ServiceName)
			var qpd, qps int64
			if ok {
				impl := counter.(*manager.RedisFlowCountService)
				qpd, _ = impl.GetDayData(time.Now())
				qps = impl.QPS
			}

			item := ServiceListItem{
				ID:          info.ID,
				ServiceName: info.ServiceName,
				ServiceDesc: info.ServiceDesc,
				LoadType:    loadType,
				Address:     serviceAddr,
				Qps:         qps,
				Qpd:         qpd,
				TotalNode:   uint(len(res.IpList.GetSlice())),
			}
			outE.List = append(outE.List, item)
		}
		out = outE
		return
	}

}

type DeleteServiceInput struct {
	ID uint `json:"id" form:"id" example:"96" validate:"required"`
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
				case dao.Load_HTTP:
					{
						http := dao.ServiceHTTPRule{
							ServiceID: serviceInfo.ID,
						}
						err = http.DeleteByID(c, tx)
						break
					}
				case dao.Load_GRPC:
					{
						grpc := dao.ServiceGrpcRule{
							ServiceID: serviceInfo.ID,
						}
						err = grpc.DeleteByID(c, tx)
						break
					}
				case dao.Load_TCP:
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
	ServiceID uint `json:"service_id" form:"service_id" example:"133" validate:"required"`
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

	return func(c *gin.Context) (out interface{}, err error) {
		data, err := handle(c)
		if err != nil {
			return
		}
		o, ok := data.(*dao.ServiceDetail)
		if ok {
			esi := EditServiceInfo{
				ServiceName: o.ServiceName,
				ServiceDesc: o.ServiceDesc,
			}
			eac := EditServiceAccessControlRule{
				OpenAuth:          o.OpenAuth,
				BlackList:         o.BlackList,
				WhiteList:         o.WhiteList,
				WhiteHostName:     o.WhiteHostName,
				ClientipFlowLimit: o.ClientipFlowLimit,
				ServiceFlowLimit:  o.ServiceFlowLimit,
			}
			elb := EditServiceLoadBalance{
				RoundType:              o.RoundType,
				IpList:                 o.IPList,
				WeightList:             o.WeightList,
				ForbidList:             o.ForbidList,
				UpstreamConnectTimeout: o.UpstreamConnectTimeout,
				UpstreamHeaderTimeout:  o.UpstreamHeaderTimeout,
				UpstreamIdleTimeout:    o.UpstreamIDleTimeout,
				UpstreamMaxIdle:        o.UpstreamMaxIDle,
			}
			switch item := o.Rule.(type) {
			case *dao.ServiceHTTPRuleExceptModel:
				{
					eh := EditServiceHTTPRule{
						RuleType:        item.RuleType,
						Rule:            item.Rule,
						NeedHttps:       item.NeedHTTPs,
						NeedStripUri:    item.NeedStripURI,
						NeedWebsocket:   item.NeedWebsocket,
						UrlRewrite:      item.URLRewrite,
						HeaderTransform: item.HeaderTransform,
					}
					return GetServiceDetailForHttpOutput{
						UpdateHttpServiceInput{
							ServiceID: o.ID,
							AddHttpServiceInput: AddHttpServiceInput{
								EditServiceInfo:              esi,
								EditServiceHTTPRule:          eh,
								EditServiceAccessControlRule: eac,
								EditServiceLoadBalance:       elb,
							},
						},
					}, nil
				}
			case *dao.ServiceTCPRuleExceptModel:
				{
					et := EditServiceTCPRule{
						Port: item.Port,
					}
					return GetServiceDetailForTcpOutput{
						UpdateTcpServiceInput{
							ServiceID: o.ID,
							AddTcpServiceInput: AddTcpServiceInput{
								EditServiceInfo:              esi,
								EditServiceTCPRule:           et,
								EditServiceAccessControlRule: eac,
								EditServiceLoadBalance:       elb,
							},
						},
					}, nil
				}

			case *dao.ServiceGRPCRuleExceptModel:
				{
					eg := EditServiceGRPCRule{
						Port:              item.Port,
						MetadataTransform: item.MetadataTransform,
					}
					return GetServiceDetailForGrpcOutput{
						UpdateGrpcServiceInput{
							ServiceID: o.ID,
							AddGrpcServiceInput: AddGrpcServiceInput{
								EditServiceInfo:              esi,
								EditServiceGRPCRule:          eg,
								EditServiceAccessControlRule: eac,
								EditServiceLoadBalance:       elb,
							},
						},
					}, nil
				}
			}
			//switch o.LoadType {
			//case dao.Load_HTTP:
			//	{
			//		eh := EditServiceHTTPRule{
			//			RuleType:        o.RuleType,
			//			Rule:            o.Rule,
			//			NeedHttps:       o.NeedHTTPs,
			//			NeedStripUri:    o.NeedStripURI,
			//			NeedWebsocket:   o.NeedWebsocket,
			//			UrlRewrite:      o.URLRewrite,
			//			HeaderTransform: o.ServiceHTTPRuleExceptModel.HeaderTransform,
			//		}
			//		return GetServiceDetailForHttpOutput{
			//			UpdateHttpServiceInput{
			//				ServiceID: o.ID,
			//				AddHttpServiceInput: AddHttpServiceInput{
			//					EditServiceInfo:              esi,
			//					EditServiceHTTPRule:          eh,
			//					EditServiceAccessControlRule: eac,
			//					EditServiceLoadBalance:       elb,
			//				},
			//			},
			//		}, nil
			//	}
			//case dao.Load_TCP:
			//	{
			//		et := EditServiceTCPRule{
			//			Port: o.ServiceTCPRuleExceptModel.Port,
			//		}
			//		return GetServiceDetailForTcpOutput{
			//			UpdateTcpServiceInput{
			//				ServiceID: o.ID,
			//				AddTcpServiceInput: AddTcpServiceInput{
			//					EditServiceInfo:              esi,
			//					EditServiceTCPRule:           et,
			//					EditServiceAccessControlRule: eac,
			//					EditServiceLoadBalance:       elb,
			//				},
			//			},
			//		}, nil
			//	}
			//case dao.Load_GRPC:
			//	{
			//		eg := EditServiceGRPCRule{
			//			Port:              o.ServiceGRPCRuleExceptModel.Port,
			//			MetadataTransform: o.ServiceGRPCRuleExceptModel.MetadataTransform,
			//		}
			//		return GetServiceDetailForGrpcOutput{
			//			UpdateGrpcServiceInput{
			//				ServiceID: o.ID,
			//				AddGrpcServiceInput: AddGrpcServiceInput{
			//					EditServiceInfo:              esi,
			//					EditServiceGRPCRule:          eg,
			//					EditServiceAccessControlRule: eac,
			//					EditServiceLoadBalance:       elb,
			//				},
			//			},
			//		}, nil
			//	}
			//}
		}
		return nil, errors.New("data handle error")
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

type GetServiceStatInput struct {
	ServiceID uint `json:"service_id" form:"service_id" example:"156"`
}

type GetServiceStatOutput struct {
	TodayList     []int64 `json:"today_list"`
	YesterdayList []int64 `json:"yesterday_list"`
}

func (g *GetServiceStatInput) BindValidParam(c *gin.Context) (params interface{}, err error) {
	err = public.DefaultGetValidParams(c, g)
	params = g
	return
}

func (g *GetServiceStatInput) ExecHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		inter, err := handle(c)
		serviceName := manager.RedisTotalPrefix
		if err != nil {
			serviceName = manager.RedisTotalPrefix
		}
		params := inter.(*GetServiceStatInput)
		db := lib.GetDefaultDB()
		type Select struct {
			ServiceName string
		}
		s := &Select{}
		err = (&dao.ServiceInfo{ID: params.ServiceID}).FindOneScan(c, db, s)
		if err == nil {
			serviceName = manager.RedisServicePrefix + s.ServiceName
		}

		data := &GetServiceStatOutput{}
		redisService, ok := manager.Default().GetRedisService(serviceName)
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

func (g *GetServiceStatInput) OutputHandle(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		return handle(c)
	}
}

func (g *GetServiceStatInput) ErrorHandle(handle FunctionalHandle) func(c *gin.Context) {
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
