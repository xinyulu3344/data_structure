package union

type node[V comparable] struct {
    value V
    parent *node[V]
    rank int
}

func NewNode[V comparable](v V) *node[V] {
    node := &node[V] {
        value: v,
        rank: 1,
    }
    node.parent = node
    return node
}

type GenericUnionFind[V comparable] struct {
    nodes map[V]*node[V]
}

func NewGenericUnionFind[V comparable]() *GenericUnionFind[V] {
    return &GenericUnionFind[V]{
        nodes: make(map[V]*node[V]),
    }
}

func (uf *GenericUnionFind[V]) MakeSet(v V) {
    if _, ok := uf.nodes[v]; ok {
        return
    }
    uf.nodes[v] = NewNode[V](v)
}

func (uf *GenericUnionFind[V]) findNode(v V) *node[V] {
    node, ok := uf.nodes[v]
    if !ok {
        return nil
    }
    for node.value != node.parent.value {
        node.parent = node.parent.parent
        node = node.parent
    }
    return node
}

func (uf *GenericUnionFind[V]) Find(v V) V {
    node := uf.findNode(v)
    if node == nil {
        return Zero[V]()
    }
    return node.value
}

func (uf *GenericUnionFind[V]) Union(v1, v2 V) {
    p1 := uf.findNode(v1)
    p2 := uf.findNode(v2)
    if p1 == nil || p2 == nil {
        return
    }
    if p1.value == p2.value {
        return
    }
    if p1.rank < p2.rank {
        p1.parent = p2
    } else if p1.rank > p2.rank {
        p2.parent = p1
    } else {
        p1.parent = p2
        p2.rank += 1
    }
}

func (uf *GenericUnionFind[V]) IsSame(v1, v2 V) bool {
    return uf.Find(v1) == uf.Find(v2)
}

func Zero[T comparable]() T {
    var zero T
    return zero
}