package queue

import (
    "fmt"
)

type E interface {
	Equal(e E) bool
}

type IQueue interface {
	Size() int
	IsEmpty() bool
	EnQueue(e E)
	DeQueue() E
	Front() E
	Clear()
}


const (
    ELEMENT_NOT_FOUND = -1
)

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

type Queue struct {
    size int
    last *node
    root *node
}

func NewQueue() *Queue {
    return &Queue{}
}

func (q *Queue) Size() int {
    return q.size
}

func (q *Queue) IsEmpty() bool {
    return q.size == 0
}


func (q *Queue) EnQueue(e E) {
    q.add(q.size, e)
}

func (q *Queue) Front() E {
    return q.getNodeByIndex(0).element
}


func (q *Queue) add(index int, e E) {
    q.rangeCheckForAdd(index)
    
    if index == q.size { // 处理index == size
        q.last = newNode(q.last, e, nil)
        if q.last.prev == nil { // 处理添加链表第一个元素
            q.root = q.last
        } else {
            q.last.prev.next = q.last
        }
    } else {
        // 待插入节点后面的节点
        next := q.getNodeByIndex(index)
        // 待插入节点前面的节点
        prev := next.prev
        addNode := newNode(prev, e, next)
        next.prev = addNode
        
        // 处理index == 0的情况
        if prev == nil {
            q.root = addNode
        } else {
            prev.next = addNode
        }
    }
    q.size++
}

func (q *Queue) DeQueue() E {
    q.rangeCheck(0)
    
    n := q.getNodeByIndex(0)
    prev := n.prev
    next := n.next
    
    if prev == nil { // index == 0
        q.root = next
    } else {
        prev.next = next
    }
    
    if next == nil { // index == size - 1
        q.last = prev
    } else {
        next.prev = prev
    }
    
    q.size--
    return n.element
}


func (q *Queue) Clear() {
    q.size = 0
    q.root = nil
    q.last = nil
}

func (q *Queue) String() string {
    alStr := "["
    n := q.root
    for i := 0; i < q.size; i++ {
        if i != 0 {
            alStr = alStr + ", "
        }
        alStr = alStr + fmt.Sprintf("%#v", n.element)
        n = n.next
    }
    return alStr + "]"
}

// 获取index位置对应的节点对象
func (q *Queue) getNodeByIndex(index int) *node {
    q.rangeCheck(index)
    
    if index < (q.size >> 1) {
        n := q.root
        for i := 0; i < index; i++ {
            n = n.next
        }
        return n
    } else {
        n := q.last
        for i := q.size - 1; i > index; i-- {
            n = n.prev
        }
        return n
    }
}

func (q *Queue) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, q.size))
}

func (q *Queue) rangeCheck(index int) {
    if index < 0 || index >= q.size {
        q.outOfBounds(index)
    }
}

func (q *Queue) rangeCheckForAdd(index int) {
    if index < 0 || index > q.size {
        q.outOfBounds(index)
    }
}

