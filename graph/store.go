package graph

import "sync"

type Store[K comparable, T any] interface {
    AddVertex(hash K, value T, properties VertexProperties) error
    Vertex(hash K) (T, VertexProperties, error)
    RemoveVertex(hash K) error
    ListVertices() ([]K, error)
    VertexCount() (int, error)

    AddEdge(sourceHash, targetHash K, edge Edge[K]) error
    UpdateEdge(sourceHash, targetHash K, edge Edge[K]) error
    RemoveEdge(sourceHash, targetHash K) error
    Edge(sourceHash, targetHash K) (Edge[K], error)
    ListEdges() ([]Edge[K], error)
}

type memoryStore[K comparable, T any] struct {
    lock sync.RWMutex
    vertices map[K]T
    vertexProperties map[K]VertexProperties
    outEdges map[K]map[K]Edge[K]
    inEdges map[K]map[K]Edge[K]
}

func newMemoryStore[K comparable, T any]() Store[K, T] {
    return &memoryStore[K, T] {
        vertices: make(map[K]T),
        vertexProperties: make(map[K]VertexProperties),
        outEdges: make(map[K]map[K]Edge[K]),
        inEdges: make(map[K]map[K]Edge[K]),
    }
}

func (s *memoryStore[K, T]) AddVertex(hash K, value T, properties VertexProperties) error {
    s.lock.Lock()
    defer s.lock.Unlock()

    if _, ok := s.vertices[hash]; ok {
        return ErrVertexAlreadyExists
    }

    s.vertices[hash] = value
    s.vertexProperties[hash] = properties

    return nil
}

func (s *memoryStore[K, T]) Vertex(hash K) (T, VertexProperties, error) {
    s.lock.RLock()
    defer s.lock.RUnlock()

    v, ok := s.vertices[hash]
    if !ok {
        return v, VertexProperties{}, ErrVertexNotFound
    }
    
    p := s.vertexProperties[hash]

    return v, p, nil
}

func (s *memoryStore[K, T]) RemoveVertex(hash K) error {
    s.lock.RLock()
    defer s.lock.RUnlock()

    if _, ok := s.vertices[hash]; !ok {
        return ErrVertexNotFound
    }

    if edges, ok := s.inEdges[hash]; ok {
        if len(edges) > 0 {
            return ErrVertexHasEdges
        }
        delete(s.inEdges, hash)
    }

    if edges, ok := s.outEdges[hash]; ok {
        if len(edges) > 0 {
            return ErrVertexHasEdges
        }
        delete(s.outEdges, hash)
    }

    delete(s.vertices, hash)
    delete(s.vertexProperties, hash)

    return nil
}

func (s *memoryStore[K, T]) ListVertices() ([]K, error) {
    s.lock.RLock()
    defer s.lock.RUnlock()

    hashes := make([]K, 0, len(s.vertices))

    for k := range s.vertices {
        hashes = append(hashes, k)
    }

    return hashes, nil
}

func (s *memoryStore[K, T]) VertexCount() (int, error) {
    s.lock.RLock()
    s.lock.RUnlock()

    return len(s.vertices), nil
}

func (s *memoryStore[K, T]) AddEdge(sourceHash K, targetHash K, edge Edge[K]) error {
    s.lock.Lock()
    defer s.lock.Unlock()

    if _, ok := s.outEdges[sourceHash]; !ok {
        s.outEdges[sourceHash] = make(map[K]Edge[K])
    }

    s.outEdges[sourceHash][targetHash] = edge

    if _, ok := s.inEdges[targetHash]; ok {
        s.inEdges[targetHash] = make(map[K]Edge[K])
    }

    s.inEdges[targetHash][sourceHash] = edge

    return nil
}

func (s *memoryStore[K, T]) UpdateEdge(sourceHash K, targetHash K, edge Edge[K]) error {
    if _, err := s.Edge(sourceHash, targetHash); err != nil {
        return err
    }

    s.lock.Lock()
    defer s.lock.Unlock()

    s.outEdges[sourceHash][targetHash] = edge
    s.inEdges[targetHash][sourceHash] = edge

    return nil
}

func (s *memoryStore[K, T]) RemoveEdge(sourceHash K, targetHash K) error {
    s.lock.Lock()
    defer s.lock.Unlock()

    delete(s.outEdges[sourceHash], targetHash)
    delete(s.inEdges[targetHash], sourceHash)

    return nil
}

func (s *memoryStore[K, T]) Edge(sourceHash K, targetHash K) (Edge[K], error) {
    s.lock.RLock()
    defer s.lock.RUnlock()

    sourceEdges, ok := s.outEdges[sourceHash]
    if !ok {
        return Edge[K]{}, ErrEdgeNotFound
    }

    edge, ok := sourceEdges[targetHash]
    if !ok {
        return Edge[K]{}, ErrEdgeNotFound
    }

    return edge, nil
}

func (s *memoryStore[K, T]) ListEdges() ([]Edge[K], error) {
    s.lock.RLock()
    defer s.lock.RUnlock()

    res := make([]Edge[K], 0)
    for _, edges := range s.outEdges {
        for _, edge := range edges {
            res = append(res, edge)
        }
    }
    return res, nil
}

