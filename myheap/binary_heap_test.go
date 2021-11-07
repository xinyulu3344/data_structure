package myheap

import (
    "fmt"
    "testing"
)

// 大顶堆比较器
type IntComparator int

func (i *IntComparator) CompareTo(e1 interface{}, e2 interface{}) int {
    i1, _ := e1.(int)
    i2, _ := e2.(int)
    return i1 - i2
}

func (i IntComparator) Compare(e1 interface{}) int {
    return int(i) - int(e1.(IntComparator))
}

// 小顶堆比较器
type IntComparator1 int

func (i *IntComparator1) CompareTo(e1 interface{}, e2 interface{}) int {
    i1, _ := e1.(int)
    i2, _ := e2.(int)
    return i2 - i1
}

func (i IntComparator1) Compare(e1 interface{}) int {
    return int(e1.(IntComparator)) - int(i)
}

// 测试不带比较器的大顶堆
func TestNewBinaryHeap(t *testing.T) {
    bheap := NewBinaryHeap()
    bheap.Add(IntComparator(68))
    bheap.Add(IntComparator(72))
    bheap.Add(IntComparator(43))
    bheap.Add(IntComparator(50))
    bheap.Add(IntComparator(38))
    bheap.Add(IntComparator(10))
    bheap.Add(IntComparator(90))
    bheap.Add(IntComparator(65))
    fmt.Println(bheap.elements)
}

// 测试带比较器的大顶堆
func TestNewBinaryHeapWithComparator(t *testing.T) {
    var intComparator *IntComparator
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
    bheap.Remove()
    fmt.Println(bheap.elements)
    bheap.Replace(70)
    fmt.Println(bheap.elements)
}

func TestNewBinaryHeapHeapify(t *testing.T) {
    var intComparator *IntComparator
    p1 := []interface{}{68, 72, 43, 50, 38, 10, 90 ,65}
    bheap := NewBinaryHeapHeapify(p1, intComparator)
    fmt.Println(bheap.elements)
}

func TestNewBinaryHeapWithComparator1(t *testing.T) {
    var intComparator1 *IntComparator1
    p1 := []interface{}{68, 72, 43, 50, 38, 10, 90 ,65}
    bheap := NewBinaryHeapHeapify(p1, intComparator1)
    fmt.Println(bheap.elements)
}

// 找出最大的前k个数
func TestTopK(t *testing.T) {
    k := 5
    var intComparator1 *IntComparator1
    data := []interface{}{19, 11, 13, 15, 94, 68, 84, 80, 58, 42, 37, 46, 85, 49, 97, 59, 44, 63, 92, 21, 87, 70, 81}
    bheap := NewBinaryHeapWithComparator(intComparator1)
    for i := 0; i < len(data); i++ {
        if bheap.Size() < k {
            bheap.Add(data[i])
        } else if data[i].(int) > bheap.Get().(int){
            bheap.Replace(data[i])
        }
    }
    s1 := make([]interface{}, 5)
    size := bheap.Size()
    for i := 0; i < size; i++ {
        s1[size - i - 1] = bheap.Remove()
    }
    fmt.Println(s1)
}