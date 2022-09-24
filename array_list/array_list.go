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

func NewArrayListWithCap(capacity int) *ArrayList {
    if capacity < 0 {
        capacity = DEFAULT_CAPACITY
    }
    return &ArrayList{
        elements: make([]Item, capacity),
        size:     0,
    }
}

func NewArrayList() *ArrayList {
    return NewArrayListWithCap(DEFAULT_CAPACITY)
}

func (al *ArrayList) Size() int {
    return al.size
}

func (al *ArrayList) IsEmpty() bool {
    return al.size == 0
}

// Contains
// 判断数组是否包含元素
func (al *ArrayList) Contains(e Item) bool {
    return al.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (al *ArrayList) Append(e Item) {
    al.Add(al.size, e)
}

func (al *ArrayList) Get(index int) Item {
    al.rangeCheck(index)
    return al.elements[index]
}

func (al *ArrayList) Set(index int, e Item) Item {
    al.rangeCheck(index)
    old := al.elements[index]
    al.elements[index] = e
    return old
}

func (al *ArrayList) Add(index int, e Item) {
    al.rangeCheckForAdd(index)
    al.ensureCapacity(al.size + 1)
    for i := al.size; i > index; i-- {
        al.elements[i] = al.elements[i-1]
    }
    al.elements[index] = e
    al.size++
}

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

// IndexOf
// 返回元素在数组中的索引，若不存在，则返回-1
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

func (al *ArrayList) Clear() {
    for i := 0; i < al.size; i++ {
        al.elements[i] = nil
    }
    al.size = 0
    if al.elements != nil && len(al.elements) > DEFAULT_CAPACITY {
        al.elements = make([]Item, DEFAULT_CAPACITY)
    }
}

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

func (al *ArrayList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, al.size))
}

func (al *ArrayList) rangeCheck(index int) {
    if index < 0 || index >= al.size {
        al.outOfBounds(index)
    }
}

func (al *ArrayList) rangeCheckForAdd(index int) {
    if index < 0 || index > al.size {
        al.outOfBounds(index)
    }
}

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
