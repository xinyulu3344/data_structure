/*
堆的一个重要性质：任意节点的值总是>=或<=子节点的值
>=: 最大堆、大根堆、大顶堆
<=: 最小堆、小根堆、小顶堆
*/
package myheap

type Heap interface {
    Size() int                         // 元素数量
    IsEmpty() bool                     // 是否为空
    Clear()                            // 清空
    Add(e interface{})                 // 添加元素
    Get() interface{}                  // 获取堆顶元素
    Remove() interface{}               // 删除堆顶元素
    Replace(e interface{}) interface{} // 删除堆顶元素的同时插入一个新元素
}