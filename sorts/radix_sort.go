package sorts

type RadixSort struct {
    elements Elements
}

func NewRadixSort() *RadixSort {
    return &RadixSort{}
}

func (rs *RadixSort) Sort(elements Elements) {
    rs.elements = elements
    max := elements.GetElement(0).GetNum()
    length := elements.Len()
    for i := 1; i < length; i++ {
        if elements.GetElement(i).GetNum() > max {
            max = elements.GetElement(i).GetNum()
        }
    }
    for divider := 1; divider < max; divider *= 10 {
        rs.countingSort(divider)
    }
}

func (rs *RadixSort) SortInt(elements []int) {
	rs.Sort(ElementsInt(elements))
}

// Sort
// 支持对象排序
func (rs *RadixSort) countingSort(divider int) {
    length := rs.elements.Len()

    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, 10)
    countsLength := len(counts)
    // 统计每个整数出现的次数
    for i := 0; i < length; i++ {
        counts[rs.elements.GetElement(i).GetNum() / divider % 10]++
    }
    // 累加次数
    for i := 1; i < countsLength; i++ {
        counts[i] += counts[i - 1]
    }
    
    // 从后往前遍历元素，将它放到有序数组中的合适位置
    newElements := make([]Element, length)
    for i := length - 1; i >= 0; i-- {
        counts[rs.elements.GetElement(i).GetNum() / divider % 10]--
        newElements[counts[rs.elements.GetElement(i).GetNum() / divider % 10]] = rs.elements.GetElement(i)
    }
    for i := 0; i < length; i++ {
        rs.elements.SetElement(i, newElements[i]) 
    }
}
