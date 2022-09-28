package circle_queue

import "fmt"

const (
	DEFAULT_CAPACITY  = 10
	ELEMENT_NOT_FOUND = -1
)

type ICircleQueue interface {
	Size() int
	IsEmpty() bool
	EnQueue(e E)
	DeQueue() E
	Front() E
}

type E interface {
	Equal(e E) bool
}

type CircleQueue struct {
	front    int // 存储对头元素的下标
	size     int
	elements []E
}

func NewCircleQueue() *CircleQueue {
	return &CircleQueue{
		size:     0,
		elements: make([]E, DEFAULT_CAPACITY),
	}
}

func (q *CircleQueue) Size() int {
	return q.size
}

func (q *CircleQueue) IsEmpty() bool {
	return q.size == 0
}

// 入队
func (q *CircleQueue) EnQueue(e E) {
	q.ensureCapacity(q.size + 1)
	q.elements[q.index(q.size)] = e
	q.size++
}

// 出队
func (q *CircleQueue) DeQueue() E {
	frontElement := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = q.index(1)
	q.size--
	return frontElement
}

func (q *CircleQueue) Front() E {
	return q.elements[q.front]
}

func (q *CircleQueue) Clear() {
	for i := 0; i < q.size; i++ {
		q.elements[q.index(i)] = nil
	}
	q.size = 0
	q.front = 0
}

func (q *CircleQueue) String() string {
	alStr := "["
	for i := 0; i < len(q.elements); i++ {
		if i != 0 {
			alStr = alStr + ", "
		}
		alStr = alStr + fmt.Sprintf("%#v", q.elements[i])
	}
	return alStr + fmt.Sprintf("]  front=%v size=%v capacity=%v", q.front, q.size, len(q.elements))
}

// 队列扩容
func (q *CircleQueue) ensureCapacity(capacity int) {
	oldCapacity := len(q.elements)
    if oldCapacity >= capacity {
        return
    }
    // 新容量为旧容量的1.5倍
    newCapacity := oldCapacity + (oldCapacity >> 1)
    newElements := make([]E, newCapacity)
    for i := 0; i < q.size; i++ {
        newElements[i] = q.elements[q.index(i)]
    }
    q.elements = newElements
	// 重置front
	q.front = 0
}

// 获取循环队列中的真实索引
func (q *CircleQueue) index(i int) int {
	return (i + q.front) % len(q.elements)
}