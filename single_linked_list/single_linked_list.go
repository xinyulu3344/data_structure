package single_linked_list

import (
    "fmt"
)

const (
    DEFAULT_CAPACITY  = 10
    ELEMENT_NOT_FOUND = -1
)

type E interface {
    Equal(e E) bool
}

type node struct {
    element E
    next    *node
}

type ILinkedList interface {
    Size() int
    IsEmpty() bool
    Contains(e E) bool
    Append(e E)
    Get(index int) E
    Set(index int, e E) E
    Add(index int, e E)
    Remove(index int) E
    IndexOf(e E) int
    Clear()
}

type SingleLinkedList struct {
    size int
    root *node
}

func NewSingleLinkedList() *SingleLinkedList {
    return &SingleLinkedList{}
}

func (l *SingleLinkedList) Size() int {
    return l.size
}

func (l *SingleLinkedList) IsEmpty() bool {
    return l.size == 0
}

func (l *SingleLinkedList) Contains(e E) bool {
    return l.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (l *SingleLinkedList) Append(e E) {
    l.Add(l.size, e)
}

func (l *SingleLinkedList) Get(index int) E {
    return l.getNodeByIndex(index).element
}

func (l *SingleLinkedList) Set(index int, e E) E {
    n := l.getNodeByIndex(index)
    old := n.element
    n.element = e
    return old
}

func (l *SingleLinkedList) Add(index int, e E) {
    l.rangeCheckForAdd(index)
    if index == 0 {
        l.root = &node{element: e, next: l.root}
    } else {
        prev := l.getNodeByIndex(index - 1)
        prev.next = &node{element: e, next: prev.next}
    }
    l.size++
}

func (l *SingleLinkedList) Remove(index int) E {
    l.rangeCheck(index)
    n := l.root
    if index == 0 {
        l.root = l.root.next
    } else {
        prev := l.getNodeByIndex(index - 1)
        n = prev.next
        prev.next = prev.next.next
    }
    l.size--
    return n.element
}

func (l *SingleLinkedList) IndexOf(e E) int {
    n := l.root
    if e == nil {
        for i := 0; i < l.size; i++ {
            if n.element == nil {
                return i
            }
            n = n.next
        }
    } else {
        for i := 0; i < l.size; i++ {
            if e.Equal(n.element) {
                return i
            }
            n = n.next
        }
    }
    return ELEMENT_NOT_FOUND
}

func (l *SingleLinkedList) Clear() {
    l.size = 0
    l.root = nil
}

func (l *SingleLinkedList) String() string {
    alStr := "["
    n := l.root
    for i := 0; i < l.size; i++ {
        if i != 0 {
            alStr = alStr + ", "
        }
        alStr = alStr + fmt.Sprintf("%#v", n.element)
        n = n.next
    }
    return alStr + "]"
}

// 获取index位置对应的节点对象
func (l *SingleLinkedList) getNodeByIndex(index int) *node {
    l.rangeCheck(index)
    n := l.root
    for i := 0; i < index; i++ {
        n = n.next
    }
    return n
}

func (l *SingleLinkedList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, l.size))
}

func (l *SingleLinkedList) rangeCheck(index int) {
    if index < 0 || index >= l.size {
        l.outOfBounds(index)
    }
}

func (l *SingleLinkedList) rangeCheckForAdd(index int) {
    if index < 0 || index > l.size {
        l.outOfBounds(index)
    }
}
