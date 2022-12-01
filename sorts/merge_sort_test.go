package sorts

import (
    "math/rand"
    "testing"
)

func TestMergeSort_Sort(t *testing.T) {
    data := make([]interface{}, 10000)
    for i := 0; i < 10000; i++ {
        data[i] = Integer(rand.Intn(10000))
    }
    ms := NewMergeSort()
    t.Log(data)
    ms.Sort(data)
    t.Log(data)
}

func TestMergeSortWithComparator_Sort(t *testing.T) {
    data := make([]interface{}, 10000)
    for i := 0; i < 10000; i++ {
        data[i] = Integer(rand.Intn(10000))
    }
    var cmp Integer
    ms := NewMergeSortWithComparator(cmp)
    t.Log(data)
    ms.Sort(data)
    t.Log(data)
}