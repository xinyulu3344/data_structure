package union_find

type UnionFind struct {
    parents []int
}

func NewUnionFind(capacity int) *UnionFind {
    buf := make([]int, capacity)
    for i := 0; i < capacity; i++ {
        buf[i] = i
    }
    return &UnionFind{
        parents: buf,
    }
}

func (uf *UnionFind) rangeCheck(v int) {
    if v < 0 || v >= len(uf.parents) {
        panic("v is out of bounds")
    }
}