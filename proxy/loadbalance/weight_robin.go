package loadbalance

import (
	"net/url"
)

type WeightRobinLoadBalance struct {
	hostList      []*WeightNode
	currentWeight int
}

type WeightNode struct {
	Weight        int
	CurrentWeight int
	Addr          *url.URL
}

// Format:  192.168.1.1:9999 60
func (r *WeightRobinLoadBalance) AddHost(host *HostUrl, hosts ...*HostUrl) error {
	//params := strings.Split(host, " ")
	//addr, err := url.Parse(params[0])
	//if err != nil {
	//	return err
	//}
	//weight, err := strconv.Atoi(params[1])

	//if err != nil {
	//	return Error_AddNode
	//}
	node := &WeightNode{
		Weight:        host.Weight,
		CurrentWeight: 0,
		Addr:          host.URL,
	}
	r.hostList = append(r.hostList, node)
	if len(hosts) != 0 {
		for _, h := range hosts {

			node := &WeightNode{
				Weight:        h.Weight,
				CurrentWeight: 0,
				Addr:          h.URL,
			}
			r.hostList = append(r.hostList, node)
		}
	}

	return nil
}

func (w *WeightRobinLoadBalance) GetHost(key string) (*HostUrl, error) {
	total := 0
	var best *WeightNode
	for _, node := range w.hostList {
		total += node.Weight
		node.CurrentWeight += node.Weight
		if best == nil || node.CurrentWeight > best.CurrentWeight {
			best = node
		}
	}
	if best != nil {
		best.CurrentWeight -= total
		return &HostUrl{
			URL:    best.Addr,
			Weight: best.Weight,
		}, nil
	}
	return nil, Error_NoAvailableHost
}

func (w *WeightRobinLoadBalance) Update() {
	panic("implement me")
}
