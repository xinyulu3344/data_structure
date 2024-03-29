// 选择排序
// 1. 从序列中找出最大的元素，然后与最末尾的元素交换位置
// 2. 忽略 1 中曾经找到的最大元素，重复执行步骤 1

// 最好、最坏、平均时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 属于稳定排序

package sorts

type SelectionSort struct {
    //elements []interface{}
    ascend bool
    elements Interface
    cmpCount int
    swapCount int
}

func NewSelectionSort(ascend bool) *SelectionSort {
    return &SelectionSort{
        ascend: ascend,
    }
}

func (s *SelectionSort) SetAscend(ascend bool) {
    s.ascend = ascend
}

func (s *SelectionSort) SortInt(elements []int) {
    s.Sort(IntSlice(elements))
}

func (s *SelectionSort) Sort(elements Interface) {
    s.cmpCount = 0
    s.swapCount = 0
    s.elements = elements
    length := s.elements.Len()
    if s.ascend {
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
    } else {
        for end := length - 1; end > 0; end-- {
            maxIndex := 0
            // 找出最大值
            for begin := 1; begin <= end; begin++ {
                if s.compare(maxIndex, begin) > 0 {
                    maxIndex = begin
                }
            }
            // 交换
            s.swap(maxIndex, end)
        }
    }
}

func (s *SelectionSort) swap(index1, index2 int) {
    s.swapCount++
    s.elements.Swap(index1, index2)
}

func (s *SelectionSort) compare(index1, index2 int) int {
    s.cmpCount++
    return s.elements.Compare(index1, index2)
}