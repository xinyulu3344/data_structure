package priorityqueue

type IPriorityQueue interface {
	Size() int
	IsEmpty() bool
	EnQueue(e any) error
	Dequeue() any
	Front() any
	Clear()
}