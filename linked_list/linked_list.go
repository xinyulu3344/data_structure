package linked_list

import (
    "container/list"
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

type LinkedList struct {
    size int
    root *node
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (l *LinkedList) Size() int {
    return l.size
}

func (l *LinkedList) IsEmpty() bool {
    return l.size == 0
}

func (l *LinkedList) Contains(e E) bool {
    return l.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (l *LinkedList) Append(e E) {
    l.Add(l.size, e)
}

func (l *LinkedList) Get(index int) E {
    return l.getNodeByIndex(index).element
}

func (l *LinkedList) Set(index int, e E) E {
    n := l.getNodeByIndex(index)
    old := n.element
    n.element = e
    return old
}

func (l *LinkedList) Add(index int, e E) {
    if index == 0 {
        l.root = &node{element: e, next: l.root}
    } else {
        prev := l.getNodeByIndex(index - 1)
        prev.next = &node{element: e, next: prev.next}
    }
    l.size++
}

func (l *LinkedList) Remove(index int) E {
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

func (l *LinkedList) IndexOf(e E) int {
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
            if n.element.Equal(n.element) {
                return i
            }
            n = n.next
        }
    }
    return ELEMENT_NOT_FOUND
}

func (l *LinkedList) Clear() {
    l.size = 0
    l.root = nil
}

func (l *LinkedList) String() string {
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
func (l *LinkedList) getNodeByIndex(index int) *node {
    l.rangeCheck(index)
    n := l.root
    for i := 0; i < index; i++ {
        n = n.next
    }
    return n
}

func (l *LinkedList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, l.size))
}

func (l *LinkedList) rangeCheck(index int) {
    if index < 0 || index >= l.size {
        l.outOfBounds(index)
    }
}

func (l *LinkedList) rangeCheckForAdd(index int) {
    if index < 0 || index > l.size {
        l.outOfBounds(index)
    }
}
