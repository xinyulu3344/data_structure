package hashmap

type Key interface {
	HashCode() int
}

type IMap interface {
    Size() int
    IsEmpty() bool
    Clear()
    Put(key Key, value any)
    Get(key Key) any
    Remove(key Key) any
    ContainsKey(key Key) bool
    ContainsValue(value any) bool
    Traversal(v Visit)
}

type Visit func(key Key, value any) bool

type Compare func(k1, k2 Key) int

type Equals func(v1, v2 any) bool