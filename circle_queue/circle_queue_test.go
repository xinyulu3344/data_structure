package circle_queue

import "testing"

type Int int

func (i Int) Equal(j E) bool {
	return i == j.(Int)
}

func testCircleQueue(t *testing.T, elements []Int) {
	q := NewCircleQueue()
	size := len(elements)
	for _, v := range elements {
		q.EnQueue(v)
	}

	if q.Size() != size {
		t.Errorf("q.Size() == %v != %v", q.Size(), size)
	}
	if (q.Size() == 0 && !q.IsEmpty()) || (q.Size() > 0 && q.IsEmpty()) {
		t.Errorf("q.Size() == %v but q.IsEmpty is %v\n", q.Size(), q.IsEmpty())
	}
	if q.Front() != elements[0] {
		t.Errorf("q.Front() == %v != %v", q.Front(), elements[0])
	}
	for i := 0; i < size; i++ {
		q.DeQueue()
	}
	t.Log(q)
}

func TestCircleQueue(t *testing.T) {
	testCircleQueue(t, []Int{0})
}
