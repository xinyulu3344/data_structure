package myheap

import (
    "fmt"
    "testing"
)

type IntComparator []int

func (m *IntComparator) CompareTo(e1 interface{}, e2 interface{}) int {
    i1, _ := e1.(int)
    i2, _ := e2.(int)
    return i1 - i2
}

func TestBinaryHeap(t *testing.T) {
    intComparator := &IntComparator{}
    bheap := NewBinaryHeapWithComparator(intComparator)
    bheap.Add(68)
    bheap.Add(72)
    bheap.Add(43)
    bheap.Add(50)
    bheap.Add(38)
    bheap.Add(10)
    bheap.Add(90)
    bheap.Add(65)
    fmt.Println(bheap.elements)
    //bheap.Remove()
    //fmt.Println(bheap.elements)
    bheap.Replace(70)
    fmt.Println(bheap.elements)
}