package linked_list

import (
	"container/list"
	"fmt"
)

const (
    DEFAULT_CAPACITY  = 10
    ELEMENT_NOT_FOUND = -1
)

type node struct {
	element any
	next *node
}

type ILinkedList interface {
    Size() int
    IsEmpty() bool
    Contains(e any) bool
    Append(e any)
    Get(index int) any
    Set(index int, e any) any
    Add(index int, e any)
    Remove(index int) any
    IndexOf(e any) int
    Clear()
}

type LinkedList struct {
	size int
	root *node
}

func NewLinkedList() *LinkedList {
	list.New()
	return &LinkedList{}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Contains(e any) bool {
	return l.IndexOf(e) != ELEMENT_NOT_FOUND
}

func (l *LinkedList) Append(e any) {
	l.Add(l.size, e)
}

func (l *LinkedList) Get(index int) any {
	return l.getNodeByIndex(index).element
}

func (l *LinkedList) Set(index int, e any) any {
	n := l.getNodeByIndex(index)
	old := n.element
	n.element = e
	return old
}

func (l *LinkedList) Add(index int, e any) {
	if index == 0 {
		l.root = &node{element: e, next: l.root}
	} else {
		prev := l.getNodeByIndex(index - 1)
		prev.next = &node{element: e, next: prev.next}
	}
	l.size++
}

func (l *LinkedList) Remove(index int) any {
	var old any
	if index == 0 {
		old = l.root.element
		l.root = l.root.next
	} else {
		prev := l.getNodeByIndex(index - 1)
		old = prev.next.element
		prev.next = prev.next.next
	}
	l.size--
	return old
}

func (l *LinkedList) IndexOf(e any) int {}

func (l *LinkedList) Clear() {
	l.size = 0
	l.root = nil
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