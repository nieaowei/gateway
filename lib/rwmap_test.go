package lib

import (
	"sync"
	"testing"
)

func BenchmarkThreadMap_Get(b *testing.B) {
	data := map[int]int{}
	mu := sync.RWMutex{}
	m := ThreadMap{
		get: func(name interface{}) (interface{}, bool) {
			mu.RLock()
			da, ok := data[name.(int)]
			mu.RUnlock()
			return da, ok
		},
		set: func(name, value interface{}) {
			mu.Lock()
			data[name.(int)] = value.(int)
			mu.Unlock()
		},
	}
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go m.Set(i, i)
		//m.Get(i)
		go func(threadMap ThreadMap) {
			for i := 0; i < 5; i++ {
				threadMap.Get(i)
			}
		}(m)
	}
}

func BenchmarkRWMap_Get(b *testing.B) {
	m := NewRWMap()
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go m.Set(i, i)
		//m.Get(i)
		go func(threadMap *RWMap) {
			for i := 0; i < 5; i++ {
				threadMap.Get(i)
			}
		}(m)
	}
}
