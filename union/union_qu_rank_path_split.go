// 路径分裂
package union

type UnionFindQURPathSplit struct {
    *UnionFindQUR
}

func NewUnionFindQURPathSplit(capacity int) *UnionFindQURPathSplit {
    uf := &UnionFindQURPathSplit{}
    uf.UnionFindQUR = NewUnionFindQUR(capacity)
    return uf
}

func (uf *UnionFindQURPathSplit) Find(v int) int {
    uf.rangeCheck(v)

    for v != uf.parents[v] {
        p := uf.parents[v]
        uf.parents[v] = uf.parents[uf.parents[v]]
        v = p
    }
    return v
}