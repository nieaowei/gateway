package dto

import (
	"encoding/gob"
	"gateway/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(&dao.AdminSessionInfo{})
	//gob.Register(&GetServiceDetailOutput{})
}

//@弃用
type FunctionalHandle func(c *gin.Context) (out interface{}, err error)

type Functional interface {
	BindValidParam(c *gin.Context) (params interface{}, err error)
	ExecHandle(handle FunctionalHandle) FunctionalHandle
	OutputHandle(handle FunctionalHandle) FunctionalHandle
	ErrorHandle(handle FunctionalHandle) func(c *gin.Context)
}

type EditServiceInfo struct {
	ServiceName string `json:"service_name" example:"service_test" validate:"required,min=6,max=255"`
	ServiceDesc string `json:"service_desc" example:"service_test" validate:"required,max=255"`
}

type EditServiceAccessControlRule struct {
	OpenAuth          int8   `json:"open_auth" validate:"oneof=0 1"`
	BlackList         string `json:"black_list"  example:"192.168.1.0\n122.12.12.3" validate:"min=0,max=1000"`
	WhiteList         string `json:"white_list"  example:"172.17.12.1" validate:"min=0,max=1000"`
	WhiteHostName     string `json:"white_host_name" example:"nekilc.com" validate:"min=0,max=1000"`
	ClientipFlowLimit int    `json:"clientip_flow_limit" example:"23" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" example:"12" validate:"min=0"`
}

type EditServiceLoadBalance struct {
	RoundType              int8   `json:"round_type" example:"1" validate:"oneof=0 1 2 3"`
	IpList                 string `json:"ip_list" example:"172.1.1.1:80\n172.11.1.2:87" validate:"required,min=0,max=2000,valid_ip_list"`
	WeightList             string `json:"weight_list" example:"1\n2" validate:"required,min=0,max=2000,valid_weight_list"`
	ForbidLIst             string `json:"forbid_l_ist" validate:"min=0,max=2000"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" example:"122" validate:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" example:"322" validate:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" example:"321" validate:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" example:"12" validate:"min=0"`
}

type EditServiceHTTPRule struct {
	RuleType        int8   `json:"rule_type" example:"1" validate:"oneof=0 1"`
	Rule            string `json:"rule" example:"/dsads" validate:"required,min=0,max=255"`
	NeedHttps       int8   `json:"need_https"  example:"1" validate:"oneof=0 1"`
	NeedStripUri    int8   `json:"need_strip_uri" example:"1" validate:"oneof=0 1"`
	NeedWebSocket   int8   `json:"need_web_socket" example:"1" validate:"oneof=0 1"`
	UrlRewrite      string `json:"url_rewrite" example:"add w\ndel 1" validate:"valid_url_rewrite,min=0,max=5000"`
	HeaderTransform string `json:"header_transform" example:"add a 12\nadd b 13" validate:"min=0,max=5000,valid_header_transform" `
}

type EditServiceGRPCRule struct {
	Port            int    `json:"port"  example:"7777" validate:"min=0,max=65535"`
	HeaderTransform string `json:"header_transform" example:"add a 32" validate:"min=0,max=5000,valid_header_transform"`
}

type EditServiceTCPRule struct {
	Port int `json:"port" example:"9999" validate:"min=0,max=65535"`
}
