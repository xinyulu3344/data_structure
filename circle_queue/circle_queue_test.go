package circle_queue

import "testing"

type Int int

func (i Int) Equal(j E) bool {
	return i == j.(Int)
}

func TestCircleQueue(t *testing.T) {
	q := NewCircleQueue()
	for i := 0; i < 10; i++ {
		q.EnQueue(Int(i))
	}
	t.Log(q.elements)
	q.DeQueue()
	t.Log(q.elements)
	q.EnQueue(Int(88))
	t.Log(q.elements)
	for i := 0; i < 10; i++ {
		q.DeQueue()
	}
	t.Log(q.elements)
}