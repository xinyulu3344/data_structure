package linked_list

import (
    "testing"
)

type Int int

func (i Int) Equal(j E) bool {
    return i == j.(Int)
}

type Str string

func (s Str) Equal(j E) bool {
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

func (p Persion) Equal(j E) bool {
    if persion, ok := j.(Persion); ok {
        return p.Age == persion.Age
    }
    return false
}

func TestLinkedList(t *testing.T) {
    list := NewLinkedList()
    list.Append(Int(1))
    list.Append(Int(2))
    list.Add(0, Int(3))
    list.Append(Int(4))
    list.Add(list.Size(), Int(5))
    list.Remove(1)
    
    if list.Size() != 4 {
        t.Errorf("%v list.Size() = %v != 5\n", list, list.Size())
    }
    
    if !list.Get(0).Equal(Int(3)) || !list.Get(1).Equal(Int(2)) || !list.Get(2).Equal(Int(4)) || !list.Get(3).Equal(Int(5)) {
        t.Errorf("list.Get() err\n")
    }
    
    list.Set(0, Int(6))
    list.Set(list.Size()-1, Int(7))
    list.Set(2, Int(8))
    
    if !list.Get(0).Equal(Int(6)) || !list.Get(1).Equal(Int(2)) || !list.Get(2).Equal(Int(8)) || !list.Get(3).Equal(Int(7)) {
        t.Errorf("list.Get() err\n")
    }
    
    list.Clear()
    
    if list.Size() != 0 {
        t.Errorf("list.Clear(), Size err")
    }
}
