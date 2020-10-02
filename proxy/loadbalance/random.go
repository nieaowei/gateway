package loadbalance

import (
	"math/rand"
	"net/url"
)

type RandomBalance struct {
	index    uint
	hostList []*url.URL
	conf     BalanceConf
}

// Format:  192.168.1.1:9999
func (r *RandomBalance) Add(host string, hosts ...string) error {
	addr, err := url.Parse(host)
	if err != nil {
		return err
	}
	r.hostList = append(r.hostList, addr)
	for _, h := range hosts {
		a, err := url.Parse(h)
		if err != nil {
			continue
		}
		r.hostList = append(r.hostList, a)
	}
	return nil
}

func (r *RandomBalance) Get(key string) (*url.URL, error) {
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
