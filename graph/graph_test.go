package graph

import "testing"

func InitGraph() *ListGraph {
    graph := NewListGraph()
    graph.AddEdge("v1", "v0", 9)
    graph.AddEdge("v1", "v2", 3)
    graph.AddEdge("v2", "v0", 2)
    graph.AddEdge("v2", "v3", 5)
    graph.AddEdge("v3", "v4", 1)
    graph.AddEdge("v0", "v4", 8)
    graph.AddEdge("v0", "v4", 6)
    return graph
}

func TestListGraph_AddEdge(t *testing.T) {
    graph := InitGraph()
    graph.Print()
}

func TestListGraph_RemoveEdge(t *testing.T) {
    graph := InitGraph()
    graph.RemoveEdge("v0", "v4", 6)
    graph.Print()
}

func TestListGraph_RemoveVertex(t *testing.T) {
    graph := InitGraph()
    t.Log("初始顶点数量: ", graph.GetVerticesSize())
    t.Log("初始边数量: ", graph.GetEdgesSize())
    graph.RemoveVertex("v0")
    t.Log("删除v0后顶点数量: ", graph.GetVerticesSize())
    t.Log("删除v0后边数量: ", graph.GetEdgesSize())
    graph.Print()
}
