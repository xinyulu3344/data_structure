package array_list

import (
    "fmt"
)

const (
    DEFAULT_CAPACITY  = 10
    ELEMENT_NOT_FOUND = -1
)

type Item interface {
    Equal(e Item) bool
}

type IArrayList interface {
    Size() int
    IsEmpty() bool
    Contains(e Item) bool
    Append(e Item)
    Get(index int) Item
    Set(index int, e Item) Item
    Add(index int, e Item)
    Remove(index int) Item
    IndexOf(e Item) int
    Clear()
}

type ArrayList struct {
    size     int // 元素数量
    elements []Item
}

// NewArrayListWithCap 创建一个动态数组对象
// capacity 动态数组起始容量， 如果小于DEFAULT_CAPACITY(10)，则设置为DEFAULT_CAPACITY
func NewArrayListWithCap(capacity int) *ArrayList {
    if capacity < 0 {
        capacity = DEFAULT_CAPACITY
    }
    return &ArrayList{
        elements: make([]Item, capacity),
        size:     0,
    }
}

// NewArrayList 创建一个容量为DEFAULT_CAPACITY的动态数组
func NewArrayList() *ArrayList {
    return NewArrayListWithCap(DEFAULT_CAPACITY)
}

// Size 获取动态数组的元素数量
func (al *ArrayList) Size() int {
    return al.size
}

// IsEmpty 判断动态数组是否为空数组
func (al *ArrayList) IsEmpty() bool {
    return al.size == 0
}

// Contains 判断动态数组是否包含指定元素
func (al *ArrayList) Contains(e Item) bool {
    return al.IndexOf(e) != ELEMENT_NOT_FOUND
}

// Append 向动态数组最后位置插入元素
func (al *ArrayList) Append(e Item) {
    al.Add(al.size, e)
}

// Get 获取指定索引位置的元素
func (al *ArrayList) Get(index int) Item {
    al.rangeCheck(index)
    return al.elements[index]
}

// Set 设置指定索引位置的元素
func (al *ArrayList) Set(index int, e Item) Item {
    al.rangeCheck(index)
    old := al.elements[index]
    al.elements[index] = e
    return old
}

// Add 向指定索引位置插入元素
func (al *ArrayList) Add(index int, e Item) {
    al.rangeCheckForAdd(index)
    al.ensureCapacity(al.size + 1)
    for i := al.size; i > index; i-- {
        al.elements[i] = al.elements[i-1]
    }
    al.elements[index] = e
    al.size++
}

// Remove 移除指定索引位置的元素
func (al *ArrayList) Remove(index int) Item {
    al.rangeCheck(index)
    old := al.Get(index)
    for i := index; i < al.size-1; i++ {
        al.elements[i] = al.elements[i+1]
    }
    al.size--
    al.elements[al.size] = nil
    al.trim()
    return old
}

// IndexOf 返回元素在数组中的索引，若不存在，则返回-1
func (al *ArrayList) IndexOf(e Item) int {
    if e == nil {
        for i := 0; i < al.size; i++ {
            if al.elements[i] == nil {
                return i
            }
        }
    } else {
        for i := 0; i < al.size; i++ {
            if e.Equal(al.elements[i]) {
                return i
            }
        }
    }
    return ELEMENT_NOT_FOUND
}

// Clear 清空数组，并缩容
func (al *ArrayList) Clear() {
    for i := 0; i < al.size; i++ {
        al.elements[i] = nil
    }
    al.size = 0
    if al.elements != nil && len(al.elements) > DEFAULT_CAPACITY {
        al.elements = make([]Item, DEFAULT_CAPACITY)
    }
}

// 扩容数组，将数组容量扩容为原来1.5倍
func (al *ArrayList) ensureCapacity(capacity int) {
    oldCapacity := len(al.elements)
    if oldCapacity >= capacity {
        return
    }
    // 新容量为旧容量的1.5倍
    newCapacity := oldCapacity + (oldCapacity >> 1)
    newElements := make([]Item, newCapacity)
    for i := 0; i < al.size; i++ {
        newElements[i] = al.elements[i]
    }
    al.elements = newElements
}

// 缩容数组，将数组缩容为原来一半
func (al *ArrayList) trim() {
    oldCapacity := len(al.elements)
    if al.size >= (oldCapacity>>1) || oldCapacity <= DEFAULT_CAPACITY {
        return
    }
    // 剩余空间还很多
    newCapacity := oldCapacity >> 1
    newElements := make([]Item, newCapacity)
    for i := 0; i < al.size; i++ {
        newElements[i] = al.elements[i]
    }
    al.elements = newElements
}

// 索引越界panic
func (al *ArrayList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, al.size))
}

// 索引越界检查
func (al *ArrayList) rangeCheck(index int) {
    if index < 0 || index >= al.size {
        al.outOfBounds(index)
    }
}

// 针对新增元素的索引越界检查
func (al *ArrayList) rangeCheckForAdd(index int) {
    if index < 0 || index > al.size {
        al.outOfBounds(index)
    }
}

// 打印数组时的字符串表示
// [xx, xx, xx, xx]
func (al *ArrayList) String() string {
    alStr := "["
    for i := 0; i < al.size; i++ {
        if i != 0 {
            alStr = alStr + ", "
        }
        alStr = alStr + fmt.Sprintf("%#v", al.elements[i])
    }
    return alStr + "]"
}
