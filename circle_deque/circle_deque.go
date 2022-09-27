package circle_deque

const (
	DEFAULT_CAPACITY  = 10
	ELEMENT_NOT_FOUND = -1
)

type ICircleDeque interface {
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

func (d *CircleDeque) EnQueueFront(e E){
	d.add(0, e)
}

func (d *CircleDeque) EnQueueRear(e E) {
    d.add(d.size, e)
}

func (d *CircleDeque) DeQueueFront() E {
	return d.CircleDeque(0)
}

func (d *CircleDeque) DeQueueRear() E {
	return d.remove(d.size - 1)
}

func (d *CircleDeque) Front() E {
	return d.get(0)
}

func (d *CircleDeque) Rear() E {
	return d.get(d.size - 1)
}