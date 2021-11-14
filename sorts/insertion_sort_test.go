package sorts

import (
    "data_structure/sorts/utils"
    "math/rand"
    "testing"
)

func TestInsertionSortAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    is := NewInsertionSort()
    is.AsSortInt(randInts)
    
    t.Log(randInts)
    t.Log(utils.IsAsSortedInts(randInts))
}

func BenchmarkInsertionSortAsSortInt(b *testing.B) {
    is := NewInsertionSort()
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        is.AsSortInt(randInts)
    }
}

func TestInsertionSortAsSort(t *testing.T) {
    s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
    t.Log(s1)
    is := NewInsertionSort()
    is.AsSort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
}

func TestNewInsertionSortWithComparator(t *testing.T) {
    s1 := []interface{}{47, 9, 38, 61, 73, 59, 52, 56, 27, 90}
    var intComparator *IntComparator
    is := NewInsertionSortWithComparator(intComparator)
    t.Log(s1)
    is.AsSort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
}