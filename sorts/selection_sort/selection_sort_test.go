package selection_sort

import (
    "data_structure/sorts/utils"
    "math/rand"
    "testing"
)


func TestSelectionSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    t.Log(utils.IsAsSortedInts(randInts))
    
    SelectionSortInt(randInts)
    
    t.Log(randInts)
    t.Log(utils.IsAsSortedInts(randInts))
}


func BenchmarkSelectionSortInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        SelectionSortInt(randInts)
    }
/*
   goos: darwin
   goarch: arm64
   pkg: data_structure/sorts/selection_sort
   BenchmarkSelectionSortInt
   BenchmarkSelectionSortInt-8   	      36	  32255755 ns/op
   PASS
*/
}

