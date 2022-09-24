package doubly_linked_list

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
    prev    *node
}

func newNode(prev *node, element E, next *node) *node {
    return &node{
        element: element,
        next:    next,
        prev:    prev,
    }
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

type DoublyLinkedList struct {
    size int
    last *node
    root *node
}

func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Size() int {
    return l.size
}

func (l *DoublyLinkedList) IsEmpty() bool {
    return l.size == 0
}

func (l *DoublyLinkedList) Contains(e E) bool {
    return l.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (l *DoublyLinkedList) Append(e E) {
    l.Add(l.size, e)
}

func (l *DoublyLinkedList) Get(index int) E {
    return l.getNodeByIndex(index).element
}

func (l *DoublyLinkedList) Set(index int, e E) E {
    n := l.getNodeByIndex(index)
    old := n.element
    n.element = e
    return old
}

func (l *DoublyLinkedList) Add(index int, e E) {
    l.rangeCheckForAdd(index)
    
    if index == l.size { // 处理index == size
        l.last = newNode(l.last, e, nil)
        if l.last.prev == nil { // 处理添加链表第一个元素
            l.root = l.last
        } else {
            l.last.prev.next = l.last
        }
    } else {
        // 待插入节点后面的节点
        next := l.getNodeByIndex(index)
        // 待插入节点前面的节点
        prev := next.prev
        addNode := newNode(prev, e, next)
        next.prev = addNode
        
        // 处理index == 0的情况
        if prev == nil {
            l.root = addNode
        } else {
            prev.next = addNode
        }
    }
    l.size++
}

func (l *DoublyLinkedList) Remove(index int) E {
    l.rangeCheck(index)
    
    n := l.getNodeByIndex(index)
    prev := n.prev
    next := n.next
    
    if prev == nil { // index == 0
        l.root = next
    } else {
        prev.next = next
    }
    
    if next == nil { // index == size - 1
        l.last = prev
    } else {
        next.prev = prev
    }
    
    l.size--
    return n.element
}

func (l *DoublyLinkedList) IndexOf(e E) int {
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

func (l *DoublyLinkedList) Clear() {
    l.size = 0
    l.root = nil
    l.last = nil
}

func (l *DoublyLinkedList) String() string {
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
func (l *DoublyLinkedList) getNodeByIndex(index int) *node {
    l.rangeCheck(index)
    
    if index < (l.size >> 1) {
        n := l.root
        for i := 0; i < index; i++ {
            n = n.next
        }
        return n
    } else {
        n := l.last
        for i := l.size - 1; i > index; i-- {
            n = n.prev
        }
        return n
    }
}

func (l *DoublyLinkedList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, l.size))
}

func (l *DoublyLinkedList) rangeCheck(index int) {
    if index < 0 || index >= l.size {
        l.outOfBounds(index)
    }
}

func (l *DoublyLinkedList) rangeCheckForAdd(index int) {
    if index < 0 || index > l.size {
        l.outOfBounds(index)
    }
}
