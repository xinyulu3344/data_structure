package hashmap

import (
	"fmt"
	"testing"
)

type Person struct {
    Name string
    Age int
}

func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age: age,
    }
}

func (p *Person) HashCode() int {
    hash := 0
    length := len(p.Name)
    for i := 0; i < length; i++ {
        c := int(p.Name[i])
        hash = 31 * hash + c
    }
    hash = 31 * hash + p.Age
    return hash
}

func (p *Person) Equals(key Key) bool {
    p1 := key.(*Person)
    if p.Name == p1.Name && p.Age == p1.Age {
        return true
    }
    return false
}

type Str string

func (s Str) HashCode() int {
    h := 0
    for i := 0; i < len(s); i++ {
        h = 31 * h + int(s[i])
    }
    return h
}

func (s Str) Equals(key Key) bool {
    return s == key.(Str)
}

type Int int

func (i Int) HashCode() int {
    return int(i)
}

func (i Int) Equals(key Key) bool {
    return i == key.(Int)
}

func TestXxx(t *testing.T) {
    p1 := NewPerson("jack", 1)
    p2 := NewPerson("rose", 2)
    hm := NewHashMap()
    hm.Put(p1, 1)
    hm.Put(p2, 2)
    fmt.Printf("hm.Size(): %v\n", hm.Size())
    fmt.Printf("hm.ContainsKey(p1): %v\n", hm.ContainsKey(p2))
}

func TestRemove(t *testing.T) {
    p1 := NewPerson("jack", 1)
    p2 := NewPerson("rose", 2)
    p3 := NewPerson("tom", 3)
    hm := NewHashMap()
    hm.Put(p1, 1)
    hm.Put(p2, 2)
    hm.Put(p3, 3)

    fmt.Printf("hm.Size(): %v\n", hm.Size())
    fmt.Printf("hm.ContainsKey(p1): %v\n", hm.ContainsKey(p1))
    fmt.Printf("hm.Get(p1): %v\n", hm.Get(p1))

    fmt.Printf("hm.Remove(p1): %v\n", hm.Remove(p1))
    fmt.Printf("hm.Size(): %v\n", hm.Size())
    fmt.Printf("hm.ContainsKey(p1): %v\n", hm.ContainsKey(p1))
    fmt.Printf("hm.Get(p1): %v\n", hm.Get(p1))
}

type PersonJob struct {
    JobName string
}

func NewPersonJob(jobName string) *PersonJob {
    return &PersonJob{
        JobName: jobName,
    }
}

func TestContainsValue(t *testing.T) {
    p1 := NewPerson("jack", 1)
    p2 := NewPerson("rose", 2)
    p3 := NewPerson("tom", 3)

    hm := NewHashMap()
    hm.Put(p1, NewPersonJob("IT"))
    hm.Put(p2, NewPersonJob("Internet"))
    hm.Put(p3, NewPersonJob("Car"))

    r := hm.ContainsValue(NewPersonJob("Car"), func(v1, v2 any) bool {
        return v1.(*PersonJob).JobName == v2.(*PersonJob).JobName
    })
    t.Log(r) // true
}

func TestTraversal(t *testing.T) {
    p1 := NewPerson("jack", 1)
    p2 := NewPerson("rose", 2)
    p3 := NewPerson("tom", 3)

    hm := NewHashMap()
    hm.Put(p1, NewPersonJob("IT"))
    hm.Put(p2, NewPersonJob("Internet"))
    hm.Put(p3, NewPersonJob("Car"))

    hm.Traversal(func(key Key, value any) bool {
        t.Log(key, value)
        return false
    })
}

type K struct {
    value int
}

func NewK(value int) *K {
    return &K{
        value: value,
    }
}

func (k *K) HashCode() int {
    return k.value / 9
}

func (k *K) Equals(key Key) bool {
    return k.value == key.(*K).value
}

func TestPut(t *testing.T) {
    hm := NewHashMap()
    for i := 0; i < 9; i++ {
        hm.Put(NewK(i), i)
    }

    hm.Traversal(func(key Key, value any) bool {
        t.Log(key, value)
        return false
    })
}

func Test1(t *testing.T) {
    hm := NewHashMap()
    for i := 1; i <= 20; i++ {
        hm.Put(NewK(i), i)
    }
    for i := 5; i <= 7; i++ {
        hm.Put(NewK(i), i + 5)
    }
    assert(t, hm.Size() == 20)
    assert(t, hm.Get(NewK(4)) == 4)
    assert(t, hm.Get(NewK(5)) == 10)
    assert(t, hm.Get(NewK(6)) == 11)
    assert(t, hm.Get(NewK(7)) == 12)
    assert(t, hm.Get(NewK(8)) == 8)
}

func Test2(t *testing.T) {
    hm := NewHashMap()
    hm.Put(Str("jack"), 4)
    hm.Put(Int(10), 4)
    hm.Put(Str("jack"), 6)
    hm.Put(Int(10), nil)
    hm.Put(NewK(10), 8)
    assert(t, hm.Size() == 3)
    assert(t, hm.Get(Str("jack")) == 6)
    assert(t, hm.Get(Int(10)) == nil)
    assert(t, hm.ContainsKey(Int(10)))
    assert(t, hm.ContainsValue(8, func(v1, v2 any) bool {
        return v1 == v2
    }))
    assert(t, !hm.ContainsValue(1, func(v1, v2 any) bool {
        return v1 == v2
    }))
}

func Test3(t *testing.T) {
    hm := NewHashMap()
    hm.Put(Str("jack"), 1)
    hm.Put(Str("rose"), 2)
    hm.Put(Str("jim"), 3)
    hm.Put(Str("jake"), 4)
    for i := 1; i <= 10; i++ {
        hm.Put(Str(fmt.Sprintf("%s%d", "test", i)), i)
        hm.Put(NewK(i), i)
    }
    for i := 5; i <= 7; i++ {
        assert(t, hm.Remove(NewK(i)) == i)
    }
    for i := 1; i <= 3; i++ {
        hm.Put(NewK(i), i + 5)
    }
    assert(t, hm.Size() == 21)
}

func assert(t *testing.T, ok bool) {
    if !ok {
        t.Fatal()
    }
}

func TestResize(t *testing.T) {
    hm := NewHashMap()
    for i := 0; i < 10000000; i++ {
        hm.Put(Int(i), i)
    }
}

func TestMap(t *testing.T) {
    m := make(map[Int]int)
    for i := 0; i < 1000000; i++ {
        m[Int(i)] = i
    }
}