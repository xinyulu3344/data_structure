package rbtree

import (
	"fmt"
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
    for _, v := range data {
        n := rbtree.getNodeByElement(v)
        if n.color {
            t.Log(n.e, "Black")
        } else {
            t.Log(n.e, "Red")
        }
    }
}

func TestRemove(t *testing.T) {
    rbtree := NewRBTree()
    data := []Int{55, 87, 56, 74, 96, 22, 62, 20, 70, 68, 90, 50}
    for _, v := range data {
        rbtree.Add(v)
    }
    for _, v := range data {
        n := rbtree.getNodeByElement(v)
        if n.color {
            t.Log(n.e, "Black")
        } else {
            t.Log(n.e, "Red")
        }
    }
    fmt.Println("===")
    for i, v := range data {
        rbtree.Remove(v)
        for j := i+1; j < len(data); j++ {
            n := rbtree.getNodeByElement(data[j])
            if n.color {
                t.Log(n.e, "Black")
            } else {
                t.Log(n.e, "Red")
            }
        }
        fmt.Println("===")
    }
}