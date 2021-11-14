package sorts

import (
    "data_structure/sorts/utils"
    "math/rand"
    "testing"
)


func TestHeapSortAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    
    t.Log(randInts)
    hs := NewHeapSort()
    hs.AsSortInt(randInts)
    t.Log(randInts)
    t.Log("cmpCount: ", hs.cmpCount)
    t.Log("swapCount: ", hs.swapCount)
    
    t.Log(utils.IsAsSortedInts(randInts))
}

func BenchmarkHeapSortAsSort(b *testing.B) {
    hs := NewHeapSort()
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        hs.AsSortInt(randInts)
    }
}

func TestHeapSortAsSort(t *testing.T) {
    hs := NewHeapSort()
    s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
    t.Log(s1)
    hs.AsSort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", hs.cmpCount)
    t.Log("swapCount: ", hs.swapCount)
}