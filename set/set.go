package set

type ISet interface {
    Size() int
    IsEmpty() bool
    Clear()
    Contains(e E) bool
    Add(e E)
    Remove(e E)
	Traversal(v Visit)
}