package union_find

type UnionFindQU struct {
    *UnionFind
}

func NewUnionFindQU(capacity int) *UnionFindQU {
    return &UnionFindQU{
        NewUnionFind(capacity),
    }
}

// Find 找到根元素
func (qu *UnionFindQU) Find(v int) int {
    qu.rangeCheck(v)
    
    // 找到根节点
    for v != qu.parents[v] {
        v = qu.parents[v]
    }
    return v
}

// Union 将v1的根节点，指向v2的根节点
func (qu *UnionFindQU) Union(v1, v2 int) {
    p1 := qu.Find(v1)
    p2 := qu.Find(v2)
    if p1 == p2 { return }
    
    qu.parents[p1] = p2
}

func (qu *UnionFindQU) IsSame(v1, v2 int) bool {
    return qu.Find(v1) == qu.Find(v2)
}