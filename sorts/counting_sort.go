package sorts

type CountingSort struct {
    elements Elements
}

func NewCountingSort() *CountingSort {
    return &CountingSort{}
}


// Sort
// 支持对象排序
func (cs *CountingSort) Sort(elements Elements) {
    cs.elements = elements
    length := cs.elements.Len()
    max := cs.elements.GetElement(0).GetNum()
    min := cs.elements.GetElement(0).GetNum()
    // 找出最大、小值
    for i := 1; i < length; i++ {
        if cs.elements.GetElement(i).GetNum() > max {
            max = cs.elements.GetElement(i).GetNum()
        }
        if cs.elements.GetElement(i).GetNum() < min {
            min = cs.elements.GetElement(i).GetNum()
        }
    }
    
    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, max - min + 1)
    countsLength := len(counts)
    // 统计每个整数出现的次数
    for i := 0; i < length; i++ {
        counts[cs.elements.GetElement(i).GetNum() - min]++
    }
    // 累加次数
    for i := 1; i < countsLength; i++ {
        counts[i] += counts[i - 1]
    }
    
    // 从后往前遍历元素，将它放到有序数组中的合适位置
    newElements := make([]Element, length)
    for i := length - 1; i >= 0; i-- {
        counts[cs.elements.GetElement(i).GetNum() - min]--
        newElements[counts[cs.elements.GetElement(i).GetNum() - min]] = cs.elements.GetElement(i)
    }
    for i := 0; i < length; i++ {
        cs.elements.SetElement(i, newElements[i])
    }
}

// SortInt
// 最好、最坏、平均时间复杂度: O(n+k)
// 空间复杂度: O(n+k)
// k 是整数的取值范围
// 属于稳定排序
func (cs *CountingSort) SortInt(elements []int) {
	cs.Sort(ElementsInt(elements))
}