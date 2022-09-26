package circle_queue


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

func (q *CircleQueue) EnQueue(e E) {
	q.elements[(q.front+q.size) % len(q.elements)] = e
	q.size++
}

func (q *CircleQueue) DeQueue() E {
	frontElement := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = (q.front + 1) % len(q.elements)
	q.size--
	return frontElement
}

func (q *CircleQueue) Front() E {
	return q.elements[q.front]
}

func (q *CircleQueue) String() string {
	
}
