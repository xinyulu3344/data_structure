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