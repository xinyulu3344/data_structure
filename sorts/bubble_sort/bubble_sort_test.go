package bubble_sort

import (
    "fmt"
    "testing"
)

type IntComparator int

func (i *IntComparator) CompareTo(e1 interface{}, e2 interface{}) int {
    return e1.(int) - e2.(int)
}

type Integer int

func (i Integer) Compare(e interface{}) int {
    return int(i - e.(Integer))
}

func TestBubbleSort(t *testing.T) {
    s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
    BubbleSort(s1)
    fmt.Println(s1)
}

func BenchmarkBubbleSort(b *testing.B) {
    s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
    for i := 0; i < b.N; i++ {
        BubbleSort(s1)
    }
}

func BenchmarkBubbleSortWithComparator(b *testing.B) {
    var intComparator *IntComparator
    s1 := []interface{}{47, 9, 38, 61, 73, 59, 52, 56, 27, 90}
    for i := 0; i < b.N; i++ {
        BubbleSortWithComparator(s1, intComparator)
    }
}
