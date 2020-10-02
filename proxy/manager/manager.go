package manager

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/proxy/loadbalance"
	"github.com/gin-gonic/gin"
	"log"
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

type HTTPService struct {
	dao.ServiceInfoExceptModel
	dao.ServiceHTTPRuleExceptModel
	dao.ServiceLoadBalanceExceptModel
	dao.ServiceAccessControlExceptModel
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
	GRPCServiceMap  lib.SafeMap
	TCPServiceMap   lib.SafeMap
	HTTPServiceMap  lib.SafeMap
	loadbalanceMap  lib.SafeMap
	redisServiceMap lib.SafeMap
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
	}
	return s.(lib.RedisService), ok
}

func (m *ServiceMgr) SetRedisService(name string, val lib.RedisService) {
	m.redisServiceMap.Set(name, val)
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

		if service.RuleType == dao.HttpRuleType_Domain {
			if service.Rule == host {
				return true
			}
		}
		if service.RuleType == dao.HttpRuleType_PrefixURL {
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
			serviceDetail, err := serviceInfo.FindOneServiceDetail(nil, db)
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
			var hosts []string
			switch serviceDetail.RoundType {
			case dao.RoundType_Random:
				lb = loadbalance.NewInst(loadbalance.RandomBalance{})
			case dao.RoundType_IpHash:
				lb = loadbalance.NewInst(loadbalance.ConsistentHashLoadBalancer{})
			case dao.RoundType_RoudRobin:
				lb = loadbalance.NewInst(loadbalance.RoundRobinLoadBalancer{})
			case dao.RoundType_WeightRound:
				lb = loadbalance.NewInst(loadbalance.WeightRobinLoadBalance{})
			}
			hosts = serviceDetail.GetIpList()
			if lb != nil && hosts != nil {
				if len(hosts) >= 2 {
					lb.Add(hosts[0], hosts[1:]...)
				} else {
					lb.Add(hosts[0])
				}
			}
			m.loadbalanceMap.Set(serviceDetail.ServiceName, lb)
			// add service to management map
			switch item := serviceDetail.Rule.(type) {
			case *dao.ServiceHTTPRuleExceptModel:
				{
					service := HTTPService{
						ServiceInfoExceptModel:          *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalanceExceptModel:   *serviceDetail.ServiceLoadBalanceExceptModel,
						ServiceAccessControlExceptModel: *serviceDetail.ServiceAccessControlExceptModel,
						ServiceHTTPRuleExceptModel:      *item,
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
	return nil, ok
}
