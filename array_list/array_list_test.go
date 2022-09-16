package array_list

import (
    "fmt"
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

func TestArrayList(t *testing.T) {
    al := NewArrayListWithCap(10)
    al.Append(Int(10))
    al.Append(Int(12))
    al.Append(Int(14))
    fmt.Println(al)
    
    al2 := NewArrayListWithCap(10)
    al2.Append(Str("a"))
    al2.Append(Str("b"))
    al2.Append(Str("c"))
    fmt.Println(al2)
}
