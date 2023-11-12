package union

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenericUnionFind(t *testing.T) {
    uf := NewGenericUnionFind[int]()
    uf.MakeSet(1)
    uf.MakeSet(2)
    uf.MakeSet(3)
    uf.MakeSet(4)
    uf.MakeSet(5)
    uf.MakeSet(6)
    uf.MakeSet(7)
    uf.MakeSet(8)
    uf.MakeSet(9)
    uf.MakeSet(10)
    uf.MakeSet(11)



    fmt.Printf("uf.IsSame(4, 6): %v\n", uf.IsSame(4, 6))
	uf.Union(4, 6)
    fmt.Printf("uf.IsSame(4, 6): %v\n", uf.IsSame(4, 6))
    testGenericUnionFindTime[int](5000000, uf)
}

func TestGenericUnionFindTime(t *testing.T) {
    uf := NewGenericUnionFind[int]()
    testGenericUnionFindTime[int](5000000, uf)
}

func testGenericUnionFindTime[T comparable](count int, uf *GenericUnionFind[int]) {
	for j := 0; j < count; j++ {
        uf.MakeSet(rand.Intn(count))
	}

	start := time.Now()
    for j := 0; j < count; j++ {
		uf.Union(rand.Intn(count), rand.Intn(count))
    }

	for j := 0; j < count; j++ {
		uf.IsSame(rand.Intn(count), rand.Intn(count))
	}
	fmt.Println(time.Since(start))
}