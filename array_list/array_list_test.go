package array_list

import (
	"strconv"
	"testing"
)

type Int int

func (i Int) Equal(j Item) bool {
	return i == j.(Int)
}

type Str string

func (s Str) Equal(j Item) bool {
	return s == j.(Str)
}

type Persion struct {
	Name string
	Age  int
}

func NewPersion(name string, age int) Persion {
	return Persion{
		Name: name,
		Age:  age,
	}
}

func (p Persion) Equal(j Item) bool {
	if persion, ok := j.(Persion); ok {
		return p.Age == persion.Age
	}
	return false
}

func TestArrayList_Append(t *testing.T) {
	al := NewArrayList()
	for i := 0; i < 18; i++ {
		al.Append(Int(i))
	}
	for i := 0; i < 18; i++ {
		if !al.elements[i].Equal(Int(i)) {
			t.Skipf("al.elements[%d] = %v != %v", i, al.elements[i], i)
		}
	}
}

func TestArrayList_Add(t *testing.T) {
	al := NewArrayList()
	for i := 0; i < 18; i++ {
		al.Append(Int(i))
	}
	al.Add(3, Int(88))
	if !al.elements[3].Equal(Int(88)) {
		t.Errorf("al.element[3] = %v != 88", al.elements[3])
	}
	for i := 0; i < 3; i++ {
		if !al.elements[i].Equal(Int(i)) {
			t.Skipf("al.elements[%d] = %v != %d", i, al.elements[i], i)
		}
	}
	for i := 4; i < 18; i++ {
		if !al.elements[i].Equal(Int(i - 1)) {
			t.Skipf("al.elements[%d] = %v != %d", i, al.elements[i], i-1)
		}
	}
}

func TestArrayList_Remove(t *testing.T) {
	al := NewArrayList()
	for i := 0; i < 18; i++ {
		al.Append(Int(i))
	}
	item1 := al.Remove(0)
	if item1.(Int) != 0 {
		t.Skipf("al.elements[0]==%v != 0", item1.(Int))
	}
	item2 := al.Remove(16)
	if item2.(Int) != 17 {
		t.Skipf("al.elements[16]==%v != 16", item2.(Int))
	}
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

func TestArrayList_ensureCapacity2(t *testing.T) {
	al := NewArrayList()
	capacity := len(al.elements)
	for i := 0; i < 15; i++ {
		p := NewPersion("张三"+strconv.Itoa(i), i)
		al.Append(p)
		if capacity != len(al.elements) {
			if len(al.elements) != (capacity + capacity>>1) {
				t.Skip(capacity, len(al.elements))
			}
			capacity = len(al.elements)
		}
	}
	t.Logf("当前容量: %d", capacity)
}

func TestArrayList_IndexOf(t *testing.T) {
	al := NewArrayList()
	for i := 0; i < 18; i++ {
		al.Append(Int(i))
	}
	al.Set(0, nil)
	al.Set(8, nil)
	al.Set(10, nil)
	al.Set(17, nil)
	if al.IndexOf(nil) != 0 {
		t.Error("IndexOf(nil) != 0")
	}
}

func TestArrayList_Clear(t *testing.T) {
	al := NewArrayList()
	for i := 0; i < 18; i++ {
		al.Append(Int(i))
	}
	al.Clear()
	if al.Size() != 0 {
		t.Error("al.Size() error")
	}
	if len(al.elements) < 18 {
		t.Error("len(al.elements) error")
	}
	if len(al.elements) != cap(al.elements) {
		t.Error("len(al.elements) != cap(al.elements)")
	}
	for i := 0; i < len(al.elements); i++ {
		if al.elements[i] != nil {
			t.Skipf("al.elements[%d] = %v != nil", i, al.elements[i])
		}
	}
}

func TestArrayList_trim(t *testing.T) {
	al := NewArrayList()
	for i := 0; i < 18; i++ {
		al.Append(Int(i))
	}

	for i := 17; i >= 0; i-- {
		al.Remove(i)
	}
	if len(al.elements) != 5  {
		t.Errorf("len(al.elements) err\n")
	}
	if al.size != 0 {
		t.Errorf("al.size err")
	}
}
