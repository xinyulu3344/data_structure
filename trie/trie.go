package trie

type Trie interface {
    Size() int
    IsEmpty() bool
    Clear()
    Get(key string) interface{}
    Contains(key string) bool
    Add(key string, value interface{}) interface{}
    Remove(key string) interface{}
    StartWith(prefix string) bool
}
