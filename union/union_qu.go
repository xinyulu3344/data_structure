package union

type UnionFindQU struct {
	parents []int
}

func NewUnionFindQU(capacity int) *UnionFindQU {
	if capacity < 0 {
		return &UnionFindQU{
			parents: make([]int, 10, 10),
		}
	}
	unionFind := &UnionFindQU{
		parents: make([]int, capacity, capacity),
	}
	for i := 0; i < capacity; i++ {
		unionFind.parents[i] = i
	}
	return unionFind
}

func (uf *UnionFindQU) Find(v int) int {
	uf.rangeCheck(v)
	for v != uf.parents[v] {
		v = uf.parents[v]
	}
	return v
}

func (uf *UnionFindQU) Union(v1 int, v2 int) {
	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}
	uf.parents[p1] = p2
}

func (uf *UnionFindQU) isSame(v1 int, v2 int) bool {
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFindQU) rangeCheck(v int) {
	if v > len(uf.parents) {
		panic("v is out of bounds")
	}
}
