package dto

import (
	"encoding/gob"
	"errors"
	"gateway/dao"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
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
	ServiceName string `json:"service_name" example:"service_test" validate:"required,min=6,max=255" example:"servicename_1234567"`
	ServiceDesc string `json:"service_desc" example:"service_test" validate:"required,max=255" example:"sevicedesc_1234567"`
}

type EditServiceAccessControlRule struct {
	OpenAuth          dao.OpenAuthType `json:"open_auth" validate:"oneof=0 1" example:"0"`
	BlackList         dao.IpListType   `json:"black_list"  example:"192.168.1.0\n122.12.12.3" validate:"min=0,max=1000"`
	WhiteList         dao.IpListType   `json:"white_list"  example:"172.17.12.1" validate:"min=0,max=1000"`
	WhiteHostName     string           `json:"white_host_name" example:"nekilc.com" validate:"min=0,max=1000"`
	ClientipFlowLimit int              `json:"clientip_flow_limit" example:"23" validate:"min=0"`
	ServiceFlowLimit  int              `json:"service_flow_limit" example:"12" validate:"min=0"`
}

type EditServiceLoadBalance struct {
	RoundType              dao.RoundType      `json:"round_type" example:"1" validate:"oneof=0 1 2 3"`
	IpList                 dao.IpListType     `json:"ip_list" example:"172.1.1.1:80\n172.11.1.2:87" validate:"required,min=0,max=2000,valid_ip_list"`
	WeightList             dao.WeightListType `json:"weight_list" example:"1\n2" validate:"required,min=0,max=2000,valid_weight_list"`
	ForbidList             dao.IpListType     `json:"forbid_list" validate:"min=0,max=2000"`
	UpstreamConnectTimeout int                `json:"upstream_connect_timeout" example:"122" validate:"min=0"`
	UpstreamHeaderTimeout  int                `json:"upstream_header_timeout" example:"322" validate:"min=0"`
	UpstreamIdleTimeout    int                `json:"upstream_idle_timeout" example:"321" validate:"min=0"`
	UpstreamMaxIdle        int                `json:"upstream_max_idle" example:"12" validate:"min=0"`
}

type EditServiceHTTPRule struct {
	RuleType        dao.HttpRuleType        `json:"rule_type" example:"1" validate:"oneof=0 1"`
	Rule            string                  `json:"rule" example:"/dsads" validate:"required,max=255"`
	NeedHttps       dao.NeedHttpsType       `json:"need_https"  example:"1" validate:"oneof=0 1"`
	NeedStripUri    dao.NeedStripUriType    `json:"need_strip_uri" example:"1" validate:"oneof=0 1"`
	NeedWebsocket   dao.NeedWebsocketType   `json:"need_websocket" example:"1" validate:"oneof=0 1"`
	UrlRewrite      dao.URLRewriteType      `json:"url_rewrite" example:"add w\ndel 1" validate:"valid_url_rewrite,min=0,max=5000"`
	HeaderTransform dao.HeaderTransformType `json:"header_transform" example:"add a 12\nadd b 13" validate:"min=0,max=5000,valid_header_transform" `
}

type EditServiceGRPCRule struct {
	Port              int                     `json:"port"  example:"7777" validate:"min=0,max=65535"`
	MetadataTransform dao.HeaderTransformType `json:"metadata_transform" example:"add a 32" validate:"min=0,max=5000,valid_header_transform"`
}

type EditServiceTCPRule struct {
	Port int `json:"port" example:"9999" validate:"min=0,max=65535"`
}

func IpListAndWeightListNumValid(handle FunctionalHandle) FunctionalHandle {
	return func(c *gin.Context) (out interface{}, err error) {
		data, err := handle(c)
		if err != nil {
			return
		}
		e := reflect.ValueOf(data).Elem()
		if dao.RoundType(e.FieldByName("RoundType").Int()) == dao.Round_WeightRound {
			ips := e.FieldByName("IpList").String()
			weight := e.FieldByName("WeightList").String()
			if len(strings.Split(ips, "\n")) != len(strings.Split(weight, "\n")) {
				return nil, errors.New("IP列表数量和权重数量不匹配")
			}
		}
		return data, nil
	}
}
