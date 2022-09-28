package circle_deque

import "testing"

type Int int

func (i Int) Equal(j E) bool {
	return i == j.(Int)
}


func TestCircleDeque(t *testing.T) {
	q := NewCircleDeque()
	for i := 0; i < 10; i++ {
		q.EnQueueFront(Int(i + 1))
		q.EnQueueRear(Int(i + 100))
	}
	t.Log(q) // [8, 7, 6, 5, 4, 3, 2, 1, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, <nil>, <nil>, 10, 9]  front=20 size=20 capacity=22

	for i := 0; i < 3; i++ {
		q.DeQueueFront()
		q.DeQueueRear()
	}
	q.EnQueueFront(Int(11))
	q.EnQueueFront(Int(12))
	t.Log(q) // [11, 7, 6, 5, 4, 3, 2, 1, 100, 101, 102, 103, 104, 105, 106, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, 12]  front=21 size=16 capacity=22
	for !q.IsEmpty() {
		t.Log(q.DeQueueFront())
	}
}
