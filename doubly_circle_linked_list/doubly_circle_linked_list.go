package doubly_circle_linked_list

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

type DoublyCircleLinkedList struct {
    size int
    last *node
    root *node
}

func NewDoublyCircleLinkedList() *DoublyCircleLinkedList {
    return &DoublyCircleLinkedList{}
}

func (l *DoublyCircleLinkedList) Size() int {
    return l.size
}

func (l *DoublyCircleLinkedList) IsEmpty() bool {
    return l.size == 0
}

func (l *DoublyCircleLinkedList) Contains(e E) bool {
    return l.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (l *DoublyCircleLinkedList) Append(e E) {
    l.Add(l.size, e)
}

func (l *DoublyCircleLinkedList) Get(index int) E {
    return l.getNodeByIndex(index).element
}

func (l *DoublyCircleLinkedList) Set(index int, e E) E {
    n := l.getNodeByIndex(index)
    old := n.element
    n.element = e
    return old
}

func (l *DoublyCircleLinkedList) Add(index int, e E) {
    l.rangeCheckForAdd(index)
    
    if index == l.size { // 处理在最后位置添加节点
        l.last = newNode(l.last, e, l.root) // 让last指向新增节点，并让新增节点的next，指向头结点
        if l.last.prev == nil {             // 如果链表为空
            l.root = l.last      // 让root指向新添加的节点
            l.last.prev = l.root // 让新添加节点的prev指向自己
            l.last.next = l.root // 让新添加节点的next指向自己
        } else { // 如果链表不为空
            l.last.prev.next = l.last // 让新增节点的上一个节点的next，指向新增节点
            l.root.prev = l.last      // 让头结点的prev，指向新增节点
        }
    } else {
        // 获取待插入节点后面的节点
        next := l.getNodeByIndex(index)
        // 获取待插入节点前面的节点
        prev := next.prev
        
        addNode := newNode(prev, e, next)
        
        // 修改前后节点prev、next的指向
        prev.next = addNode
        next.prev = addNode
        
        if next == l.root { // 处理index == 0 的情况
            l.root = addNode
        }
    }
    l.size++
}

func (l *DoublyCircleLinkedList) Remove(index int) E {
    l.rangeCheck(index)
    
    n := l.root
    
    if l.size == 1 {
        l.root = nil
        l.last = nil
    } else {
        n = l.getNodeByIndex(index)
        prev := n.prev
        next := n.next
        prev.next = next
        next.prev = prev
        
        if n == l.root { // index == 0
            l.root = next
        }
        
        if n == l.last { // index == size - 1
            l.last = prev
        }
    }
    
    l.size--
    return n.element
}

func (l *DoublyCircleLinkedList) IndexOf(e E) int {
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

func (l *DoublyCircleLinkedList) Clear() {
    l.size = 0
    l.root = nil
    l.last = nil
}

func (l *DoublyCircleLinkedList) String() string {
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
func (l *DoublyCircleLinkedList) getNodeByIndex(index int) *node {
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

func (l *DoublyCircleLinkedList) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, l.size))
}

func (l *DoublyCircleLinkedList) rangeCheck(index int) {
    if index < 0 || index >= l.size {
        l.outOfBounds(index)
    }
}

func (l *DoublyCircleLinkedList) rangeCheckForAdd(index int) {
    if index < 0 || index > l.size {
        l.outOfBounds(index)
    }
}
