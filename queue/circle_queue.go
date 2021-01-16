/**
 * @Author: xinyulu
 * @Date: 2020-11-30 16:12:15
 * @Description: 循环队列
**/
package queue

import "errors"

type CircleQueue struct {
    front int
    size int
    elements []interface{}
}

func NewCircleQueue(cap int) *CircleQueue {
    return &CircleQueue{
        front: 0,
        size: 0,
        elements: make([]interface{}, cap),
    }
}

func (c *CircleQueue) GetSize() int {
    return c.size
}

func (c *CircleQueue) IsEmpty() bool {
    return c.size == 0
}

func (c *CircleQueue) EnQueue(e interface{}) {
    if c.IsFull() {
        c.resize(c.size+1)
    }
    c.elements[c.index(c.size)] = e
    c.size++
}

func (c *CircleQueue) DeQueue() (interface{}, error) {
    if c.IsEmpty() {
        return nil, errors.New("队列是空队列")
    }
    e := c.elements[c.front]
    c.elements[c.front] = nil
    c.front = c.index(1)
    c.size--
    return e, nil
}

func (c *CircleQueue) Front() (interface{}, error) {
    if c.IsEmpty() {
        return nil, errors.New("队列是空队列")
    }
    return c.elements[c.front], nil
}

func (c *CircleQueue) IsFull() bool {
    return c.size == len(c.elements)
}

func (c *CircleQueue) resize(cap int) {
    oldCapacity := len(c.elements)
    if oldCapacity >= cap {
        return
    }
    newCapacity := oldCapacity + (oldCapacity >> 1)
    newElements := make([]interface{}, newCapacity)
    for i := 0; i < c.size; i++ {
        newElements[i] = c.elements[c.index(i)]
    }
    c.elements = newElements
    c.front = 0
}

func (c *CircleQueue) index(i int) int {
    return (c.front + i) % len(c.elements)
}

func (c *CircleQueue) Clear() {
    for i := 0; i < c.size; i++ {
        c.elements[c.index(i)] = nil
    }
    c.size = 0
    c.front = 0
}