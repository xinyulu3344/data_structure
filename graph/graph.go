package graph

import (
    "fmt"
)

type Graph interface {
    GetEdgesSize() int
    GetVerticesSize() int
    AddVertex(v interface{})
    AddEdge(from, to interface{}, weight interface{})
    RemoveVertex(v interface{})
    RemoveEdge(from, to interface{})
}

type ListGraph struct {
    vertices map[interface{}]*Vertex // 顶点元素为key，顶点为value的map
    edges map[Edge]struct{}
}

func NewListGraph() *ListGraph {
    return &ListGraph{
        vertices: make(map[interface{}]*Vertex),
        edges: make(map[Edge]struct{}),
    }
}

func (lg *ListGraph) GetEdgesSize() int {
    return len(lg.edges)
}

func (lg *ListGraph) GetVerticesSize() int {
    return len(lg.vertices)
}

func (lg *ListGraph) AddVertex(v interface{}) {
    if _, ok := lg.vertices[v]; ok {
        return
    }
    lg.vertices[v] = NewVertex(v)
}

// AddEdge 添加一条边
func (lg *ListGraph) AddEdge(from, to interface{}, weight interface{}) {
    // 判断from、to是否存在
    fromVertex, ok := lg.vertices[from]
    if !ok {
        // 如果from不存在，就新增一个顶点，存到顶点集的map里
        fromVertex = NewVertex(from)
        lg.vertices[from] = fromVertex
    }
    toVertex, ok := lg.vertices[to]
    if !ok {
        // 如果to不存在，就新增一个顶点，存到顶点集的map里
        toVertex = NewVertex(to)
        lg.vertices[to] = toVertex
    }
    
    // 能来到这里，说明from、to都有值
    edge := NewEdge(fromVertex, toVertex, weight)
    fromVertex.outEdges[*edge] = struct{}{}
    toVertex.inEdges[*edge] = struct{}{}
    lg.edges[*edge] = struct{}{}
}

func (lg *ListGraph) RemoveVertex(v interface{}) {
    vertex := lg.vertices[v]
    if vertex == nil { return }
    
    // 删除顶点
    delete(lg.vertices, v)
    
    // 删除顶点关联的边
    for edge, _ := range vertex.outEdges {
        delete(edge.to.inEdges, edge)
        delete(edge.from.outEdges, edge)
        delete(lg.edges, edge)
    }
    for edge, _ := range vertex.inEdges {
        delete(edge.to.outEdges, edge)
        delete(edge.from.inEdges, edge)
        delete(lg.edges, edge)
    }
}

func (lg *ListGraph) RemoveEdge(from, to interface{}, weight interface{}) {
    // 判断from to是否存在
    fromVertex, ok := lg.vertices[from]
    if !ok { return }
    toVertex, ok := lg.vertices[to]
    if !ok { return }
    
    // 从edges中删除这条边
    edge := NewEdge(fromVertex, toVertex, weight)
    delete(lg.edges, *edge)
    // 从源顶点的出边集合中删除这条边
    delete(fromVertex.outEdges, *edge)
    // 从目标顶点的入边集合中删除这条边
    delete(toVertex.inEdges, *edge)
}

func (lg *ListGraph) Print() {
    for k, v := range lg.vertices {
        fmt.Println(k)
        fmt.Println("==> in")
        fmt.Println(v.inEdges)
        fmt.Println("==> out")
        fmt.Println(v.outEdges)
        fmt.Println()
    }
    for k, _ := range lg.edges {
        fmt.Println(k)
    }
}

// Vertex 顶点
type Vertex struct {
    Value    interface{}
    inEdges  map[Edge]struct{} // 以这个节点为起点的边的集合
    outEdges map[Edge]struct{} // 以这个节点为终点的边的集合
}

func NewVertex(v interface{}) *Vertex {
    return &Vertex{
        Value: v,
        inEdges: make(map[Edge]struct{}),
        outEdges: make(map[Edge]struct{}),
    }
}

func (v *Vertex) String() string {
    return fmt.Sprint(v.Value)
}

// Edge 边
type Edge struct {
    from   *Vertex
    to     *Vertex
    weight interface{}
}

func NewEdge(from, to *Vertex, weight interface{}) *Edge {
    return &Edge{
        from: from,
        to: to,
        weight: weight,
    }
}

func (e Edge) String() string {
    return "Edge [from=" + e.from.String() + ", to=" + e.to.String() + ", weight=" + fmt.Sprint(e.weight) + "]"
}