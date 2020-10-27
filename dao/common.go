package dao

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
)

//const (
//	Load_HTTP = 0
//	Load_TCP  = 1
//	Load_GRPC = 2
//
//	HttpRuleTypePrefixURL = 0
//	HttpRuleTypeDomain    = 1
//)

type PageSize struct {
	Size int
	No   int
	Info string
}

func GetFields(p interface{}) []string {
	t := reflect.TypeOf(p).Elem()
	all := []string{}
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == reflect.TypeOf(gorm.Model{}) {
			for j := 0; j < t.Field(i).Type.NumField(); j++ {
				all = append(all, t.Field(i).Type.Field(j).Name)
			}
			continue
		}
		all = append(all, t.Field(i).Name)
	}
	return all
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func ErrorHandleForDB(res *gorm.DB) (err error) {
	if res.Error != nil {
		err = res.Error
		return
	}
	if res.RowsAffected == 0 {
		err = errors.New("not updated")
		return
	}
	return
}

type ServiceHTTPRuleExceptModel struct {
	RuleType        HttpRuleType        `json:"rule_type"`        // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule            string              `json:"rule"`             // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHTTPs       NeedHttpsType       `json:"need_https"`       // 支持https 1=支持
	NeedStripURI    NeedStripUriType    `json:"need_strip_uri"`   // 启用strip_uri 1=启用
	NeedWebsocket   NeedWebsocketType   `json:"need_websocket"`   // 是否支持websocket 1=支持
	URLRewrite      URLRewriteType      `json:"url_rewrite"`      // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransform HeaderTransformType `json:"header_transform"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type ServiceTCPRuleExceptModel struct {
	Port int `json:"port"` // 端口号
}

type ServiceGRPCRuleExceptModel struct {
	Port              int                 `json:"port"`               // 端口
	MetadataTransform HeaderTransformType `json:"metadata_transform"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type ServiceInfoExceptModel struct {
	ID          uint     `gorm:"primarykey" json:"id"`
	LoadType    LoadType `json:"load_type"`    // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string   `json:"service_name"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string   `json:"service_desc"` // 服务描述
}

type ServiceAccessControlExceptModel struct {
	OpenAuth          OpenAuthType `json:"open_auth"`           // 是否开启权限 1=开启
	BlackList         IpListType   `json:"black_list"`          // 黑名单ip
	WhiteList         IpListType   `json:"white_list"`          // 白名单ip
	WhiteHostName     string       `json:"white_host_name"`     // 白名单主机
	ClientipFlowLimit int          `json:"clientip_flow_limit"` // BS客户端ip限流
	ServiceFlowLimit  int          `json:"service_flow_limit"`  // 服务端限流
}

type ServiceLoadBalanceExceptModel struct {
	CheckMethod            int8           `json:"check_method"`             // 检查方法 0=tcpchk,检测端口是否握手成功
	CheckTimeout           int            `json:"check_timeout"`            // check超时时间,单位s
	CheckInterval          int            `json:"check_interval"`           // 检查间隔, 单位s
	RoundType              RoundType      `json:"round_type"`               // 轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IPList                 IpListType     `json:"ip_list"`                  // ip列表
	WeightList             WeightListType `json:"weight_list"`              // 权重列表
	ForbidList             IpListType     `json:"forbid_list"`              // 禁用ip列表
	UpstreamConnectTimeout int            `json:"upstream_connect_timeout"` // 建立连接超时, 单位s
	UpstreamHeaderTimeout  int            `json:"upstream_header_timeout"`  // 获取header超时, 单位s
	UpstreamIDleTimeout    int            `json:"upstream_idle_timeout"`    // 链接最大空闲时间, 单位s
	UpstreamMaxIDle        int            `json:"upstream_max_idle"`        // 最大空闲链接数
}
