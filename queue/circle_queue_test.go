package queue

import (
    "fmt"
    "testing"
)

func TestCircleQueue(t *testing.T) {
    queue := NewCircleQueue(10)
    for i := 0; i < 5; i++ {
        queue.EnQueue(i)
    }
    fmt.Println(queue)
    for i := 0; i < 4; i++ {
       _, err := queue.DeQueue()
       if err != nil {
           fmt.Println(err)
           return
       }
    }
    fmt.Println(queue)
    for i := 5; i < 20; i++ {
        queue.EnQueue(i)
    }
    fmt.Println(queue)
}
