package loadbalance

import (
	"hash/crc32"
)

type ConsistentHashLoadBalancer struct {
	hostList []*HostUrl
}

// Format:  192.168.1.1:9999
func (r *ConsistentHashLoadBalancer) Add(host *HostUrl, hosts ...*HostUrl) error {

	r.hostList = append(r.hostList, host)
	for _, h := range hosts {
		r.hostList = append(r.hostList, h)
	}
	return nil
}

func (c *ConsistentHashLoadBalancer) Get(key string) (*HostUrl, error) {
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
