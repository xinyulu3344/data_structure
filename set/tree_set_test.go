package set

import (
	"fmt"
	"testing"
)

type Int int

func (i Int) CompareTo(e E) int {
    return int(i - e.(Int))
}

func TestXxx(t *testing.T) {
	// data := []int{11, 22, 33, 44, 11, 22, 33, 44}
	data := []Int{55, 66, 77, 55, 66, 77}

	s := NewTreeSet()
	for _, v := range data {
		s.Add(v)
	}
	s.Traversal(func(e E) bool {
		t.Log(e)
		return false
	})

	fmt.Println("==")
	s.Remove(Int(66))
	s.Traversal(func(e E) bool {
		t.Log(e)
		return false
	})

	t.Log(s.Contains(Int(66)))

	fmt.Println("==")
	s.Clear()
	s.Traversal(func(e E) bool {
		t.Log(e)
		return false
	})

	t.Log(s.IsEmpty())
}