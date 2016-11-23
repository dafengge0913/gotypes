package set_test

import (
	"github.com/dafengge0913/gotypes/set"
	"testing"
)

func TestSet(t *testing.T) {
	s := set.NewSet("1", "2", "3")
	s.Add("4")
	t.Log(s.Len())
	t.Log(s.Contain("2", "1"))
	s.Del("2")
	t.Log(s.Contain("2"))
	t.Log(s.Len())
	t.Log(s)

	for _, x := range s.List() {
		t.Log(x)
	}

	s.Clear()
	t.Log(s.Len())



}
