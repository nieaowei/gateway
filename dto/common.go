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
	ServiceName string `json:"service_name" validate:"required,min=6,max=255"`
	ServiceDesc string `json:"service_desc" validate:"required,min=6,max=255"`
}

type EditServiceAccessControlRule struct {
	OpenAuth          int8   `json:"open_auth" validate:"oneof=0 1"`
	BlackList         string `json:"black_list" validate:"min=0,max=1000"`
	WhiteList         string `json:"white_list" validate:"min=0,max=1000"`
	WhiteHostName     string `json:"white_host_name" validate:"min=0,max=1000"`
	ClientipFlowLimit int    `json:"clientip_flow_limit" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" validate:"min=0"`
}

type EditServiceLoadBalance struct {
	CheckMethod            int8   `json:"check_method" validate:"oneof=0 1"`
	CheckTimeout           int    `json:"check_timeout" validate:"min=0"`
	CheckInterval          int    `json:"check_interval" validate:"min=0"`
	RoundType              int8   `json:"round_type" validate:"oneof=0 1 2 3"`
	IpList                 string `json:"ip_list" validate:"min=0,max=2000,valid_ip_list"`
	WeightList             string `json:"weight_list" validate:"min=0,max=2000,valid_weight_list"`
	ForbidLIst             string `json:"forbid_l_ist" validate:"min=0,max=2000"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" validate:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" validate:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" validate:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" validate:"min=0"`
}

type EditServiceHTTPRule struct {
	RuleType        int8   `json:"rule_type" validate:"oneof=0 1 2"`
	Rule            string `json:"rule" validate:"min=0,max=255"`
	NeedHttps       int8   `json:"need_https" validate:"oneof=0 1"`
	NeedStripUri    int8   `json:"need_strip_uri" validate:"oneof=0 1"`
	NeedWebSocket   int8   `json:"need_web_socket" validate:"oneof=0 1"`
	UrlRewrite      string `json:"url_rewrite" validate:"min=0,max=5000"`
	HeaderTransform string `json:"header_transform" validate:"min=0,max=5000,valid_header_transform" `
}

type EditServiceGRPCRule struct {
	Port            int    `json:"port" validate:"min=0,max=65535"`
	HeaderTransform string `json:"header_transform" validate:"min=0,max=5000,valid_header_transform"`
}

type EditServiceTCPRule struct {
	Port int `json:"port" validate:"min=0,max=65535"`
}
