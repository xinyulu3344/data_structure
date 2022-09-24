package doubly_linked_list

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

func TestDoublyLinkedList(t *testing.T) {
    list := NewDoublyLinkedList()
    list.Append(Int(11))
    list.Append(Int(22))
    list.Append(Int(33))
    list.Append(Int(44)) // [11, 22, 33, 44]
    
    list.Add(0, Int(55))           // [55, 11, 22, 33, 44]
    list.Add(2, Int(66))           // [55, 11, 66, 22, 33, 44]
    list.Add(list.Size(), Int(77)) // [55, 11, 66, 22, 33, 44, 77]
    
    list.Remove(0)               // [11, 66, 22, 33, 44, 77]
    list.Remove(2)               // [11, 66, 33, 44, 77]
    list.Remove(list.Size() - 1) // [11, 66, 33, 44]
    
    if list.IndexOf(Int(44)) != 3 {
        t.Fail()
    }
    if list.IndexOf(Int(22)) != ELEMENT_NOT_FOUND {
        t.Log(list.IndexOf(Int(22)))
        t.Errorf("list.IndexOf(Int(22)) != -1")
    }
    
    if !list.Contains(Int(33)) {
        t.Errorf("list.Contains(Int(33)) is false")
    }
    if list.Get(0) != Int(11) {
        t.Errorf("list.Get(0) != 11")
    }
    if list.Get(1) != Int(66) {
        t.Errorf("list.Get(1) != 66")
    }
    if list.Get(list.Size()-1) != Int(44) {
        t.Errorf("list.Get(list.Size() - 1 != 44")
    }
}
