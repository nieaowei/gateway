package manager

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/proxy/loadbalance"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type TCPService struct {
	dao.ServiceInfoExceptModel
	dao.ServiceTCPRuleExceptModel
	dao.ServiceLoadBalanceExceptModel
	dao.ServiceAccessControlExceptModel
}

type ServiceLoadBalance struct {
	CheckMethod            int8                 `json:"check_method"`             // 检查方法 0=tcpchk,检测端口是否握手成功
	CheckTimeout           int                  `json:"check_timeout"`            // check超时时间,单位s
	CheckInterval          int                  `json:"check_interval"`           // 检查间隔, 单位s
	RoundType              dao.RoundType        `json:"round_type"`               // 轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IPList                 []dao.IpListItem     `json:"ip_list"`                  // ip列表
	WeightList             []dao.WeightListItem `json:"weight_list"`              // 权重列表
	ForbidList             []dao.IpListItem     `json:"forbid_list"`              // 禁用ip列表
	UpstreamConnectTimeout int                  `json:"upstream_connect_timeout"` // 建立连接超时, 单位s
	UpstreamHeaderTimeout  int                  `json:"upstream_header_timeout"`  // 获取header超时, 单位s
	UpstreamIDleTimeout    int                  `json:"upstream_idle_timeout"`    // 链接最大空闲时间, 单位s
	UpstreamMaxIDle        int                  `json:"upstream_max_idle"`        // 最大空闲链接数
}

type HTTPServiceRule struct {
	RuleType        dao.HttpRuleType          `json:"rule_type"`        // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule            string                    `json:"rule"`             // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHTTPs       bool                      `json:"need_https"`       // 支持https 1=支持
	NeedStripURI    bool                      `json:"need_strip_uri"`   // 启用strip_uri 1=启用
	NeedWebsocket   bool                      `json:"need_websocket"`   // 是否支持websocket 1=支持
	URLRewrite      string                    `json:"url_rewrite"`      // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransform []dao.HeaderTransformItem `json:"header_transform"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type ServiceAccessControl struct {
	OpenAuth          bool             `json:"open_auth"`           // 是否开启权限 1=开启
	BlackList         []dao.IpListItem `json:"black_list"`          // 黑名单ip
	WhiteList         []dao.IpListItem `json:"white_list"`          // 白名单ip
	WhiteHostName     string           `json:"white_host_name"`     // 白名单主机
	ClientipFlowLimit int              `json:"clientip_flow_limit"` // BS客户端ip限流
	ServiceFlowLimit  int              `json:"service_flow_limit"`  // 服务端限流
}

type HTTPService struct {
	dao.ServiceInfoExceptModel
	//dao.ServiceHTTPRuleExceptModel
	HTTPServiceRule
	ServiceLoadBalance
	//dao.ServiceAccessControlExceptModel
	ServiceAccessControl
}

type GRPCService struct {
	dao.ServiceInfoExceptModel
	dao.ServiceGRPCRuleExceptModel
	dao.ServiceLoadBalanceExceptModel
	dao.ServiceAccessControlExceptModel
}

const (
	ServicePrefix = "service_"
	TotalPrefix   = "total_"
)

type ServiceMgr struct {
	//TCPServiceList  []TCPService
	//HTTPServiceList []HTTPService
	//GRPCServiceList []GRPCService
	GRPCServiceMap  lib.SafeMap //GRPCService
	TCPServiceMap   lib.SafeMap //TCPService
	HTTPServiceMap  lib.SafeMap //HTTPService
	loadbalanceMap  lib.SafeMap //loadbalance.LoadBalancer
	redisServiceMap lib.SafeMap //lib.RedisService
	transportMap    lib.SafeMap //TransportItem
	init            sync.Once
	err             error
}

var (
	defaultServiceMgr = NewServiceMgr()
)

func InitManager() {
	err := Default().LoadOnce()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ServiceMgr) Load(serviceName string) {

}

func Default() *ServiceMgr {
	return defaultServiceMgr
}

func NewServiceMgr() *ServiceMgr {
	return &ServiceMgr{
		GRPCServiceMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
			return []byte(key.(string))
		}),
		TCPServiceMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
			return []byte(key.(string))
		}),
		HTTPServiceMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
			return []byte(key.(string))
		}),
		redisServiceMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
			return []byte(key.(string))
		}),
		loadbalanceMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
			return []byte(key.(string))
		}),
		transportMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
			return []byte(key.(string))
		}),
		init: sync.Once{},
		err:  nil,
	}
}

