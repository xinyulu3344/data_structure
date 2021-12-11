package union_find

import (
    "math/rand"
    "testing"
    "time"
)

func TestNewUnionFindQF(t *testing.T) {
    uf := NewUnionFindQF(12)
    uf.Union(0, 1)
    uf.Union(0, 3)
    uf.Union(0, 4)
    uf.Union(2, 3)
    uf.Union(2, 5)
    uf.Union(6, 7)
    uf.Union(8, 10)
    uf.Union(9, 10)
    uf.Union(9, 11)
    t.Log(uf.IsSame(0, 6))
    t.Log(uf.IsSame(0, 5))
    
    t.Log(uf.IsSame(2, 7))
    uf.Union(4, 6)
    t.Log(uf.IsSame(2, 7))
}

func TestNewUnionFindQU(t *testing.T) {
    uf := NewUnionFindQU(12)
    uf.Union(0, 1)
    uf.Union(0, 3)
    uf.Union(0, 4)
    uf.Union(2, 3)
    uf.Union(2, 5)
    uf.Union(6, 7)
    uf.Union(8, 10)
    uf.Union(9, 10)
    uf.Union(9, 11)
    t.Log(uf.IsSame(0, 6))
    t.Log(uf.IsSame(0, 5))
    
    t.Log(uf.IsSame(2, 7))
    uf.Union(4, 6)
    t.Log(uf.IsSame(2, 7))
}

func TestNewUnionFindQUSize(t *testing.T) {
    uf := NewUnionFindQUSize(12)
    uf.Union(0, 1)
    uf.Union(0, 3)
    uf.Union(0, 4)
    uf.Union(2, 3)
    uf.Union(2, 5)
    uf.Union(6, 7)
    uf.Union(8, 10)
    uf.Union(9, 10)
    uf.Union(9, 11)
    t.Log(uf.IsSame(0, 6))
    t.Log(uf.IsSame(0, 5))
    
    t.Log(uf.IsSame(2, 7))
    uf.Union(4, 6)
    t.Log(uf.IsSame(2, 7))
}

func TestUnionFindQUSize_Union(t *testing.T) {
    count := 100000
    uf := NewUnionFindQUSize(count)
    start := time.Now()
    for i := 0; i < count; i++ {
        uf.Union(rand.Intn(count), rand.Intn(count))
    }
    t.Log("UnionFindQUSize_Union: ", time.Since(start))
    start = time.Now()
    for i := 0; i < count; i++ {
        uf.IsSame(rand.Intn(count), rand.Intn(count))
    }
    t.Log("UnionFindQUSize_IsSame: ", time.Since(start))
}

func TestUnionFindQU_Union(t *testing.T) {
    count := 100000
    uf := NewUnionFindQU(count)
    start := time.Now()
    for i := 0; i < count; i++ {
        uf.Union(rand.Intn(count), rand.Intn(count))
    }
    t.Log("UnionFindQU_Union: ", time.Since(start))
    start = time.Now()
    for i := 0; i < count; i++ {
        uf.IsSame(rand.Intn(count), rand.Intn(count))
    }
    t.Log("UnionFindQU_IsSame: ", time.Since(start))
}

func TestUnionFindQF_Union(t *testing.T) {
    count := 100000
    uf := NewUnionFindQF(count)
    start := time.Now()
    for i := 0; i < count; i++ {
        uf.Union(rand.Intn(count), rand.Intn(count))
    }
    t.Log("UnionFindQF_Union: ", time.Since(start))
    start = time.Now()
    for i := 0; i < count; i++ {
        uf.IsSame(rand.Intn(count), rand.Intn(count))
    }
    t.Log("UnionFindQF_IsSame: ", time.Since(start))
}