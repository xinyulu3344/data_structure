package rbtree

import (
	"testing"
)

type Int int

func (i Int) CompareTo(e E) int {
    return int(i - e.(Int))
}

func TestAdd(t *testing.T) {
    rbtree := NewRBTree()
    data := []Int{55, 87, 56, 74, 96, 22, 62, 20, 70, 68, 90, 50}
    for _, v := range data {
        rbtree.Add(v)
    }
}