package lib

import "sync"

type SafeMap interface {
	Set(key, val interface{})
	Get(key interface{}) (interface{}, bool)
	Mutex() sync.Locker
}

type RWMap struct {
	data map[interface{}]interface{}
	mu   *sync.RWMutex
}

func NewRWMap() *RWMap {
	return &RWMap{
		data: map[interface{}]interface{}{},
		mu:   &sync.RWMutex{},
	}
}

func (m *RWMap) Set(key, val interface{}) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *RWMap) Mutex() sync.Locker {
	return m.mu
}

func (m *RWMap) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	data, ok := m.data[key]
	m.mu.RUnlock()
	return data, ok
}
