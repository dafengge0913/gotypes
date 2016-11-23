package set

import "fmt"

type ISet interface {
	Add(interface{}) bool
	Del(interface{}) bool
	Len() int
	Clear()
	Contain(...interface{}) bool
	List() []interface{}
}

//unsafe
type Set struct {
	data map[interface{}]bool
}

func NewSet(data ...interface{}) ISet {
	set := &Set{
		data: make(map[interface{}]bool, len(data)),
	}
	for _, x := range data {
		set.Add(x)
	}
	return set
}

func (s *Set) Add(x interface{}) bool {
	if _, fd := s.data[x]; fd {
		return false
	}
	s.data[x] = true
	return true
}

func (s *Set) Del(x interface{}) bool {
	if _, fd := s.data[x]; fd {
		delete(s.data, x)
		return true
	}
	return false
}

func (s *Set) Len() int {
	return len(s.data)
}

func (s *Set) Clear() {
	for x, _ := range s.data {
		delete(s.data, x)
	}
}

func (s *Set) Contain(data ...interface{}) bool {
	for _, x := range data {
		if _, fd := s.data[x]; !fd {
			return false
		}
	}
	return true
}

func (s *Set) String() string {
	if s.Len() <= 0 {
		return "{}"
	}
	str := "{"
	for x, _ := range s.data {
		str += fmt.Sprintf("%v,", x)
	}
	return str[:len(str)-1] + "}"
}

func (s *Set) List() []interface{} {
	list := make([]interface{}, s.Len())
	i := 0
	for x, _ := range s.data {
		list[i] = x
		i++
	}
	return list
}
