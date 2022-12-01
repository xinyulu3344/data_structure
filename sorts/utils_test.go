package sorts

import (
    "testing"
)

func TestIndexOfInt(t *testing.T) {
    orderInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    t.Log(IndexOfInt(orderInts, 3))
}

func TestSearch(t *testing.T) {
    orderInts := []int{2, 4, 8, 8, 8, 12, 14}
    t.Log(Search(orderInts, 5) == 2)
    t.Log(Search(orderInts, 1) == 0)
    t.Log(Search(orderInts, 15) == 7)
    t.Log(Search(orderInts, 8) == 5)
}