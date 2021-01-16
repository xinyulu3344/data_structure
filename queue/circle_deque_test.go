package queue

import (
    "fmt"
    "testing"
)

func TestCircleDeque1(t *testing.T) {
    queue := NewCircleDeque(10)
    for i := 0; i < 5; i++ {
        queue.EnQueueRear(i)
    }
    fmt.Println(queue)
    for i := 0; i < 4; i++ {
        _, err := queue.DeQueueFront()
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    fmt.Println(queue)
    for i := 5; i < 20; i++ {
        queue.EnQueueRear(i)
    }
    fmt.Println(queue)
}


func TestCircleDeque2(t *testing.T) {
    queue := NewCircleDeque(10)
    for i := 0; i < 5; i++ {
        queue.EnQueueRear(i)
    }
    fmt.Println(queue)
    for i := 0; i < 4; i++ {
        _, err := queue.DeQueueRear()
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    fmt.Println(queue)
    for i := 5; i < 20; i++ {
        queue.EnQueueFront(i)
    }
    fmt.Println(queue)

    for i := 0; i < 10; i++ {
        _, err := queue.DeQueueFront()
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    fmt.Println(queue)
}

func TestCircleDequeClear(t *testing.T) {
    queue := NewCircleDeque(10)
    for i := 0; i < 5; i++ {
        queue.EnQueueRear(i)
    }
    fmt.Println(queue)
    queue.Clear()
    fmt.Println(queue)
}