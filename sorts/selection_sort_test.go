package sorts

import (
    "math/rand"
    "testing"
)

func TestSelectionAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    ss := NewSelectionSort(true)
    ss.SortInt(randInts)
    
    t.Log(randInts)
    t.Log(IntsAreAsSorted(randInts))
}

func BenchmarkSelectionAsSortInt(b *testing.B) {
    ss := NewSelectionSort(true)
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        ss.SortInt(randInts)
    }
}

func TestSelectionAsSort(t *testing.T) {
    s1 := NewByAge(10000)
    ss := NewSelectionSort(true)
    t.Log(s1)
    ss.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", ss.cmpCount)
    t.Log("swapCount: ", ss.swapCount)
    t.Log(IsAsSorted(s1))
}

func TestSelectionSortDsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    ss := NewSelectionSort(false)
    ss.SortInt(randInts)
    
    t.Log(randInts)
    t.Log(IntsAreDsSorted(randInts))
}

func TestSelectionSortDsSort(t *testing.T) {
    s1 := NewByAge(10000)
    ss := NewSelectionSort(false)
    t.Log(s1)
    ss.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", ss.cmpCount)
    t.Log("swapCount: ", ss.swapCount)
    t.Log(IsDsSorted(s1))
}