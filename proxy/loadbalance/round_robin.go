package loadbalance

import "net/url"

type RoundRobinLoadBalancer struct {
	currentIndex int
	hostList     []*url.URL
}

// Format:  192.168.1.1:9999
func (r *RoundRobinLoadBalancer) Add(host string, hosts ...string) error {
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

func (r *RoundRobinLoadBalancer) Get(key string) (*url.URL, error) {
	length := len(r.hostList)

	if length == 0 {
		return nil, Error_NoAvailableHost
	}

	if r.currentIndex >= length {
		r.currentIndex = 0
	}

	current := r.currentIndex

	r.currentIndex = (r.currentIndex + 1) % length

	return r.hostList[current], nil
}

func (r *RoundRobinLoadBalancer) Update() {
	panic("implement me")
}
