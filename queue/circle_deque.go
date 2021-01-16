/**
 * @Author: xinyulu
 * @Date: 2020-11-30 16:12:36
 * @Description: 循环双端队列
**/
package queue

import "errors"

type CircleDeque struct {
    front int
    size int
    elements []interface{}
}

func NewCircleDeque(cap int) *CircleDeque {
    return &CircleDeque{
        front: 0,
        size: 0,
        elements: make([]interface{}, cap),
    }
}

// 获取队列元素个数
func (c *CircleDeque) GetSize() int {
    return c.size
}

// 从尾部入队
func (c *CircleDeque) EnQueueRear(e interface{}) {
    if c.isFull() {
        c.resize(c.size + 1)
    }
    c.elements[c.index(c.size)] = e
    c.size++
}

// 从头部入队
func (c *CircleDeque) EnQueueFront(e interface{}) {
    if c.isFull() {
        c.resize(c.size + 1)
    }
    c.front = c.index(-1)
    c.elements[c.front] = e
    c.size++
}

// 从尾部出队
func (c *CircleDeque) DeQueueRear() (interface{}, error) {
    if c.IsEmpty() {
        return nil, errors.New("队列是空队列")
    }
    rearIndex := c.index(c.size - 1)
    rear := c.elements[rearIndex]
    c.elements[rearIndex] =  nil
    c.size--
    return rear, nil
}


// 从头部出队
func (c *CircleDeque) DeQueueFront() (interface{}, error) {
    if c.IsEmpty() {
        return nil, errors.New("队列是空队列")
    }
    e := c.elements[c.front]
    c.elements[c.front] = nil
    c.front = c.index(1)
    c.size--
    return e, nil
}

// 获取队尾元素
func (c *CircleDeque) Rear() (e interface{}) {
    return c.elements[c.index(c.size - 1)]
}

// 获取队头元素
func (c *CircleDeque) Front() interface{} {
    return c.elements[c.front]
}

// 判断队列是否为空
func (c *CircleDeque) IsEmpty() bool {
    return c.size == 0
}

// 判断队列是否已满
func (c *CircleDeque) isFull() bool {
    return c.size == len(c.elements)
}

// 队列扩容
func (c *CircleDeque) resize(cap int) {
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

func (c *CircleDeque) index(i int) int {
    i += c.front
    if i < 0 {
        return i + len(c.elements)
    }
    return i % len(c.elements)
}

// 清空队列
func (c *CircleDeque) Clear() {
    for i := 0; i < c.size; i++ {
        c.elements[c.index(i)] = nil
    }
    c.size = 0
    c.front = 0
}