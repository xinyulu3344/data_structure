package myheap

const default_capacity = 10

type Comparator interface {
    CompareTo(e1 interface{}, e2 interface{}) int
}

type Comparable interface {
    Compare(e interface{}) int
}

type BinaryHeap struct {
    comparator Comparator
    elements []interface{}
    size int
}

func NewBinaryHeap() *BinaryHeap {
    return &BinaryHeap{
        comparator: nil,
        elements: make([]interface{}, 0, default_capacity),
        size: 0,
    }
}

func NewBinaryHeapWithComparator(cmp Comparator) *BinaryHeap {
    return &BinaryHeap{
        comparator: cmp,
        elements: make([]interface{}, 0, default_capacity),
        size: 0,
    }
}

func (b *BinaryHeap) Size() int {
    return b.size
}

func (b *BinaryHeap) IsEmpty() bool {
    return b.size == 0
}

func (b *BinaryHeap) Clear() {
    for i := 0; i < b.size; i++ {
        b.elements[i] = nil
    }
    b.size = 0
}

func (b *BinaryHeap) Add(e interface{}) {
    b.elementNotNilCheck(e)
    b.elements = append(b.elements, e)
    b.size++
    b.siftUp(b.size - 1)
}

func (b *BinaryHeap) Get() interface{} {
    b.emptyCheck()
    return b.elements[0]
}

func (b *BinaryHeap) Remove() interface{} {
    b.emptyCheck()
    lastIndex := b.size - 1
    root := b.elements[0]
    b.elements[0] = b.elements[lastIndex]
    //b.elements[lastIndex] = nil
    b.elements = b.elements[:b.size-1]
    b.size--
    b.siftDown(0)
    return root
}


func (b *BinaryHeap) Replace(e interface{}) interface{} {
    b.elementNotNilCheck(e)
    
    var root interface{}
    if b.size == 0 {
        b.elements[0] = e
        b.size++
    } else {
        root = b.elements[0]
        b.elements[0] = e
        b.siftDown(0)
    }
    return root
}


func (b *BinaryHeap) compare(e1 interface{}, e2 interface{}) int {
    if b.comparator != nil {
        return b.comparator.CompareTo(e1, e2)
    } else {
        return e1.(Comparable).Compare(e2.(Comparable))
    }
}

func (b *BinaryHeap) emptyCheck() {
    if b.size == 0 {
        panic("index out of bounds, Heap is empty")
    }
}

func (b *BinaryHeap) elementNotNilCheck(e interface{}) {
    if e == nil {
        panic("element must not be empty")
    }
}

// 让index位置的元素上滤
func (b *BinaryHeap) siftUp(index int) {
    e := b.elements[index]
    for index > 0 {
        pindex := (index - 1) >> 1
        p := b.elements[pindex]
        if b.compare(e, p) <= 0 {
            break
        }
        
        // 交换index、pindex位置的内容
        //tmp := b.elements[index]
        //b.elements[index] = b.elements[pindex]
        //b.elements[pindex] = tmp
        
        // 将父元素存储在index位置
        b.elements[index] = p
        
        index = pindex
    }
    b.elements[index] = e
}

func (b *BinaryHeap) siftDown(index int) {
    e := b.elements[index]
    half := b.size >> 1
    // 第一个叶子节点的索引 == 非叶子节点的数量
    // 必须保证index位置是非叶子节点
    for index <  half {
        // index的节点有2种情况
        // 1. 只有左子节点
        // 2. 同时有左右子节点
        
        // 默认为左子节点的索引跟它进行比较
        childIndex := (index << 1) + 1
        child := b.elements[childIndex]
        
        // 右子节点
        rightIndex := childIndex + 1
        // 选出左右子节点最大的那个
        if rightIndex < b.size && b.compare(b.elements[rightIndex], child) > 0 {
            childIndex = rightIndex
            child = b.elements[rightIndex]
        }
        
        if b.compare(e, child) >= 0 {
            break
        }
        // 将子节点存放到index位置
        b.elements[index] = child
        // 重新设置index
        index = childIndex
    }
    b.elements[index] = e
}