func (m *ServiceMgr) GetRedisService(name string) (lib.RedisService, bool) {
	s, ok := m.redisServiceMap.Get(name)
	if !ok {
		if strings.HasPrefix(name, ServicePrefix) {
			newCount := NewRedisFlowCountService(name, 3*time.Second)
			m.SetRedisService(name, newCount)
			newCount.Start()
			return newCount, true
		}
		return nil, false
	}
	return s.(lib.RedisService), ok
}

func (m *ServiceMgr) SetRedisService(name string, val lib.RedisService) {
	m.redisServiceMap.Set(name, val)
}

func (m *ServiceMgr) GetTransport(name string) (*TransportItem, bool) {
	s, ok := m.transportMap.Get(name)
	if ok {
		return s.(*TransportItem), ok
	}
	return nil, ok
}

func (m *ServiceMgr) HTTPAccessMode(c *gin.Context) (HTTPService, error) {
	// fetch ip address
	host := c.Request.Host
	host = host[0:strings.Index(host, ":")]
	//
	path := c.Request.URL.Path
	// Fetch service related according condition.
	data, ok := m.HTTPServiceMap.GetByCondition(func(key, val interface{}) bool {
		service := val.(HTTPService)

		if service.RuleType == dao.HttpRule_Domain {
			if service.Rule == host {
				return true
			}
		}
		if service.RuleType == dao.HttpRule_PrefixURL {
			if service.Rule == "" {
				return false
			}
			// Until found service
			if strings.HasPrefix(path, service.Rule) {
				return true
			}
		}
		return false
	})
	if ok {
		return data.(HTTPService), nil
	}
	return HTTPService{}, Error_NoMatchedService
}

