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

type Student struct {
	name string
	age int
}

func NewStudent(name string, age int) *Student {
	return &Student{
		name: name,
		age: age,
	}
}

func (s *Student) IsEqual(v Value) bool {
	stu := v.(*Student)
	if s.name == stu.name && s.age == stu.age {
		return true
	} else {
		return false
	}
}

func TestGenericUnionFind(t *testing.T) {
	uf := NewGenericUnionFind()
	stu1 := NewStudent("jack", 1)
	stu2 := NewStudent("rose", 2)
	stu3 := NewStudent("jack", 3)
	stu4 := NewStudent("rose", 4)
	uf.MakeSet(stu1)
	uf.MakeSet(stu2)
	uf.MakeSet(stu3)
	uf.MakeSet(stu4)
	
	uf.Union(stu1, stu2)
	uf.Union(stu3, stu4)
	uf.Union(stu1, stu4)

	t.Log(uf.IsSame(stu1, stu2))
	t.Log(uf.IsSame(stu1, stu3))
	t.Log(uf.IsSame(stu1, stu4))
	t.Log(uf.IsSame(stu2, stu3))
	t.Log(uf.IsSame(stu2, stu4))
	t.Log(uf.IsSame(stu3, stu4))
}