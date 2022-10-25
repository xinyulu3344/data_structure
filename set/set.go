package set

type ISet[T comparable] interface {
    Size() int
    IsEmpty() bool
    Clear()
    Contains(e T) bool
    Add(e T)
    Remove(e T)
	Traversal(v Visit[T])
}


// Visit 遍历二叉树时执行的操作
// return 是否中止遍历，true终止，false不终止
type Visit[T comparable] func(e T) bool