package sorts

import (
    "data_structure/sorts/utils"
    "math/rand"
    "testing"
)


func TestSelectionSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    t.Log(utils.IsAsSortedInts(randInts))
    
    ss := NewSelectionSort()
    ss.AsSortInt(randInts)
    
    t.Log(randInts)
    t.Log(utils.IsAsSortedInts(randInts))
}


func BenchmarkSelectionSortInt(b *testing.B) {
    ss := NewSelectionSort()
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        ss.AsSortInt(randInts)
    }
}

