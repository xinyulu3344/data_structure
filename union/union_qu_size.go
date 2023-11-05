// 基于size的优化
package union


type UnionFindQUS struct {
    *UnionFindQU
    sizes []int
}

func NewUnionFindQUS(capacity int) *UnionFindQUS {
    uf := &UnionFindQUS{}
    uf.UnionFindQU = NewUnionFindQU(capacity)

    uf.sizes = make([]int, capacity, capacity)

    for i := 0; i < len(uf.sizes); i++ {
        uf.sizes[i] = 1
    }
    return uf
}


func (uf *UnionFindQUS) Union(v1, v2 int) {
    p1 := uf.Find(v1)
    p2 := uf.Find(v2)

    if p1 == p2 {
        return
    }

    if uf.sizes[p1] < uf.sizes[p2] {
        uf.parents[p1] = p2
        uf.sizes[p2] += uf.sizes[p1]
    } else {
        uf.parents[p2] = p1
        uf.sizes[p1] += uf.sizes[p2]
    }
}