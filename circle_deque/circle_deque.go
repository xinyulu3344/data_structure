package circle_deque

import "fmt"

const (
	DEFAULT_CAPACITY  = 10
	ELEMENT_NOT_FOUND = -1
)

type ICircleDeque interface {
	Size() int
	IsEmpty() bool
	EnQueueFront(e E) // 从对头入队
	EnQueueRear(e E)  // 从队尾入队
	DeQueueFront() E  // 从队头出队
	DeQueueRear() E   // 从对尾出队
	Front() E         // 获取对头元素
	Rear() E          // 获取队尾元素
}

type E interface {
	Equal(e E) bool
}

type CircleDeque struct {
	front    int // 存储对头元素的下标
	size     int
	elements []E
}

func NewCircleDeque() *CircleDeque {
	return &CircleDeque{
		size:     0,
		elements: make([]E, DEFAULT_CAPACITY),
	}
}

func (q *CircleDeque) Size() int {
	return q.size
}

func (q *CircleDeque) IsEmpty() bool {
	return q.size == 0
}

// 从头部入队
func (q *CircleDeque) EnQueueFront(e E) {
	q.ensureCapacity(q.size + 1)
	q.front = q.index(-1)
	q.elements[q.front] = e
	q.size++
}

// 从尾部入队
func (q *CircleDeque) EnQueueRear(e E) {
	q.ensureCapacity(q.size + 1)
	q.elements[q.index(q.size)] = e
	q.size++
}

// 从头部出队
func (q *CircleDeque) DeQueueFront() E {
	frontElement := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = q.index(1)
	q.size--
	return frontElement
}

// 从尾部出队
func (q *CircleDeque) DeQueueRear() E {
	rearIndex := q.index(q.size - 1)
	rear := q.elements[rearIndex]
	q.elements[rearIndex] = nil
	q.size--
	return rear
}

func (q *CircleDeque) Front() E {
	return q.elements[q.front]
}

func (q *CircleDeque) Rear() E {
	return q.elements[q.index(q.size-1)]
}

func (q *CircleDeque) Clear() {
	for i := 0; i < q.size; i++ {
		q.elements[q.index(i)] = nil
	}
	q.size = 0
	q.front = 0
}

func (q *CircleDeque) String() string {
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
func (q *CircleDeque) ensureCapacity(capacity int) {
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
func (q *CircleDeque) index(i int) int {
	i = q.front + i
	if i < 0 {
		i = i + len(q.elements)
	}
	return i % len(q.elements)
}
