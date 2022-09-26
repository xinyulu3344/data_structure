package queue

import "testing"


type Int int

func (i Int) Equal(j E) bool {
    return i == j.(Int)
}

func testNewQueue(t *testing.T, elements []Int) {
    s := NewQueue()
    size := len(elements)
    for _, v := range elements {
        s.EnQueue(Int(v))
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
    
    first := s.DeQueue()
    
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
}

func TestNewQueue(t *testing.T) {
	s := make([]Int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, Int(i))
	}
	testNewQueue(t, s)
    testNewQueue(t, []Int{0})
    testNewQueue(t, []Int{11, 22, 33, 44})
}
