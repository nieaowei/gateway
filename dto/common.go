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
	ServiceName string `json:"service_name"`
	ServiceDesc string `json:"service_desc"`
}

type EditServiceAccessControlRule struct {
	OpenAuth          int8   `json:"open_auth"`
	BlackList         string `json:"black_list"`
	WhiteList         string `json:"white_list"`
	WhiteHostName     string `json:"white_host_name"`
	ClientipFlowLimit int    `json:"clientip_flow_limit"`
	ServiceFlowLimit  int    `json:"service_flow_limit"`
}

type EditServiceLoadBalance struct {
	CheckMethod            int8   `json:"check_method"`
	CheckTimeout           int    `json:"check_timeout"`
	CheckInterval          int    `json:"check_interval"`
	RoundType              int8   `json:"round_type"`
	IpList                 string `json:"ip_list"`
	WeightList             string `json:"weight_list"`
	ForbidLIst             string `json:"forbid_l_ist"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle"`
}

type EditServiceHTTPRule struct {
	RuleType       int8   `json:"rule_type"`
	Rule           string `json:"rule"`
	NeedHttps      int8   `json:"need_https"`
	NeedStripUri   int8   `json:"need_strip_uri"`
	NeedWebSocket  int8   `json:"need_web_socket"`
	UrlRewrite     string `json:"url_rewrite"`
	HeaderTransfor string `json:"header_transfor"`
}

type EditServiceGRPCRule struct {
	Port           int    `json:"port"`
	HeaderTransfor string `json:"header_transfor"`
}

type EditServiceTCPRule struct {
	Port int `json:"port"`
}
