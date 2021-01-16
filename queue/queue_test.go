package queue

import (
    "fmt"
    "testing"
)

func TestQueue(t *testing.T) {
    queue := NewQueue()
    queue.EnQueue(11)
    queue.EnQueue(22)
    queue.EnQueue(33)
    queue.EnQueue(44)
    for !queue.IsEmpty() {
        fmt.Println(queue.DeQueue())
    }
}
