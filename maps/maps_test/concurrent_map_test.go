package maps_test

import (
	"github.com/dafengge0913/gotypes/maps"
	"sync"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	cm := maps.NewConcurrentMap()
	wg := &sync.WaitGroup{}
	wg.Add(20000)
	for g := 0; g < 100; g++ {
		for i := 0; i < 100; i++ {
			go func(n int) {
				cm.Put(n, n)
				wg.Done()
			}(i)
			go func(n int) {
				cm.Get(i)
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	t.Log(cm.Len())
	t.Logf("%v", cm.Data())
	t.Logf("sorted keys: %v", cm.SortedKeys(intLess))
	cm.Clear()
	t.Log(cm.Len())
}

func intLess(a, b interface{}) bool {
	if a.(int) < b.(int) {
		return true
	}
	return false
}

func BenchmarkCcMapPut(b *testing.B) {
	cm := maps.NewConcurrentMap()
	for i := 0; i < b.N; i++ {
		cm.Put(i, i)
	}
}

func BenchmarkCcMapGet(b *testing.B) {
	cm := maps.NewConcurrentMap()
	for i := 0; i < b.N; i++ {
		cm.Get(i)
	}
}
