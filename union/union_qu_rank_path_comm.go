// 路径压缩
package union

type UnionFindQURPathComm struct {
    *UnionFindQUR
}

func NewUnionFindQURPathComm(capacity int) *UnionFindQURPathComm {
    uf := &UnionFindQURPathComm{}
    uf.UnionFindQUR = NewUnionFindQUR(capacity)
    return uf
}

func (uf *UnionFindQURPathComm) Find(v int) int {
    uf.rangeCheck(v)

    if uf.parents[v] != v {
        uf.parents[v] = uf.Find(uf.parents[v])
    }
    return uf.parents[v]
}