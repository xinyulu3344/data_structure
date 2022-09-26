package deque

import "testing"

type Int int

func (i Int) Equal(j E) bool {
	return i == j.(Int)
}

func testDeque(t *testing.T, elements []Int) {
	s := NewDeque()
	size := len(elements)
	// 从队尾入队，队头出队
	for _, v := range elements {
		s.EnQueueRear(Int(v))
	}
	if s.Size() != size {
		t.Errorf("s.Size() == %v != %v", s.Size(), size)
	}
	if (s.Size() == 0 && !s.IsEmpty()) || (s.Size() > 0 && s.IsEmpty()) {
		t.Errorf("s.Size() == %v but s.IsEmpty is %v\n", s.Size(), s.IsEmpty())
	}
	if s.Front() != elements[0] {
		t.Errorf("s.Front() == %v != %v", s.Front(), elements[0])
	}

	first := s.DeQueueFront()

	if first != elements[0] {
		t.Errorf("s.DeQueue() == %v != %v", first, elements[0])
	}

	if s.Size() != (size - 1) {
		t.Errorf("s.Size() == %v != %v", s.Size(), size-1)
	}
	if (s.Size() == 0 && !s.IsEmpty()) || (s.Size() > 0 && s.IsEmpty()) {
		t.Errorf("s.Size() == %v but s.IsEmpty is %v\n", s.Size(), s.IsEmpty())
	}
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("s.Size() == %v != %v", s.Size(), 0)
	}
	if !s.IsEmpty() {
		t.Errorf("s.IsEmpty() == %v != %v", s.IsEmpty(), !s.IsEmpty())
	}

	// 测试从队头入队，队尾出队
	for _, v := range elements {
		s.EnQueueFront(Int(v))
	}
	if s.Size() != size {
		t.Errorf("s.Size() == %v != %v", s.Size(), size)
	}
	if (s.Size() == 0 && !s.IsEmpty()) || (s.Size() > 0 && s.IsEmpty()) {
		t.Errorf("s.Size() == %v but s.IsEmpty is %v\n", s.Size(), s.IsEmpty())
	}
	if s.Rear() != elements[0] {
		t.Errorf("s.Front() == %v != %v", s.Front(), elements[0])
	}

	last := s.DeQueueRear()

	if last != elements[0] {
		t.Errorf("s.DeQueue() == %v != %v", last, elements[0])
	}

	if s.Size() != (size - 1) {
		t.Errorf("s.Size() == %v != %v", s.Size(), size-1)
	}
	if (s.Size() == 0 && !s.IsEmpty()) || (s.Size() > 0 && s.IsEmpty()) {
		t.Errorf("s.Size() == %v but s.IsEmpty is %v\n", s.Size(), s.IsEmpty())
	}
}

func TestNewDeque(t *testing.T) {
	s := make([]Int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, Int(i))
	}
	testDeque(t, s)
	testDeque(t, []Int{0})
	testDeque(t, []Int{11, 22, 33, 44})
}
