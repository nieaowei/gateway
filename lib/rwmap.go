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

type GetHandler func(name interface{}) (interface{}, bool)
type SetHandler func(name, value interface{})

type ThreadMap struct {
	get func(name interface{}) (interface{}, bool)
	set func(name, value interface{})
}

func NewThreadMap(set SetHandler, get GetHandler) *ThreadMap {
	return &ThreadMap{
		get: get,
		set: set,
	}
}

func (t *ThreadMap) Set(key, val interface{}) {
	t.set(key, val)
}

func (t *ThreadMap) Get(key interface{}) (interface{}, bool) {
	return t.get(key)
}
