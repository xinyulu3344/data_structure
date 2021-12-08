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

type Value interface {
    IsEqual(Value) bool
}

type Node struct {
    value Value
    parent *Node
    rank int
}

func NewNode(v Value) *Node {
    node := &Node {
        value: v,
    }
    node.parent = node
    return node
}

type GenericUnionFind struct {
    nodes map[Value]*Node
}

func NewGenericUnionFind() *GenericUnionFind {
    return &GenericUnionFind{
        nodes: make(map[Value]*Node),
    }
}

func (g *GenericUnionFind) MakeSet(v Value) {
    if _, ok := g.nodes[v]; ok { return }
    g.nodes[v] = NewNode(v)
}

func (g *GenericUnionFind) Find(v Value) Value {
    root := g.findNode(v)
    if root == nil { return nil }
    return root.value
}

func (g *GenericUnionFind) Union(v1, v2 Value) {
    p1 := g.findNode(v1)
    p2 := g.findNode(v2)
    if p1 == nil || p2 == nil { return }
    if p1.value.IsEqual(p2.value) { return }
    if p1.rank < p2.rank {
        p1.parent = p2
    } else if p1.rank > p2.rank {
        p2.parent = p1
    } else {
        p1.parent = p2
        p2.rank += 1
    }
}

func (g *GenericUnionFind) IsSame(v1, v2 Value) bool {
    return g.Find(v1).IsEqual(g.Find(v2))
}

// 找到v的根节点
func (g *GenericUnionFind) findNode(v Value) *Node {
    node := g.nodes[v]
    if node == nil { return nil }

    for !node.value.IsEqual(node.parent.value) {
        node.parent = node.parent.parent
        node = node.parent
    }
    return node
}