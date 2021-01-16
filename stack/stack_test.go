package stack

import (
    "container/list"
    "fmt"
    "testing"
)

type testObj struct {
    Name string
}

func TestStack(t *testing.T) {

    stack := NewStack()
    for {
        stack.Push(testObj{Name: "Bob"})
        if stack.GetSize() == 100 {
            break
        }
    }
    for !stack.IsEmpty(){
        e := stack.Pop()
        fmt.Println(e.(*list.Element).Value)
    }
}

func TestClear(t *testing.T) {
    stack := NewStack()
    for {
        stack.Push(testObj{Name: "Bob"})
        if stack.GetSize() == 100 {
            break
        }
    }
    stack.Clear()
    fmt.Println(stack.Top())
}

