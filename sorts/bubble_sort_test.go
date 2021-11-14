package sorts

import (
    "data_structure/sorts/utils"
    "math/rand"
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
    bs := NewBubbleSort()
    t.Log(s1)
    bs.AsSort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", bs.cmpCount)
    t.Log("swapCount: ", bs.swapCount)
}

func TestBubbleSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    bs := NewBubbleSort()
    bs.AsSortInt(randInts)
    t.Log(randInts)
    t.Log("cmpCount: ", bs.cmpCount)
    t.Log("swapCount: ", bs.swapCount)
    t.Log(utils.IsAsSortedInts(randInts))
}

func BenchmarkBubbleSortInt(b *testing.B) {
    bs := NewBubbleSort()
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        bs.AsSortInt(randInts)
    }
}

func BenchmarkBubbleSort(b *testing.B) {
    bs := NewBubbleSort()
    for i := 0; i < b.N; i++ {
        s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
        bs.AsSort(s1)
    }
}

func BenchmarkBubbleSortWithComparator(b *testing.B) {
    var intComparator *IntComparator
    bs := NewBubbleSortWithComparator(intComparator)
    for i := 0; i < b.N; i++ {
        s1 := []interface{}{47, 9, 38, 61, 73, 59, 52, 56, 27, 90}
        bs.AsSort(s1)
    }
}