package sorts

import (
    "math/rand"
    "testing"
)

func TestInsertionSortAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    is := NewInsertionSort(true)
    is.SortInt(randInts)
    
    t.Log(randInts)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IntsAreAsSorted(randInts))
}

func TestInsertionSortAsSortIntSwap(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    is := NewInsertionSort(true)
    is.SortSwapInt(randInts)
    
    t.Log(randInts)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IntsAreAsSorted(randInts))
}

func BenchmarkInsertionSortAsSortInt(b *testing.B) {
    is := NewInsertionSort(true)
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        is.SortInt(randInts)
    }
}

func BenchmarkInsertionSortAsSortIntSwap(b *testing.B) {
    is := NewInsertionSort(true)
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        is.SortSwapInt(randInts)
    }
}

func TestInsertionSortAsSort(t *testing.T) {
    s1 := NewByAge(10000)
    t.Log(s1)
    is := NewInsertionSort(true)
    is.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IsAsSorted(s1))
}

func BenchmarkInsertionSortAsSort(b *testing.B) {
    is := NewInsertionSort(true)
    for i := 0; i < b.N; i++ {
        s1 := NewByAge(10000)
        is.Sort(s1)
    }
}


func TestInsertionSortAsSortSwap(t *testing.T) {
    s1 := NewByAge(10000)
    t.Log(s1)
    is := NewInsertionSort(true)
    is.SortSwap(s1)
    t.Log(s1)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IsAsSorted(s1))
}

func BenchmarkInsertionSortAsSortSwap(b *testing.B) {
    is := NewInsertionSort(true)
    for i := 0; i < b.N; i++ {
        s1 := NewByAge(10000)
        is.SortSwap(s1)
    }
}

func TestInsertionSortDsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    is := NewInsertionSort(false)
    is.SortInt(randInts)
    
    t.Log(randInts)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IntsAreDsSorted(randInts))
}

func TestInsertionSortDsSortIntSwap(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    is := NewInsertionSort(false)
    is.SortSwapInt(randInts)
    
    t.Log(randInts)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IntsAreDsSorted(randInts))
}

func TestInsertionSortDsSort(t *testing.T) {
    s1 := NewByAge(10000)
    t.Log(s1)
    is := NewInsertionSort(false)
    is.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IsDsSorted(s1))
}

func TestInsertionSortDsSortSwap(t *testing.T) {
    s1 := NewByAge(10000)
    t.Log(s1)
    is := NewInsertionSort(false)
    is.SortSwap(s1)
    t.Log(s1)
    t.Log("cmpCount: ", is.cmpCount)
    t.Log("swapCount: ", is.swapCount)
    t.Log(IsDsSorted(s1))

}
