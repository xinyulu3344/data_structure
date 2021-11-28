package sorts

type MergeSort struct {
    cmpCount  int
    swapCount int
    cmp Comparator
    leftElements []interface{}
    elements  []interface{}
}

func NewMergeSort() *MergeSort {
    return &MergeSort{}
}

func NewMergeSortWithComparator(cmp Comparator) *MergeSort {
    return &MergeSort{
        cmp: cmp,
    }
}

func (ms *MergeSort) SetComparator(cmp Comparator) {
    ms.cmp = cmp
}

func (ms *MergeSort) Sort(elements []interface{}) {
    ms.elements = elements
    length := len(elements)
    ms.leftElements = make([]interface{}, length >> 1)
    ms.sort(0, length)
}

// [begin, end)
func (ms *MergeSort) sort(begin, end int) {
    if end - begin < 2 { return }
    mid := (begin + end) >> 1
    ms.sort(begin, mid)
    ms.sort(mid, end)
    ms.merge(begin, mid, end)
}

// 将 [begin, mid) 和 [mid, end) 范围的序列合并成一个有序序列
func (ms *MergeSort) merge(begin, mid, end int) {
    leftIndex := 0
    leftEnd := mid - begin
    rightIndex := mid
    rightEnd := end
    allIndex := begin
    for i := leftIndex; i < leftEnd; i++ {
        ms.leftElements[i] = ms.elements[begin + i]
    }
    for leftIndex < leftEnd {
        if rightIndex < rightEnd && ms.compare(rightIndex, leftIndex) < 0 {
            ms.elements[allIndex] = ms.elements[rightIndex]
            allIndex++
            rightIndex++
        } else {
            ms.elements[allIndex] = ms.leftElements[leftIndex]
            allIndex++
            leftIndex++
        }
    }
}


func (ms *MergeSort) compare(i1, i2 int) int {
    if ms.cmp != nil {
        return ms.cmp.CompareTo(ms.elements[i1], ms.leftElements[i2])
    }
    return ms.elements[i1].(Comparable).Compare(ms.leftElements[i2])
}
