package lib

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkRWMap_Get_Read(b *testing.B) {
	m := NewConcurrentHashMap(1024, func(i interface{}) []byte {
		return []byte(strconv.Itoa(i.(int)))
	})
	for i := 0; i < 10000; i++ {
		go m.Set(i, i)

	}
	for i := 0; i < b.N; i++ {
		c := rand.Intn(10000)
		go m.Get(c)
		//if i%100 == 0 {
		//	m.GetByCondition(func(a interface{}) bool {
		//		data := a.(int)
		//		if data == b {
		//			fmt.Println(b,data)
		//			return true
		//		}
		//		return false
		//	})
		//}
	}
}
func BenchmarkRWMap_Get1_Read(b *testing.B) {
	//m := NewConcurrentHashMap(1024, func(i interface{}) []byte {
	//	return []byte(strconv.Itoa(i.(int)))
	//})
	m := sync.Map{}
	for i := 0; i < 10000; i++ {
		go m.Store(i, i)
	}
	for i := 0; i < b.N; i++ {
		c := rand.Intn(10000)
		go m.Load(c)
		//if i%100 == 0 {
		//	m.Range(func(key, value interface{}) bool {
		//		data := value.(int)
		//		if data == b {
		//			fmt.Println(b,data)
		//			return true
		//		}
		//		return false
		//	})
		//
		//	//m.GetByCondition(func(a interface{}) bool {
		//	//	data := a.(int)
		//	//	if data == b{
		//	//		return true
		//	//	}
		//	//	return false
		//	//})
		//}
	}
}

func BenchmarkRWMap_Get_Write(b *testing.B) {
	m := NewConcurrentHashMap(1024, func(i interface{}) []byte {
		return []byte(strconv.Itoa(i.(int)))
	})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go m.Set(i, i)

	}
	//for i := 0; i < b.N; i++ {
	//	c := rand.Intn(10000)
	//	go m.GetHost(c)
	//	//if i%100 == 0 {
	//	//	m.GetByCondition(func(a interface{}) bool {
	//	//		data := a.(int)
	//	//		if data == b {
	//	//			fmt.Println(b,data)
	//	//			return true
	//	//		}
	//	//		return false
	//	//	})
	//	//}
	//}
}
func BenchmarkRWMap_Get1_Write(b *testing.B) {
	//m := NewConcurrentHashMap(1024, func(i interface{}) []byte {
	//	return []byte(strconv.Itoa(i.(int)))
	//})
	m := sync.Map{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go m.Store(i, i)
	}
	//b.ResetTimer()
	//for i := 0; i < b.N; i++ {
	//	c := rand.Intn(10000)
	//	go m.Load(c)
	//	//if i%100 == 0 {
	//	//	m.Range(func(key, value interface{}) bool {
	//	//		data := value.(int)
	//	//		if data == b {
	//	//			fmt.Println(b,data)
	//	//			return true
	//	//		}
	//	//		return false
	//	//	})
	//	//
	//	//	//m.GetByCondition(func(a interface{}) bool {
	//	//	//	data := a.(int)
	//	//	//	if data == b{
	//	//	//		return true
	//	//	//	}
	//	//	//	return false
	//	//	//})
	//	//}
	//}
}

func BenchmarkRWMap_Get_Serach(b *testing.B) {
	m := NewConcurrentHashMap(1024, func(i interface{}) []byte {
		return []byte(strconv.Itoa(i.(int)))
	})

	for i := 0; i < 10000; i++ {
		go m.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := rand.Intn(10000)
		if i%100 == 0 {
			m.GetByCondition(func(k, v interface{}) bool {
				data := v.(int)
				if data == c {
					return true
				}
				return false
			})
		}
	}
}
func BenchmarkRWMap_Get1_Serach(b *testing.B) {
	//m := NewConcurrentHashMap(1024, func(i interface{}) []byte {
	//	return []byte(strconv.Itoa(i.(int)))
	//})
	m := sync.Map{}
	for i := 0; i < 10000; i++ {
		go m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := rand.Intn(10000)
		if i%100 == 0 {
			m.Range(func(key, value interface{}) bool {
				data := value.(int)
				if data == c {
					return false
				}
				return true
			})

		}
	}
}
