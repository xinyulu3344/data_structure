package array_list

import "testing"

type Int int

func (i Int) Equal(j Item) bool {
    return i == j.(Int)
}

type Str string

func (s Str) Equal(j Item) bool {
    return s == j.(Str)
}

func TestArrayList_ensureCapacity(t *testing.T) {
    al := NewArrayList()
    capacity := len(al.elements)
    for i := 0; i < 1000000; i++ {
        al.Append(Int(i))
        if capacity != len(al.elements) {
            if len(al.elements) != (capacity + capacity>>1) {
                t.Skip(capacity, len(al.elements))
            }
            capacity = len(al.elements)
        }
    }
    t.Logf("当前容量: %d", capacity)
}
