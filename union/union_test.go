package union

import (
	"testing"
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
