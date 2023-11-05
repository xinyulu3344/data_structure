// 基于rank的优化
package union


type UnionFindQUR struct {
    *UnionFindQU
    ranks []int
}

func NewUnionFindQUR(capacity int) *UnionFindQUR {
    uf := &UnionFindQUR{}
    uf.UnionFindQU = NewUnionFindQU(capacity)

    uf.ranks = make([]int, capacity, capacity)

    for i := 0; i < len(uf.ranks); i++ {
        uf.ranks[i] = 1
    }
    return uf
}


func (uf *UnionFindQUR) Union(v1, v2 int) {
    p1 := uf.Find(v1)
    p2 := uf.Find(v2)

    if p1 == p2 {
        return
    }

    if uf.ranks[p1] < uf.ranks[p2] {
        uf.parents[p1] = p2
    } else if uf.ranks[p1] > uf.ranks[p2] {
        uf.parents[p2] = p1
    } else {
        uf.parents[p1] = p2
        uf.ranks[p2] += 1
    }
}