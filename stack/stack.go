package stack

import "container/list"

type Stack struct {
    lst *list.List
}

func NewStack() *Stack {
    s := &Stack{}
    s.lst = list.New()
    return s
}

func (s *Stack) GetSize() int {
    return s.lst.Len()
}

func (s *Stack) Push(e interface{}) {
    s.lst.PushBack(e)
}

func (s *Stack) Pop() interface{} {
    return s.lst.Remove(s.lst.Back())
}

func (s *Stack) Top() interface{} {
    return s.lst.Back().Value
}

func (s *Stack) IsEmpty() bool {
    return s.lst.Len() == 0
}

func (s *Stack) Clear() {
    s.lst.Init()
}