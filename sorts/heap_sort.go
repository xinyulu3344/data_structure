// 堆排序
/*
1. 对序列进行原地建堆(heapify)
2. 重复执行以下操作
  1) 交换堆顶元素，直到堆的元素数量为1
  2) 堆的元素数量减1
  3) 对0位置进行一次siftDown操作

最好、最坏、平均时间复杂度：O(nlogn)
空间复杂度：O(1)
属于不稳定排序
*/

package sorts

type HeapSort struct {
    heapSize int
    elements []interface{}
    elementsInt []int
    cmpCount int
    swapCount int
    cmp Comparator
}

func NewHeapSort() *HeapSort {
    return &HeapSort{}
}

func NewHeapSortWithComparator(cmp Comparator) *HeapSort {
    return &HeapSort{
        cmp: cmp,
    }
}

func (h *HeapSort) AsSortInt(elements []int) {
    h.cmpCount = 0
    h.swapCount = 0
    
    // 原地建堆
    h.elementsInt = elements
    h.heapSize = len(elements)
    h.heapifyInt()
    for h.heapSize > 1 {
       h.swapInt(0, h.heapSize - 1)
       h.heapSize--
    
       // 恢复堆的性质
       h.siftDownInt(0)
    }
}

func (h *HeapSort) AsSort(elements []interface{}) {
    h.cmpCount = 0
    h.swapCount = 0
    
    h.elements = elements
    h.heapSize = len(elements)
    h.heapify()
    for h.heapSize > 1 {
        h.swap(0, h.heapSize - 1)
        h.heapSize--
        
        // 恢复堆的性质
        h.siftDown(0)
    }
}

func (h *HeapSort) heapifyInt() {
    // 自下而上的下滤
    for i := (h.heapSize >> 1) - 1; i >= 0; i-- {
        h.siftDownInt(i)
    }
}

func (h *HeapSort) heapify() {
    // 自下而上的下滤
    for i := (h.heapSize >> 1) - 1; i >= 0; i-- {
        h.siftDown(i)
    }
}

func (h *HeapSort) siftDownInt(index int) {
    e := h.elementsInt[index]
    half := h.heapSize >> 1
    // 第一个叶子节点的索引 == 非叶子节点的数量
    // 必须保证index位置是非叶子节点
    for index <  half {
        // index的节点有2种情况
        // 1. 只有左子节点
        // 2. 同时有左右子节点
        
        // 默认为左子节点的索引跟它进行比较
        childIndex := (index << 1) + 1
        child := h.elementsInt[childIndex]
        
        // 右子节点
        rightIndex := childIndex + 1
        // 选出左右子节点最大的那个
        if rightIndex < h.heapSize && h.compareInt(h.elementsInt[rightIndex], child) > 0 {
            childIndex = rightIndex
            child = h.elementsInt[rightIndex]
        }
        
        if h.compareInt(e, child) >= 0 {
            break
        }
        // 将子节点存放到index位置
        h.elementsInt[index] = child
        // 重新设置index
        index = childIndex
    }
    h.elementsInt[index] = e
}

func (h *HeapSort) siftDown(index int) {
    e := h.elements[index]
    half := h.heapSize >> 1
    // 第一个叶子节点的索引 == 非叶子节点的数量
    // 必须保证index位置是非叶子节点
    for index <  half {
        // index的节点有2种情况
        // 1. 只有左子节点
        // 2. 同时有左右子节点
        
        // 默认为左子节点的索引跟它进行比较
        childIndex := (index << 1) + 1
        child := h.elements[childIndex]
        
        // 右子节点
        rightIndex := childIndex + 1
        // 选出左右子节点最大的那个
        if rightIndex < h.heapSize && h.compareElement(h.elements[rightIndex], child) > 0 {
            childIndex = rightIndex
            child = h.elements[rightIndex]
        }
        
        if h.compareElement(e, child) >= 0 {
            break
        }
        // 将子节点存放到index位置
        h.elements[index] = child
        // 重新设置index
        index = childIndex
    }
    h.elements[index] = e
}

func (h *HeapSort) swapInt(index1, index2 int) {
    h.swapCount++
    tmp := h.elementsInt[index1]
    h.elementsInt[index1] = h.elementsInt[index2]
    h.elementsInt[index2] = tmp
}

func (h *HeapSort) swap(index1, index2 int) {
    h.swapCount++
    tmp := h.elements[index1]
    h.elements[index1] = h.elements[index2]
    h.elements[index2] = tmp
}

func (h *HeapSort) compareInt(v1, v2 int) int {
    h.cmpCount++
    return v1 - v2
}

func (h *HeapSort) compareElement(e1 interface{}, e2 interface{}) int {
    h.cmpCount++
    h.cmpCount++
    if h.cmp != nil {
        return h.cmp.CompareTo(e1, e2)
    }
    return e1.(Comparable).Compare(e2)
}