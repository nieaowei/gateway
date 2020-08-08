package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddGrpcServiceInput struct {
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
	dao.ServiceGrpcRule
	dao.ServiceAccessControl
	dao.ServiceLoadBalance
}

func (p *AddGrpcServiceInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *AddGrpcServiceInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 10002, err)
}

func (p *AddGrpcServiceInput) OutputHandle(c *gin.Context, outIn interface{}) (out interface{}) {
	return outIn
}

func (p *AddGrpcServiceInput) Exec(c *gin.Context) (out interface{}, err error) {
	db := lib.GetDefaultDB()
	// set Grpc type
	p.LoadType = dao.LoadTypeGrpc
	// start
	err = db.Transaction(
		func(tx *gorm.DB) (err error) {
			err = p.ServiceInfo.InsertAfterCheck(c, tx, true)
			if err != nil {
				return
			}
			p.ServiceGrpcRule.ServiceID = p.ServiceInfo.ID
			err = p.ServiceGrpcRule.InsertAfterCheck(c, tx, true)
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
	out = p
	return
}

type UpdateGrpcServiceInput struct {
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
	dao.ServiceGrpcRule
	dao.ServiceAccessControl
	dao.ServiceLoadBalance
}

func (p *UpdateGrpcServiceInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *UpdateGrpcServiceInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

func (p *UpdateGrpcServiceInput) Exec(c *gin.Context) (out interface{}, err error) {
	db := lib.GetDefaultDB()
	// set Grpc type
	p.LoadType = dao.LoadTypeGrpc
	// start
	err = db.Transaction(
		func(tx *gorm.DB) (err error) {
			p.ServiceInfo.ID = p.ServiceID
			err = p.ServiceInfo.UpdateAll(c, tx)
			if err != nil {
				return
			}
			p.ServiceGrpcRule.ServiceID = p.ServiceInfo.ID
			err = p.ServiceGrpcRule.UpdateAllByServiceID(c, tx)
			if err != nil {
				return
			}
			p.ServiceLoadBalance.ServiceID = p.ServiceInfo.ID
			err = p.ServiceLoadBalance.UpdateAllByServiceID(c, tx)
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

func (p *UpdateGrpcServiceInput) OutputHandle(c *gin.Context, outIn interface{}) (out interface{}) {
	return outIn
}
