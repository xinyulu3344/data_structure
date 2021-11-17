// 选择排序
// 1. 从序列中找出最大的元素，然后与最末尾的元素交换位置
// 2. 忽略 1 中曾经找到的最大元素，重复执行步骤 1

// 最好、最坏、平均时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 属于稳定排序

package sorts

type SelectionSort struct {
    elements []interface{}
    elementsInt []int
    cmpCount int
    swapCount int
    cmp Comparator
}

func NewSelectionSort() *SelectionSort {
    return &SelectionSort{}
}

func NewSelectionSortWithComparator(cmp Comparator) *SelectionSort {
    return &SelectionSort{
        cmp: cmp,
    }
}

func (s *SelectionSort) AsSortInt(elements []int) {
    s.cmpCount = 0
    s.swapCount = 0
    s.elementsInt = elements
    length := len(s.elementsInt)
    for end := length - 1; end > 0; end-- {
        maxIndex := 0
        // 找出最大值
        for begin := 1; begin <= end; begin++ {
            if s.compareInt(maxIndex, begin) <= 0 {
                maxIndex = begin
            }
        }
        // 交换
        s.swapInt(maxIndex, end)
    }
}

func (s *SelectionSort) AsSort(elements []interface{}) {
    s.cmpCount = 0
    s.swapCount = 0
    s.elements = elements
    length := len(s.elements)
    for end := length - 1; end > 0; end-- {
        maxIndex := 0
        // 找出最大值
        for begin := 1; begin <= end; begin++ {
            if s.compare(maxIndex, begin) <= 0 {
                maxIndex = begin
            }
        }
        // 交换
        s.swap(maxIndex, end)
    }
    
}

func (s *SelectionSort) swapInt(index1, index2 int) {
    s.swapCount++
    tmp := s.elementsInt[index1]
    s.elementsInt[index1] = s.elementsInt[index2]
    s.elementsInt[index2] = tmp
}

func (s *SelectionSort) swap(index1, index2 int) {
    s.swapCount++
    tmp := s.elements[index1]
    s.elements[index1] = s.elements[index2]
    s.elements[index2] = tmp
}

func (s *SelectionSort) compareInt(index1, index2 int) int {
    s.cmpCount++
    return s.elementsInt[index1] - s.elementsInt[index2]
}

func (s *SelectionSort) compare(index1, index2 int) int {
    s.cmpCount++
    if s.cmp != nil {
        return s.cmp.CompareTo(s.elements[index1], s.elements[index2])
    }
    return s.elements[index1].(Comparable).Compare(s.elements[index2])
}