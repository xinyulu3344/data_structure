package single_circle_linked_list

import (
    "fmt"
)

const (
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

type SingleCircleLinkedList struct {
    size int
    root *node
}

func NewSingleCircleLinkedList() *SingleCircleLinkedList {
    return &SingleCircleLinkedList{}
}

func (l *SingleCircleLinkedList) Size() int {
    return l.size
}

func (l *SingleCircleLinkedList) IsEmpty() bool {
    return l.size == 0
}

func (l *SingleCircleLinkedList) Contains(e E) bool {
    return l.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (l *SingleCircleLinkedList) Append(e E) {
    l.Add(l.size, e)
}

func (l *SingleCircleLinkedList) Get(index int) E {
    return l.getNodeByIndex(index).element
}

func (l *SingleCircleLinkedList) Set(index int, e E) E {
    n := l.getNodeByIndex(index)
    old := n.element
    n.element = e
    return old
}

func (l *SingleCircleLinkedList) Add(index int, e E) {
    l.rangeCheckForAdd(index)
    if index == 0 {
        l.root = &node{element: e, next: l.root}
        var last *node
        if l.size == 0 { // 如果是空链表, 让last指向root
            last = l.root
        } else {
            // 获取最后一个节点
            last = l.getNodeByIndex(l.size - 1).next
        }
        last.next = l.root
    } else {
        prev := l.getNodeByIndex(index - 1)
        prev.next = &node{element: e, next: prev.next}
    }
    l.size++
}

func (l *SingleCircleLinkedList) Remove(index int) E {
    l.rangeCheck(index)
    n := l.root
    if index == 0 {
        if l.size == 1 {
            l.root = nil
        } else {
            last := l.getNodeByIndex(l.size - 1)
            l.root = l.root.next
            last.next = l.root
        }
    } else {
        prev := l.getNodeByIndex(index - 1)
        n = prev.next
        prev.next = prev.next.next
    }
    l.size--
    return n.element
}

func (l *SingleCircleLinkedList) IndexOf(e E) int {
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

func (l *SingleCircleLinkedList) Clear() {
    l.size = 0
    l.root = nil
}

func (l *SingleCircleLinkedList) String() string {
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
func (l *SingleCircleLinkedList) getNodeByIndex(index int) *node {
    l.rangeCheck(index)
    n := l.root
    for i := 0; i < index; i++ {
        n = n.next
    }
    return n
}

func (l *SingleCircleLinkedList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, l.size))
}

func (l *SingleCircleLinkedList) rangeCheck(index int) {
    if index < 0 || index >= l.size {
        l.outOfBounds(index)
    }
}

func (l *SingleCircleLinkedList) rangeCheckForAdd(index int) {
    if index < 0 || index > l.size {
        l.outOfBounds(index)
    }
}
