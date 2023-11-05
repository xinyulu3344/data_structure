// 路径分裂
package union

type UnionFindQURPathHalv struct {
    *UnionFindQUR
}

func NewUnionFindQURPathHalv(capacity int) *UnionFindQURPathHalv {
    uf := &UnionFindQURPathHalv{}
    uf.UnionFindQUR = NewUnionFindQUR(capacity)
    return uf
}

func (uf *UnionFindQURPathHalv) Find(v int) int {
    uf.rangeCheck(v)

    for v != uf.parents[v] {
        uf.parents[v] = uf.parents[uf.parents[v]]
        v = uf.parents[v]
    }
    return v
}