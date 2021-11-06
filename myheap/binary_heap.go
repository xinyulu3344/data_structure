package myheap

type Comparator struct {}

type BinaryHeap struct {
    elements []interface{}
    size int
}

func NewBinaryHeap() *BinaryHeap {
    return &BinaryHeap{}
}