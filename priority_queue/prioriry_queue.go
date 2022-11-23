package priorityqueue

import "data_structure/heap"

type PriorityQueue struct {
	h heap.IHeap
}

func NewPrioriryQueue() *PriorityQueue {
	return &PriorityQueue{
		heap.NewBinaryHeap(),
	}
}

func NewPrioriryQueueWithComparator(comparator heap.Compare) *PriorityQueue {
	return &PriorityQueue{
		heap.NewBinaryHeapWithComparator(comparator),
	}
}

func (q *PriorityQueue) Size() int {
	return q.h.Size()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.h.IsEmpty()
}

func (q *PriorityQueue) EnQueue(e any) {
	q.h.Add(e)
}

func (q *PriorityQueue) Dequeue() any {
	return q.h.Remove()
}

func (q *PriorityQueue) Front() any {
	return q.h.Get()
}

func (q *PriorityQueue) Clear() {
	q.h.Clear()
}

