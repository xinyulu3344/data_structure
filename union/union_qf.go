package union

type IUnion interface {
	Find(v int) int
	Union(v1, v2 int)
	isSame(v1, v2 int) bool
}

type UnionFindQF struct {
	parents []int
}

func NewUnionFindQF(capacity int) IUnion {
	if capacity < 0 {
		return &UnionFindQF{
			parents: make([]int, 10, 10),
		}
	}
	unionFind := &UnionFindQF{
		parents: make([]int, capacity, capacity),
	}
	for i := 0; i < capacity; i++ {
		unionFind.parents[i] = i
	}
	return unionFind
}

func (uf *UnionFindQF) Find(v int) int {
	uf.rangeCheck(v)
	return uf.parents[v]
}

func (uf *UnionFindQF) Union(v1, v2 int) {
	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}
	for i := 0; i < len(uf.parents); i++ {
		if uf.parents[i] == p1 {
			uf.parents[i] = p2
		}
	}
}

func (uf *UnionFindQF) isSame(v1, v2 int) bool {
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFindQF) rangeCheck(v int) {
	if v > len(uf.parents) {
		panic("v is out of bounds")
	}
}
