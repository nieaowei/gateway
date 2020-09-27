package manager

import (
	"errors"
	"gateway/dao"
	"gateway/lib"
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

type ServiceMgr struct {
	TCPServiceList  []TCPService
	HTTPServiceList []HTTPService
	GRPCServiceList []GRPCService
	redisServiceMap *lib.RWMap
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
		TCPServiceList:  []TCPService{},
		HTTPServiceList: []HTTPService{},
		GRPCServiceList: []GRPCService{},
		redisServiceMap: lib.NewRWMap(),
		init:            sync.Once{},
		err:             nil,
	}
}

func (m *ServiceMgr) GetRedisService(name string) (lib.RedisService, bool) {
	s, ok := m.redisServiceMap.Get(name)
	return s.(lib.RedisService), ok
}

func (m *ServiceMgr) SetRedisService(name string, val lib.RedisService) {
	m.redisServiceMap.Set(name, val)
}

func (m *ServiceMgr) HTTPAccessMode(c *gin.Context) (HTTPService, error) {
	host := c.Request.Host
	host = host[0:strings.Index(host, ":")]
	path := c.Request.URL.Path
	for _, service := range m.HTTPServiceList {
		if service.LoadType != dao.LoadType_HTTP {
			continue
		}
		if service.RuleType == dao.HttpRuleType_Domain {
			if service.Rule == host {
				return service, nil
			}
		}
		if service.RuleType == dao.HttpRuleType_PrefixURL {
			if service.Rule == "" {
				continue
			}
			if strings.HasPrefix(path, service.Rule) {
				return service, nil
			}
		}
	}
	return HTTPService{}, errors.New("not matched service")
}

func (m *ServiceMgr) LoadOnce() (err error) {
	m.init.Do(func() {
		db := lib.GetDefaultDB()
		serviceInfo := &dao.ServiceInfo{}
		serviceInfoList, _, err := serviceInfo.PageListIdDesc(nil, db, &dao.PageSize{})
		if err != nil {
			return
		}
		for _, serviceInfo := range serviceInfoList {
			serviceDetail, err := serviceInfo.FindOneServiceDetail(nil, db)
			if err != nil {
				m.err = err
				return
			}
			statistics := NewRedisFlowCountService(serviceInfo.ServiceName, time.Second)

			statistics.Start()

			m.redisServiceMap.Set(serviceDetail.ServiceName, statistics)

			switch item := serviceDetail.Rule.(type) {
			case *dao.ServiceHTTPRuleExceptModel:
				{
					service := HTTPService{
						ServiceInfoExceptModel:          *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalanceExceptModel:   *serviceDetail.ServiceLoadBalanceExceptModel,
						ServiceAccessControlExceptModel: *serviceDetail.ServiceAccessControlExceptModel,
						ServiceHTTPRuleExceptModel:      *item,
					}
					m.HTTPServiceList = append(m.HTTPServiceList, service)
				}
			case *dao.ServiceTCPRuleExceptModel:
				{
					service := TCPService{
						ServiceInfoExceptModel:          *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalanceExceptModel:   *serviceDetail.ServiceLoadBalanceExceptModel,
						ServiceAccessControlExceptModel: *serviceDetail.ServiceAccessControlExceptModel,
						ServiceTCPRuleExceptModel:       *item,
					}
					m.TCPServiceList = append(m.TCPServiceList, service)
				}
			case *dao.ServiceGRPCRuleExceptModel:
				{
					service := GRPCService{
						ServiceInfoExceptModel:          *serviceDetail.ServiceInfoExceptModel,
						ServiceLoadBalanceExceptModel:   *serviceDetail.ServiceLoadBalanceExceptModel,
						ServiceAccessControlExceptModel: *serviceDetail.ServiceAccessControlExceptModel,
						ServiceGRPCRuleExceptModel:      *item,
					}
					m.GRPCServiceList = append(m.GRPCServiceList, service)
				}
			}
		}
	})
	return m.err
}
