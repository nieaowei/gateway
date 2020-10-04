package loadbalance

import (
	"net/url"
)

type HostUrl struct {
	*url.URL
	Weight int
}

type LoadBalancer interface {
	Add(*HostUrl, ...*HostUrl) error
	Get(key string) (*HostUrl, error)
	Update()
}

type BalanceConf interface {
}

type LoadBalancerMgr interface {
	GetLoadBalancer(serviceName string) (LoadBalancer, bool)
}

func NewInst(t interface{}) LoadBalancer {
	switch t.(type) {
	case RoundRobinLoadBalancer:
		return &RoundRobinLoadBalancer{
			currentIndex: 0,
			hostList:     []*HostUrl{},
		}
	case WeightRobinLoadBalance:
		return &WeightRobinLoadBalance{
			hostList:      []*WeightNode{},
			currentWeight: 0,
		}
	case RandomBalance:
		return &RandomBalance{
			index:    0,
			hostList: []*HostUrl{},
			conf:     nil,
		}
	case ConsistentHashLoadBalancer:
		return &ConsistentHashLoadBalancer{
			hostList: []*HostUrl{},
		}
	}
	return nil
}

//var DefaultLoadBalanceMgr *Mgr
//
//func init() {
//	DefaultLoadBalanceMgr = NewMgr()
//}
//
//type Mgr struct {
//	LoadBalancerMap lib.SafeMap
//}
//
//func NewMgr() *Mgr {
//	return &Mgr{
//		LoadBalancerMap: lib.NewConcurrentHashMap(1024, func(key interface{}) []byte {
//			return []byte(key.(string))
//		}),
//	}
//}
//
//func (m *Mgr) GetLoadBalancer(serviceName string) (LoadBalancer, bool) {
//	lb, ok := m.LoadBalancerMap.Get(serviceName)
//	if ok {
//		return lb.(LoadBalancer), ok
//	}
//	return lb.(LoadBalancer), ok
//}
