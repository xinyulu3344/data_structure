package deque

import (
    "fmt"
)

const (
    ELEMENT_NOT_FOUND = -1
)

type IDeque interface {
	Size() int
	IsEmpty() bool
	EnQueueFront(e E) // 从对头入队
	EnQueueRear(e E) // 从队尾入队
	DeQueueFront() E // 从队头出队
	DeQueueRear() E // 从对尾出队
	Front() E // 获取对头元素
	Rear() E // 获取队尾元素
}

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

type Deque struct {
    size int
    last *node
    root *node
}

func NewDeque() *Deque {
    return &Deque{}
}

func (d *Deque) Size() int {
    return d.size
}

func (d *Deque) IsEmpty() bool {
    return d.size == 0
}

func (d *Deque) EnQueueFront(e E){
	d.add(0, e)
}

func (d *Deque) EnQueueRear(e E) {
    d.add(d.size, e)
}

func (d *Deque) DeQueueFront() E {
	return d.remove(0)
}

func (d *Deque) DeQueueRear() E {
	return d.remove(d.size - 1)
}

func (d *Deque) Front() E {
	return d.get(0)
}

func (d *Deque) Rear() E {
	return d.get(d.size - 1)
}

func (d *Deque) get(index int) E {
    return d.getNodeByIndex(index).element
}

func (d *Deque) add(index int, e E) {
    d.rangeCheckForAdd(index)
    
    if index == d.size { // 处理index == size
        d.last = newNode(d.last, e, nil)
        if d.last.prev == nil { // 处理添加链表第一个元素
            d.root = d.last
        } else {
            d.last.prev.next = d.last
        }
    } else {
        // 待插入节点后面的节点
        next := d.getNodeByIndex(index)
        // 待插入节点前面的节点
        prev := next.prev
        addNode := newNode(prev, e, next)
        next.prev = addNode
        
        // 处理index == 0的情况
        if prev == nil {
            d.root = addNode
        } else {
            prev.next = addNode
        }
    }
    d.size++
}

func (d *Deque) remove(index int) E {
    d.rangeCheck(index)
    
    n := d.getNodeByIndex(index)
    prev := n.prev
    next := n.next
    
    if prev == nil { // index == 0
        d.root = next
    } else {
        prev.next = next
    }
    
    if next == nil { // index == size - 1
        d.last = prev
    } else {
        next.prev = prev
    }
    
    d.size--
    return n.element
}

func (d *Deque) Clear() {
    d.size = 0
    d.root = nil
    d.last = nil
}

func (d *Deque) String() string {
    alStr := "["
    n := d.root
    for i := 0; i < d.size; i++ {
        if i != 0 {
            alStr = alStr + ", "
        }
        alStr = alStr + fmt.Sprintf("%#v", n.element)
        n = n.next
    }
    return alStr + "]"
}

// 获取index位置对应的节点对象
func (d *Deque) getNodeByIndex(index int) *node {
    d.rangeCheck(index)
    
    if index < (d.size >> 1) {
        n := d.root
        for i := 0; i < index; i++ {
            n = n.next
        }
        return n
    } else {
        n := d.last
        for i := d.size - 1; i > index; i-- {
            n = n.prev
        }
        return n
    }
}

func (d *Deque) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, d.size))
}

func (d *Deque) rangeCheck(index int) {
    if index < 0 || index >= d.size {
        d.outOfBounds(index)
    }
}

func (d *Deque) rangeCheckForAdd(index int) {
    if index < 0 || index > d.size {
        d.outOfBounds(index)
    }
}
