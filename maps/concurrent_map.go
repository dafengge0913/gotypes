package maps

import (
	"sync"
)

type ConcurrentMap struct {
	*sync.RWMutex
	data map[interface{}]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		RWMutex: &sync.RWMutex{},
		data:    make(map[interface{}]interface{}),
	}
}

func (m *ConcurrentMap) Put(key interface{}, elem interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = elem
}

func (m *ConcurrentMap) Get(key interface{}) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	v, fd := m.data[key]
	return v, fd
}

func (m *ConcurrentMap) Remove(key interface{}) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := m.data[key]
	if ok {
		delete(m.data, key)
		return true
	}
	return false
}

func (m *ConcurrentMap) Data() map[interface{}]interface{} {
	return m.data
}

func (m *ConcurrentMap) Len() int {
	return len(m.data)
}

func (m *ConcurrentMap) Clear() {
	for k, _ := range m.data {
		delete(m.data, k)
	}
}

func (m *ConcurrentMap) SortedKeys(less func(a, b interface{}) bool) []interface{} {
	keys := make([]interface{}, 0, m.Len())
	for k := range m.Data() {
		keys = append(keys, k)
	}
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if less(keys[j], keys[i]) {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	return keys
}
