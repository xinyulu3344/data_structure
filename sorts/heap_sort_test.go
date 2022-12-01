package sorts

import (
    "math/rand"
    "testing"
)


func TestHeapSortAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    
    t.Log(randInts)
    hs := NewHeapSort(true)
    hs.SortInt(randInts)
    t.Log(randInts)
    t.Log("cmpCount: ", hs.cmpCount)
    t.Log("swapCount: ", hs.swapCount)
    
    t.Log(IntsAreAsSorted(randInts))
}

func BenchmarkHeapSortAsSort(b *testing.B) {
    hs := NewHeapSort(true)
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        hs.SortInt(randInts)
    }
}

func TestHeapSortAsSort(t *testing.T) {
    hs := NewHeapSort(true)
    //s1 := []interface{}{Integer(47), Integer(9), Integer(38), Integer(61), Integer(73), Integer(59), Integer(52), Integer(56), Integer(27), Integer(90)}
    s1 := NewByAge(10000)
    t.Log(s1)
    hs.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", hs.cmpCount)
    t.Log("swapCount: ", hs.swapCount)
    t.Log(IsAsSorted(s1))
}

func TestHeapSortDsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    
    t.Log(randInts)
    hs := NewHeapSort(false)
    hs.SortInt(randInts)
    t.Log(randInts)
    t.Log("cmpCount: ", hs.cmpCount)
    t.Log("swapCount: ", hs.swapCount)
    
    t.Log(IntsAreDsSorted(randInts))
}