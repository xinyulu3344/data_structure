package set

import (
	"fmt"
	"testing"
)

func TestTraversl(t *testing.T) {
    data := []int{11, 22, 33, 44, 11, 22, 33, 44}

    s := NewMapSet[int]()
    for _, v := range data {
        s.Add(v)
    }

    s.Traversal(func(e int) bool {
        fmt.Printf("%d ", e)
        return false
    })
}