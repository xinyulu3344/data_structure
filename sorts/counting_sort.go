package sorts

type CountingSort struct {
    elementsInt []int
    elements Elements
}

func NewCountingSort() *CountingSort {
    return &CountingSort{}
}


// Sort
// 支持对象排序
func (cs *CountingSort) Sort(elements Elements) {
    cs.elements = elements
    length := len(cs.elements)
    max := cs.elements[0].GetNum()
    min := cs.elements[0].GetNum()
    // 找出最大、小值
    for i := 1; i < length; i++ {
        if cs.elements[i].GetNum() > max {
            max = cs.elements[i].GetNum()
        }
        if cs.elements[i].GetNum() < min {
            min = cs.elements[i].GetNum()
        }
    }
    
    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, max - min + 1)
    countsLength := len(counts)
    // 统计每个整数出现的次数
    for i := 0; i < length; i++ {
        counts[cs.elements[i].GetNum() - min]++
    }
    // 累加次数
    for i := 1; i < countsLength; i++ {
        counts[i] += counts[i - 1]
    }
    
    // 从后往前遍历元素，将它放到有序数组中的合适位置
    newElements := make(Elements, length)
    for i := length - 1; i >= 0; i-- {
        counts[cs.elements[i].GetNum() - min]--
        newElements[counts[cs.elements[i].GetNum() - min]] = cs.elements[i]
    }
    for i := 0; i < length; i++ {
        cs.elements[i] =  newElements[i]
    }
}

// SortInt
// 最好、最坏、平均时间复杂度: O(n+k)
// 空间复杂度: O(n+k)
// k 是整数的取值范围
// 属于稳定排序
func (cs *CountingSort) SortInt(elements []int) {
    cs.elementsInt = elements
    length := len(cs.elementsInt)
    max := cs.elementsInt[0]
    min := cs.elementsInt[0]
    // 找出最大、小值
    for i := 1; i < length; i++ {
        if cs.elementsInt[i] > max {
            max = cs.elementsInt[i]
        }
        if cs.elementsInt[i] < min {
            min = cs.elementsInt[i]
        }
    }
    
    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, max - min + 1)
    countsLength := len(counts)
    // 统计每个整数出现的次数
    for i := 0; i < length; i++ {
        counts[cs.elementsInt[i] - min]++
    }
    // 累加次数
    for i := 1; i < countsLength; i++ {
        counts[i] += counts[i - 1]
    }
    
    // 从后往前遍历元素，将它放到有序数组中的合适位置
    newElements := make([]int, length)
    for i := length - 1; i >= 0; i-- {
        counts[cs.elementsInt[i] - min]--
        newElements[counts[cs.elementsInt[i] - min]] = cs.elementsInt[i]
    }
    for i := 0; i < length; i++ {
        cs.elementsInt[i] =  newElements[i]
    }
}

// SortInt1
// 最简单的实现
// 不稳定、内存浪费极大、无法给负数排序
func (cs *CountingSort) SortInt1(elements []int) {
    cs.elementsInt = elements
    length := len(cs.elements)
    max := cs.elementsInt[0]
    
    // 找出最大值
    for i := 1; i < length; i++ {
        if cs.elementsInt[i] > max {
            max = cs.elementsInt[i]
        }
    }
    
    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, max + 1)
    for i := 0; i < length; i++ {
        counts[cs.elementsInt[i]]++
    }
    
    // 根据整数的出现次数，对整数进行排序
    index := 0
    countsLength := len(counts)
    for i := 0; i < countsLength; i++ {
        for ;counts[i] > 0; counts[i]-- {
            cs.elementsInt[index] = i
            index++
        }
    }
}