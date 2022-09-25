package stack

import "testing"

type Int int

func (i Int) Equal(j E) bool {
    return i == j.(Int)
}

func testNewStack(t *testing.T, elements []Int) {
    s := NewStack()
    size := len(elements)
    for _, v := range elements {
        s.Push(Int(v))
    }
    if s.Size() != size {
        t.Errorf("s.Size() == %v != %v", s.Size(), size)
    }
    if (s.Size() == 0 && !s.IsEmpty()) || (s.Size() > 0 && s.IsEmpty()) {
        t.Errorf("s.Size() == %v but s.IsEmpty is %v\n", s.Size(), s.IsEmpty())
    }
    if s.Top() != elements[size-1] {
        t.Errorf("s.Top() == %v != %v", s.Top(), elements[size-1])
    }
    
    last := s.Pop()
    
    if last != elements[size-1] {
        t.Errorf("s.Pop() == %v != %v", last, elements[size-1])
    }
    
    if s.Size() != (size - 1) {
        t.Errorf("s.Size() == %v != %v", s.Size(), size-1)
    }
    if (s.Size() == 0 && !s.IsEmpty()) || (s.Size() > 0 && s.IsEmpty()) {
        t.Errorf("s.Size() == %v but s.IsEmpty is %v\n", s.Size(), s.IsEmpty())
    }
}

func TestNewStack(t *testing.T) {
    testNewStack(t, []Int{0})
    testNewStack(t, []Int{11, 22, 33, 44})
}
