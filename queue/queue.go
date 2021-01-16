/**
 * @Author: xinyulu
 * @Date: 2020-11-30 16:11:50
 * @Description: 队列
**/
package queue

import "container/list"

type Queue struct {
    lst *list.List
}

func NewQueue() *Queue {
    queue := &Queue{}
    queue.lst = list.New()
    return queue
}

// 获取队列长度
func (q *Queue) GetSize() int {
    return q.lst.Len()
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
    return q.lst.Len() == 0
}

// 队尾入队
func (q *Queue) EnQueue(e interface{}) {
    q.lst.PushBack(e)
}

// 队头出队
func (q *Queue) DeQueue() (e interface{}) {
    return q.lst.Remove(q.lst.Front())
}

// 查看队头元素
func (q *Queue) Front() (e interface{}) {
    return q.lst.Front().Value
}

// 清空队列
func (q *Queue) Clear() {
    q.lst.Init()
}
