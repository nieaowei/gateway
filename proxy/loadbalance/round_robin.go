package loadbalance

type RoundRobinLoadBalancer struct {
	currentIndex int
	hostList     []*HostUrl
}

// Format:  192.168.1.1:9999
func (r *RoundRobinLoadBalancer) Add(host *HostUrl, hosts ...*HostUrl) error {
	r.hostList = append(r.hostList, host)
	for _, h := range hosts {
		r.hostList = append(r.hostList, h)
	}
	return nil
}

func (r *RoundRobinLoadBalancer) Get(key string) (*HostUrl, error) {
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
