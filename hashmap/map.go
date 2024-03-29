package hashmap

type Key interface {
	HashCode() int
    Equals(key Key) bool
}

type IMap interface {
    Size() int
    IsEmpty() bool
    Clear()
    Put(key Key, value any)
    Get(key Key) any
    Remove(key Key) any
    ContainsKey(key Key) bool
    ContainsValue(value any, equals Equals) bool
    Traversal(visit Visit)
}

type Visit func(key Key, value any) bool

type Compare func(k1, k2 Key) int

type Equals func(v1, v2 any) bool