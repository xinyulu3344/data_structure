package heap

type IHeap interface {
	// 元素的数量
	Size() int
	// 是否为空
	IsEmpty() bool
	// 清空
	Clear()
	// 添加元素
	Add(e any) error
	// 获取堆顶元素
	Get() any
	// 删除堆顶元素
	Remove() any
	// 删除堆顶元素的同时插入一个新元素
	Replace(e any) (any, error)
}
