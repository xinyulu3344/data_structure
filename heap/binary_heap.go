package heap

import (
	"errors"
)

const DEFAULT_CAPACITY = 10

type BinaryHeap struct {
    elements   []any
    size       int
    comparator Compare
}

func NewBinaryHeapWithComparator(comparator Compare) *BinaryHeap {
    return &BinaryHeap{
        comparator: comparator,
        elements:   make([]any, DEFAULT_CAPACITY),
    }
}

func NewBinaryHeap() *BinaryHeap {
    return &BinaryHeap{
        elements: make([]any, DEFAULT_CAPACITY),
    }
}

func NewBinaryHeapify(elements []any, comparator Compare) *BinaryHeap {
	if len(elements) == 0 {
        bheap := NewBinaryHeapWithComparator(comparator)
        return bheap
    } else {
        bheap := &BinaryHeap{
            comparator: comparator,
            elements: make([]any, max(len(elements), DEFAULT_CAPACITY)),
            size: len(elements),
        }
        copy(bheap.elements, elements)
        bheap.heapify()
        return bheap
    }
}

// 元素的数量
func (bh *BinaryHeap) Size() int {
    return bh.size
}

// 是否为空
func (bh *BinaryHeap) IsEmpty() bool {
    return bh.size == 0
}

// 清空
func (bh *BinaryHeap) Clear() {
    for i := 0; i < bh.size; i++ {
        bh.elements[i] = nil
    }
    bh.size = 0
}

// 添加元素
func (bh *BinaryHeap) Add(e any) error {
    err := bh.elementNotNullCheck(e)
    if err != nil {
        return err
    }
    bh.ensureCapacity(bh.size + 1)
    bh.elements[bh.size] = e
    bh.siftUp(bh.size)
    bh.size++
    return nil
}

// 获取堆顶元素
func (bh *BinaryHeap) Get() any {
    if bh.size == 0 {
        return nil
    }
    return bh.elements[0]
}

// 删除堆顶元素
func (bh *BinaryHeap) Remove() any {
    if bh.size == 0 {
        return nil
    }

    root := bh.elements[0]
    bh.elements[0] = bh.elements[bh.size - 1]
    bh.elements[bh.size - 1] = nil
    bh.size--
    bh.siftDown(0)
    return root
}

// 删除堆顶元素的同时插入一个新元素
func (bh *BinaryHeap) Replace(e any) (any, error) {
    err := bh.elementNotNullCheck(e)
    if err != nil {
        return nil, err
    }

    var root any
    if bh.size == 0 {
        bh.elements[0] = e
        bh.size++
    } else {
        root = bh.elements[0]
        bh.elements[0] = e
        bh.siftDown(0)
    }
    return root, nil
}

func (bh *BinaryHeap) compare(e1, e2 any) int {
    if bh.comparator != nil {
        return bh.comparator(e1, e2)
    }
    return e1.(Comparable).CompareTo(e2)
}

// 扩容数组，将数组容量扩容为原来1.5倍
func (bh *BinaryHeap) ensureCapacity(capacity int) {
    oldCapacity := len(bh.elements)
    if oldCapacity >= capacity {
        return
    }
    // 新容量为旧容量的1.5倍
    newCapacity := oldCapacity + (oldCapacity >> 1)
    newElements := make([]any, newCapacity)
    for i := 0; i < bh.size; i++ {
        newElements[i] = bh.elements[i]
    }
    bh.elements = newElements
}

func (bh *BinaryHeap) elementNotNullCheck(e any) error {
    if e == nil {
        return errors.New("element must not be nil")
    }
    return nil
}

// 让index位置的元素上滤
// 不断和父节点比较
func (bh *BinaryHeap) siftUp(index int) {
    current := bh.elements[index]
    for index > 0 {
        // 获取父节点索引
        parentIndex := (index - 1) >> 1
        // 找到父节点
        parent := bh.elements[parentIndex]
        if bh.compare(current, parent) <= 0 { // 和父节点比较，父节点大，跳出循环
            break
        }
        bh.elements[index] = parent
        index = parentIndex
    }
    bh.elements[index] = current
}


func (bh *BinaryHeap) siftDown(index int) {
    current := bh.elements[index]
    indexOfFirstLeaf := bh.size >> 1
    // 第一个叶子节点的索引 == 非叶子节点的数量
    for index < indexOfFirstLeaf { // 小于第一个叶子节点的索引
        // index 的节点有2种情况
        // 1. 只有左子节点
        // 2. 同时有左右子节点

        // 默认为左子节点的索引和index进行比较
        childIndex := (index << 1) + 1
        child := bh.elements[childIndex]

        // 右子节点索引
        rightIndex := childIndex + 1
        // 选出左右子节点大的那个
        if rightIndex < bh.size && bh.compare(bh.elements[rightIndex], child) > 0 {
            childIndex = rightIndex
            child = bh.elements[rightIndex]
        }

        if bh.compare(current, child) >= 0 {
            break
        }

        // 将子节点存放到index位置
        bh.elements[index] = child
        // 重新设置index
        index = childIndex 
    }
    bh.elements[index] = current
}

func (bh *BinaryHeap) heapify() {
	// 自上而下的上滤
	// for i := 1; i < bh.size; i++ {
	// 	bh.siftUp(i)
	// }

	// 自下而上的下滤
	for i := bh.size >> 1; i >=0; i-- {
		bh.siftDown(i)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
