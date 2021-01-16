package queue

import (
    "fmt"
    "testing"
)

func TestDeque(t *testing.T) {
    deque := NewDeque()
    deque.EnQueueFront(11)
    deque.EnQueueFront(22)
    deque.EnQueueRear(33)
    deque.EnQueueRear(44)
    for !deque.IsEmpty() {
        fmt.Println(deque.DeQueueFront())
    }
    deque.EnQueueFront(11)
    deque.EnQueueFront(22)
    deque.EnQueueRear(33)
    deque.EnQueueRear(44)
    for !deque.IsEmpty() {
        fmt.Println(deque.DeQueueRear())
    }
}
