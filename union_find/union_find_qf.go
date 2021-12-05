package union_find

type IUnionFind interface {
    Find(v int) int
    Union(v1, v2 int)
    IsSame(v1, v2 int) bool
}

type UnionFindQF struct {
    parents []int
}

func NewUnionFindQF(capacity int) *UnionFindQF {
    buf := make([]int, capacity)
    for i := 0; i < capacity; i++ {
        buf[i] = i
    }
    return &UnionFindQF{
        parents: buf,
    }
}

// Find 查找v所属的集合(根节点)
func (uf *UnionFindQF) Find(v int) int {
    uf.rangeCheck(v)
    return uf.parents[v]
}

func (uf *UnionFindQF) Union(v1, v2 int) {
    p1 := uf.Find(v1)
    p2 := uf.Find(v2)
    // 如果v1和v2是同一个集合，直接return
    if p1 == p2 { return }
    
    // 遍历数组，如果数组元素的根节点是p1，则将该元素的父节点修改为v2的父节点
    for i := 0; i < len(uf.parents); i++ {
        if uf.parents[i] == p1 {
            uf.parents[i] = p2
        }
    }
}

// IsSame 检查v1、v2是否属于同一个集合
func (uf *UnionFindQF) IsSame(v1, v2 int) bool {
    return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFindQF) rangeCheck(v int) {
    if v < 0 || v >= len(uf.parents) {
        panic("v is out of bounds")
    }
}
