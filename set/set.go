package set

type Set interface {
    Size() int
    IsEmpty() bool
    Clear()
    Contains(e interface{}) bool
    Add(e interface{})
    Remove(e interface{})
    Traversal()
}


type Visitor interface {
    Visit(e interface{}) // 操作遍历的数据
    Stop() bool          // 是否终止遍历
}
