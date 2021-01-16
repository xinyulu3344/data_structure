/**
 * @Author: xinyulu
 * @Date: 2020-11-30 16:11:25
 * @Description: 双端队列
**/
package queue

import "container/list"


type Deque struct {
    lst *list.List
}

func NewDeque() *Deque {
    queue := &Deque{}
    queue.lst = list.New()
    return queue
}

// 获取队列元素个数
func (d *Deque) GetSize() int {
    return d.lst.Len()
}

// 从尾部入队
func (d *Deque) EnQueueRear(e interface{}) {
    d.lst.PushBack(e)
}

// 从头部入队
func (d *Deque) EnQueueFront(e interface{}) {
    d.lst.PushFront(e)
}

// 从尾部出队
func (d *Deque) DeQueueRear() interface{} {
    return d.lst.Remove(d.lst.Back())
}


// 从头部出队
func (d *Deque) DeQueueFront() interface{} {
    return d.lst.Remove(d.lst.Front())
}

// 获取队尾元素
func (d *Deque) Rear() (e interface{}) {
    return d.lst.Back().Value
}

// 获取队头元素
func (d *Deque) Front() interface{} {
    return d.lst.Front().Value
}

// 判断队列是否为空
func (d *Deque) IsEmpty() bool {
    return d.lst.Len() == 0
}

// 清空队列
func (d *Deque) Clear() {
    d.lst.Init()
}