func (m *ServiceMgr) LoadOnce() (err error) {
	m.init.Do(func() {
		// fetch data from database.
		db := lib.GetDefaultDB()
		serviceInfo := &dao.ServiceInfo{}
		serviceInfoList, _, err := serviceInfo.PageListIdDesc(nil, db, &dao.PageSize{})
		if err != nil {
			return
		}
		// start total statistic redis service.
		totalStatistics := NewRedisFlowCountService(TotalPrefix, 0)
		totalStatistics.Start()
		m.redisServiceMap.Set(TotalPrefix, totalStatistics)

		for _, serviceInfo := range serviceInfoList {

			// fetch single service data.
			info := dao.ServiceInfo{ID: serviceInfo.ID, ServiceName: serviceInfo.ServiceName}
			serviceDetail, err := info.FindOneServiceDetail(nil, db)
			if err != nil {
				m.err = err
				return
			}

			// redis statistics service start
			statistics := NewRedisFlowCountService(ServicePrefix+serviceInfo.ServiceName, 0)
			statistics.Start()
			m.redisServiceMap.Set(ServicePrefix+serviceDetail.ServiceName, statistics)

			// add load balance instance for service.
			// add it to map.
			var lb loadbalance.LoadBalancer
			var hosts []dao.IpListItem
			switch serviceDetail.RoundType {
			case dao.Round_Random:
				lb = loadbalance.NewInst(loadbalance.RandomBalance{})
			case dao.Round_IpHash:
				lb = loadbalance.NewInst(loadbalance.ConsistentHashLoadBalancer{})
			case dao.Round_RoudRobin:
				lb = loadbalance.NewInst(loadbalance.RoundRobinLoadBalancer{})
			case dao.Round_WeightRound:
				lb = loadbalance.NewInst(loadbalance.WeightRobinLoadBalance{})
			}
			hosts = serviceDetail.GetHostsUrl()
			weights := serviceDetail.WeightList.GetSlice()
			for i, host := range hosts {
				newItem := host.URL
				if serviceDetail.RoundType == dao.Round_WeightRound {
					lb.AddHost(&loadbalance.HostUrl{
						URL:    &newItem,
						Weight: int(weights[i]),
					})
				} else {
					lb.AddHost(&loadbalance.HostUrl{
						URL: &newItem,
					})
				}
			}
			m.loadbalanceMap.Set(serviceDetail.ServiceName, lb)

			// add service to management map
			switch item := serviceDetail.Rule.(type) {
			case *dao.ServiceHTTPRuleExceptModel:
				{
					service := HTTPService{
						ServiceInfoExceptModel: *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalance: ServiceLoadBalance{
							CheckMethod:            serviceDetail.CheckMethod,
							CheckTimeout:           serviceDetail.CheckTimeout,
							CheckInterval:          serviceDetail.CheckInterval,
							RoundType:              serviceDetail.RoundType,
							IPList:                 serviceDetail.IPList.GetSlice(),
							WeightList:             serviceDetail.WeightList.GetSlice(),
							ForbidList:             serviceDetail.ForbidList.GetSlice(),
							UpstreamConnectTimeout: serviceDetail.UpstreamConnectTimeout,
							UpstreamHeaderTimeout:  serviceDetail.UpstreamHeaderTimeout,
							UpstreamIDleTimeout:    serviceDetail.UpstreamIDleTimeout,
							UpstreamMaxIDle:        serviceDetail.UpstreamMaxIDle,
						},
						ServiceAccessControl: ServiceAccessControl{
							OpenAuth:          serviceDetail.OpenAuth == dao.OpenAuth_Open,
							BlackList:         serviceDetail.BlackList.GetSlice(),
							WhiteList:         serviceDetail.WhiteList.GetSlice(),
							WhiteHostName:     serviceDetail.WhiteHostName,
							ClientipFlowLimit: serviceDetail.ClientipFlowLimit,
							ServiceFlowLimit:  serviceDetail.ServiceFlowLimit,
						},
						HTTPServiceRule: HTTPServiceRule{
							RuleType:        item.RuleType,
							Rule:            item.Rule,
							NeedHTTPs:       item.NeedHTTPs == dao.NeedHttps_Open,
							NeedStripURI:    item.NeedStripURI == dao.NeedStripUri_Open,
							NeedWebsocket:   item.NeedWebsocket == dao.NeedWebsocket_Open,
							URLRewrite:      item.URLRewrite,
							HeaderTransform: item.HeaderTransform.GetSlice(),
						},
					}
					//m.HTTPServiceList = append(m.HTTPServiceList, service)
					m.HTTPServiceMap.Set(service.ServiceName, service)
				}
			case *dao.ServiceTCPRuleExceptModel:
				{
					service := TCPService{
						ServiceInfoExceptModel:          *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalanceExceptModel:   *serviceDetail.ServiceLoadBalanceExceptModel,
						ServiceAccessControlExceptModel: *serviceDetail.ServiceAccessControlExceptModel,
						ServiceTCPRuleExceptModel:       *item,
					}
					m.TCPServiceMap.Set(service.ServiceName, service)
					//m.TCPServiceList = append(m.TCPServiceList, service)
				}
			case *dao.ServiceGRPCRuleExceptModel:
				{
					service := GRPCService{
						ServiceInfoExceptModel:          *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalanceExceptModel:   *serviceDetail.ServiceLoadBalanceExceptModel,
						ServiceAccessControlExceptModel: *serviceDetail.ServiceAccessControlExceptModel,
						ServiceGRPCRuleExceptModel:      *item,
					}
					m.GRPCServiceMap.Set(service.ServiceName, service)
					//m.GRPCServiceList = append(m.GRPCServiceList, service)
				}
			}
			// add transport
			trans := TransportItem{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   time.Duration(serviceDetail.UpstreamConnectTimeout) * time.Second,
						KeepAlive: 30 * time.Second,
						DualStack: true,
					}).DialContext,
					ForceAttemptHTTP2:     true,
					MaxIdleConns:          serviceDetail.UpstreamMaxIDle,
					IdleConnTimeout:       time.Duration(serviceDetail.UpstreamIDleTimeout) * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ResponseHeaderTimeout: time.Duration(serviceDetail.UpstreamHeaderTimeout) * time.Second,
				},
				name: serviceInfo.ServiceName,
			}
			m.transportMap.Set(serviceInfo.ServiceName, &trans)
		}
	})
	return m.err
}

// GetLoadBalancer fetch load balance instance according service name from internal map.
func (m *ServiceMgr) GetLoadBalancer(serviceName string) (loadbalance.LoadBalancer, bool) {
	lb, ok := m.loadbalanceMap.Get(serviceName)
	if ok {
		return lb.(loadbalance.LoadBalancer), ok
	}
	data, ok := m.HTTPServiceMap.Get(serviceName)
	if ok {
		service := data.(HTTPService)
		var newlb loadbalance.LoadBalancer
		switch service.RoundType {
		case dao.Round_Random:
			newlb = loadbalance.NewInst(loadbalance.RandomBalance{})
		case dao.Round_IpHash:
			newlb = loadbalance.NewInst(loadbalance.ConsistentHashLoadBalancer{})
		case dao.Round_RoudRobin:
			newlb = loadbalance.NewInst(loadbalance.RoundRobinLoadBalancer{})
		case dao.Round_WeightRound:
			newlb = loadbalance.NewInst(loadbalance.WeightRobinLoadBalance{})
		}
		if newlb != nil {
			for _, item := range service.IPList {
				newlb.AddHost(&loadbalance.HostUrl{URL: &item.URL})
			}
			m.loadbalanceMap.Set(serviceName, newlb)
			return newlb, ok
		}
	}

	return nil, ok
}
