package loadbalance

import (
	"math/rand"
)

type RandomBalance struct {
	index    uint
	hostList []*HostUrl
	conf     BalanceConf
}

// Format:  192.168.1.1:9999
func (r *RandomBalance) Add(host *HostUrl, hosts ...*HostUrl) error {
	r.hostList = append(r.hostList, host)
	for _, h := range hosts {
		r.hostList = append(r.hostList, h)
	}
	return nil
}

func (r *RandomBalance) Get(key string) (*HostUrl, error) {
	length := len(r.hostList)
	if length == 0 {
		return nil, Error_NoAvailableHost
	}
	index := rand.Intn(length)
	return r.hostList[index], nil
}

func (r *RandomBalance) Update() {
	panic("implement me")
}
