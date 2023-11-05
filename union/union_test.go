package union

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestUnionQF(t *testing.T) {
	uf := NewUnionFindQF(12)
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 4)
	uf.Union(2, 3)
	uf.Union(2, 5)

	uf.Union(6, 7)

	uf.Union(8, 10)
	uf.Union(9, 10)
	uf.Union(9, 11)

	t.Log(uf.isSame(0, 6))
	uf.Union(4, 6)
	t.Log(uf.isSame(0, 6))
}

func TestUnionQU(t *testing.T) {
	uf := NewUnionFindQU(12)
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 4)
	uf.Union(2, 3)
	uf.Union(2, 5)

	uf.Union(6, 7)

	uf.Union(8, 10)
	uf.Union(9, 10)
	uf.Union(9, 11)

	t.Log(uf.isSame(0, 6))
	uf.Union(4, 6)
	t.Log(uf.isSame(0, 6))
}

func TestUnionQUS(t *testing.T) {
	count := 5000000
	// uf1 := NewUnionFindQF(count)
	// uf2 := NewUnionFindQU(count)
	uf3 := NewUnionFindQUS(count)
	uf4 := NewUnionFindQUR(count)
	// testTime(count, uf1)
	// testTime(count, uf2)
	testTime(count, uf3)
	testTime(count, uf4)
}

func testTime(count int, uf IUnion) {
	start := time.Now()
	for j := 0; j < count; j++ {
		uf.Union(rand.Intn(count), rand.Intn(count))
	}

	for j := 0; j < count; j++ {
		uf.isSame(rand.Intn(count), rand.Intn(count))
	}
	fmt.Println(time.Since(start))
}