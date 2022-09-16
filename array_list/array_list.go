package array_list

import "fmt"

const (
    DEFAULT_CAPACITY = 10
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

type ArrayList struct{
    size int // 元素数量
    elements []Item
}

func NewArrayListWithCap(capacity int) *ArrayList {
    if capacity < 0 {
        capacity = DEFAULT_CAPACITY
    }
    return &ArrayList{
        elements: make([]Item, 0, capacity),
        size: 0,
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
    al.elements = append(al.elements, e)
    al.size++
}

func (al *ArrayList) Get(index int) Item {
    if index < 0 || index >= al.size {
        panic(fmt.Sprintf("Index: %d Size: %d", index, al.size))
    }
    return al.elements[index]
}

func (al *ArrayList) Set(index int, e Item) Item {
    if index < 0 || index >= al.size {
        panic(fmt.Sprintf("Index: %d Size: %d", index, al.size))
    }
    old := al.elements[index]
    al.elements[index] = e
    return old
}

func (al *ArrayList) Add(index int, e Item) {
}

func (al *ArrayList) Remove(index int) Item {
    return nil
}

// IndexOf
// 返回元素在数组中的索引，若不存在，则返回-1
func (al *ArrayList) IndexOf(e Item) int {
    for i, v := range al.elements {
        if v.Equal(e) {
            return i
        }
    }
    return ELEMENT_NOT_FOUND
}

func (al *ArrayList) Clear() {
    al.size = 0
    al.elements = al.elements[0:0]
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
