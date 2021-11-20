package sorts

import (
    "math/rand"
    "testing"
)

func TestSelectionAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    t.Log(IsAsSortedInts(randInts))
    
    ss := NewSelectionSort()
    ss.AsSortInt(randInts)
    
    t.Log(randInts)
    t.Log(IsAsSortedInts(randInts))
}

func BenchmarkSelectionSortInt(b *testing.B) {
    ss := NewSelectionSort()
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        ss.AsSortInt(randInts)
    }
}

func TestNewSelectionAsSort(t *testing.T) {
    s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
    ss := NewSelectionSort()
    t.Log(s1)
    ss.AsSort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", ss.cmpCount)
    t.Log("swapCount: ", ss.swapCount)
}

func TestNewSelectionSortWithComparator(t *testing.T) {
    s1 := []interface{}{47, 9, 38, 61, 73, 59, 52, 56, 27, 90}
    var intComparator *IntComparator
    ss := NewSelectionSortWithComparator(intComparator)
    t.Log(s1)
    ss.AsSort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", ss.cmpCount)
    t.Log("swapCount: ", ss.swapCount)
}