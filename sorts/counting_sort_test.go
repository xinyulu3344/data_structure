package sorts

import (
    "math/rand"
    "testing"
)

func TestCountingSort_Sort(t *testing.T) {
    cs := NewCountingSort()
    data := rand.Perm(100000000)
    //t.Log(data)
    cs.Sort(data)
    //t.Log(data)
    t.Log(IntsAreAsSorted(data))
}