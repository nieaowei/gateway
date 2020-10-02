package loadbalance

import (
	"hash/crc32"
	"net/url"
)

type ConsistentHashLoadBalancer struct {
	hostList []*url.URL
}

// Format:  192.168.1.1:9999
func (r *ConsistentHashLoadBalancer) Add(host string, hosts ...string) error {
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

func (c *ConsistentHashLoadBalancer) Get(key string) (*url.URL, error) {
	length := len(c.hostList)
	if length == 0 {
		return nil, Error_NoAvailableHost
	}
	hash := crc32.ChecksumIEEE([]byte(key))
	hashAddr := c.hostList[hash%uint32(length)]
	return hashAddr, nil
}

func (c *ConsistentHashLoadBalancer) Update() {
	panic("implement me")
}
