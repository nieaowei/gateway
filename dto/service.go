package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ServiceListInput struct {
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

type ServiceListOutput struct {
	Total int64             `json:"total" form:"total" validate:""`
	List  []ServiceListItem `json:"list" form:"list" validate:""`
}

func (p *ServiceListInput) GetServiceList(c *gin.Context) (out *ServiceListOutput, err error) {
	serviceInfo := &dao.ServiceInfo{}
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	serviceInfos, count, err := serviceInfo.PageList(c, db, &dao.PageSize{
		Size: p.PageSize,
		No:   p.PageNo,
		Info: p.Info,
	})
	if err != nil {
		return
	}
	out = &ServiceListOutput{
		Total: count,
		List:  []ServiceListItem{},
	}
	clusterIP := lib.GetDefaultConfBase().Cluster.Ip
	clusterPort := lib.GetDefaultConfBase().Cluster.Port
	clusterSSLPort := lib.GetDefaultConfBase().Cluster.SslPort
	for _, info := range serviceInfos {
		serviceDetail, err := info.FindOneServiceDetail(c, db)
		if err != nil {
			return nil, err
		}
		serviceAddr := ""
		loadType := ""
		switch serviceDetail.Info.LoadType {
		case dao.LoadTypeHttp:
			{
				loadType = "HTTP"
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHTTPs == 0 {
					serviceAddr = clusterIP + ":" + clusterPort + serviceDetail.HTTP.Rule

				}
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHTTPs == 1 {
					serviceAddr = clusterIP + ":" + clusterSSLPort + serviceDetail.HTTP.Rule
				}
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypeDomain {
					serviceAddr = serviceDetail.HTTP.Rule
				}
				break
			}
		case dao.LoadTypeTcp:
			{
				loadType = "TCP"
				serviceAddr = clusterIP + ":" + strconv.Itoa(int(serviceDetail.TCP.Port))
				break
			}
		case dao.LoadTypeGrpc:
			{
				loadType = "GRPC"
				serviceAddr = clusterIP + ":" + strconv.Itoa(int(serviceDetail.GRPC.Port))
				break
			}
		}
		ipList := serviceDetail.LoadBalance.GetIPListByModel()

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
		out.List = append(out.List, item)
	}

	return
}

func (p *ServiceListInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

type ServiceDeleteInput struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

func (p *ServiceDeleteInput) Delete(c *gin.Context) (err error) {
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	serviceInfo := &dao.ServiceInfo{
		Model: gorm.Model{
			ID: p.ID,
		},
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
	return err
}

func (p *ServiceDeleteInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

type ServiceAddHttpInput struct {
	//LoadType               uint   `json:"load_type"`
	//ServiceName            string `json:"service_name"`
	//ServiceDesc            string `json:"service_desc"`
	//RuleType               uint8  `json:"rule_type"`
	//Rule                   string `json:"rule"`
	//NeedHttps              uint8  `json:"need_https"`
	//NeedStripUri           uint8  `json:"need_strip_uri"`
	//NeedWebSocket          uint8  `json:"need_web_socket"`
	//UrlRewrite             string `json:"url_rewrite"`
	//HeaderTransfor         string `json:"header_transfor"`
	//OpenAuth               uint8  `json:"open_auth"`
	//BlackList              string `json:"black_list"`
	//WhiteList              string `json:"white_list"`
	//WhiteHostName          string `json:"white_host_name"`
	//ClientipFlowLimit      uint16 `json:"clientip_flow_limit"`
	//ServiceFlowLimit       uint16 `json:"service_flow_limit"`
	//CheckMethod            uint   `json:"check_method"`
	//CheckTimeout           uint   `json:"check_timeout"`
	//CheckInterval          uint   `json:"check_interval"`
	//RoundType              uint8  `json:"round_type"`
	//IpList                 string `json:"ip_list"`
	//WeightList             string `json:"weight_list"`
	//ForbidLIst             string `json:"forbid_l_ist"`
	//UpstreamConnectTimeout uint16 `json:"upstream_connect_timeout"`
	//UpstreamHeaderTimeout  uint16 `json:"upstream_header_timeout"`
	//UpstreamIdleTimeout    uint16 `json:"upstream_idle_timeout"`
	//UpstreamMaxIdle        uint16 `json:"upstream_max_idle"`
	dao.ServiceInfo
	dao.ServiceHTTPRule
	dao.ServiceAccessControl
	dao.ServiceLoadBalance
}

func (p *ServiceAddHttpInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *ServiceAddHttpInput) AddHttpService(c *gin.Context) (err error) {
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	// set http type
	p.LoadType = dao.LoadTypeHttp
	// start
	err = db.Transaction(
		func(tx *gorm.DB) (err error) {
			err = p.ServiceInfo.Insert(c, tx)
			if err != nil {
				return
			}
			p.ServiceHTTPRule.ServiceID = p.ServiceInfo.ID
			err = p.ServiceHTTPRule.InsertAfterCheck(c, tx, true)
			if err != nil {
				return
			}
			p.ServiceLoadBalance.ServiceID = p.ServiceInfo.ID
			err = p.ServiceLoadBalance.InsertAfterCheck(c, tx, true)
			if err != nil {
				return
			}
			p.ServiceAccessControl.ServiceID = p.ServiceInfo.ID
			err = p.ServiceAccessControl.InsertAfterCheck(c, tx, true)
			if err != nil {
				return
			}
			return
		})
	return
}

type HttpServiceUpdateInput struct {
	//LoadType               uint   `json:"load_type"`
	//ServiceDesc            string `json:"service_desc"`
	//RuleType               uint8  `json:"rule_type"`
	//Rule                   string `json:"rule"`
	//NeedHttps              uint8  `json:"need_https"`
	//NeedStripUri           uint8  `json:"need_strip_uri"`
	//NeedWebSocket          uint8  `json:"need_web_socket"`
	//UrlRewrite             string `json:"url_rewrite"`
	//HeaderTransfor         string `json:"header_transfor"`
	//OpenAuth               uint8  `json:"open_auth"`
	//BlackList              string `json:"black_list"`
	//WhiteList              string `json:"white_list"`
	//WhiteHostName          string `json:"white_host_name"`
	//ClientipFlowLimit      uint16 `json:"clientip_flow_limit"`
	//ServiceFlowLimit       uint16 `json:"service_flow_limit"`
	//CheckMethod            uint   `json:"check_method"`
	//CheckTimeout           uint   `json:"check_timeout"`
	//CheckInterval          uint   `json:"check_interval"`
	//RoundType              uint8  `json:"round_type"`
	//IpList                 string `json:"ip_list"`
	//WeightList             string `json:"weight_list"`
	//ForbidLIst             string `json:"forbid_l_ist"`
	//UpstreamConnectTimeout uint16 `json:"upstream_connect_timeout"`
	//UpstreamHeaderTimeout  uint16 `json:"upstream_header_timeout"`
	//UpstreamIdleTimeout    uint16 `json:"upstream_idle_timeout"`
	//UpstreamMaxIdle        uint16 `json:"upstream_max_idle"`
	ServiceID uint `json:"service_id"`
	dao.ServiceInfo
	dao.ServiceHTTPRule
	dao.ServiceAccessControl
	dao.ServiceLoadBalance
}

func (p *HttpServiceUpdateInput) UpdateHttpService(c *gin.Context) (err error) {
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	// set http type
	p.LoadType = dao.LoadTypeHttp
	// start
	err = db.Transaction(
		func(tx *gorm.DB) (err error) {
			p.ServiceInfo.ID = p.ServiceID
			err = p.ServiceInfo.UpdateAll(c, tx)
			if err != nil {
				return
			}
			p.ServiceHTTPRule.ServiceID = p.ServiceInfo.ID
			err = p.ServiceHTTPRule.UpdateAll(c, tx)
			if err != nil {
				return
			}
			p.ServiceLoadBalance.ServiceID = p.ServiceInfo.ID
			err = p.ServiceLoadBalance.UpdateAll(c, tx)
			if err != nil {
				return
			}
			p.ServiceAccessControl.ServiceID = p.ServiceInfo.ID
			err = p.ServiceAccessControl.UpdateAllByServiceID(c, tx)
			if err != nil {
				return
			}
			return
		})
	return
}

func (p *HttpServiceUpdateInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}
