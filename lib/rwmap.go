package lib

import (
	"hash/crc32"
	"sync"
)

type SafeMap interface {
	Set(key, val interface{})
	Get(key interface{}) (interface{}, bool)
	Mutexs() []sync.RWMutex
	GetByCondition(f func(key, val interface{}) bool) (interface{}, bool)
}

type ConcurrentHashMap struct {
	data []map[interface{}]interface{}
	mus  []sync.RWMutex
	// Partition size
	size uint32
	// Hash format function.
	// Example: from int to []byte].
	// func(key interface{}) []byte {
	//		return []byte(strings.Itoa(key))
	// }
	hashFormat func(key interface{}) []byte
}

func NewConcurrentHashMap(bufSize uint32, hashFormat func(key interface{}) []byte) *ConcurrentHashMap {
	// Allocate storage space for partition.
	d := make([]map[interface{}]interface{}, bufSize)
	// Allocate storage space for lock.
	m := make([]sync.RWMutex, bufSize)
	// Generate instance for partition and lock.
	for i := uint32(0); i < bufSize; i++ {
		d[i] = make(map[interface{}]interface{})
		m[i] = sync.RWMutex{}
	}

	return &ConcurrentHashMap{
		size:       bufSize,
		data:       d,
		mus:        m,
		hashFormat: hashFormat,
	}
}

func (m *ConcurrentHashMap) Set(key, val interface{}) {
	hashVal := crc32.ChecksumIEEE(m.hashFormat(key)) % m.size
	m.mus[hashVal].Lock()
	m.data[hashVal][key] = val
	m.mus[hashVal].Unlock()
}

func (m *ConcurrentHashMap) Mutexs() []sync.RWMutex {
	return m.mus
}

func (m *ConcurrentHashMap) Get(key interface{}) (interface{}, bool) {
	hashVal := crc32.ChecksumIEEE(m.hashFormat(key)) % m.size
	m.mus[hashVal].RLock()
	data, ok := m.data[hashVal][key]
	m.mus[hashVal].RUnlock()
	return data, ok
}

func (m *ConcurrentHashMap) GetByCondition(f func(key, value interface{}) bool) (interface{}, bool) {
	for i := uint32(0); i < m.size; i++ {
		for k, v := range m.data[i] {
			if f(k, v) == true {
				return v, true
			}
		}
	}
	return nil, false
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
