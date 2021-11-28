package sorts

type CountingSort struct {
    elements []int
}

func NewCountingSort() *CountingSort {
    return &CountingSort{}
}

// SortInt
// 最好、最坏、平均时间复杂度: O(n+k)
// 空间复杂度: O(n+k)
// k 是整数的取值范围
// 属于稳定排序
func (cs *CountingSort) SortInt(elements []int) {
    cs.elements = elements
    length := len(cs.elements)
    max := cs.elements[0]
    min := cs.elements[0]
    // 找出最大、小值
    for i := 1; i < length; i++ {
        if cs.elements[i] > max {
            max = cs.elements[i]
        }
        if cs.elements[i] < min {
            min = cs.elements[i]
        }
    }
    
    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, max - min + 1)
    countsLength := len(counts)
    // 统计每个整数出现的次数
    for i := 0; i < length; i++ {
        counts[cs.elements[i] - min]++
    }
    // 累加次数
    for i := 1; i < countsLength; i++ {
        counts[i] += counts[i - 1]
    }
    
    // 从后往前遍历元素，将它放到有序数组中的合适位置
    newElements := make([]int, length)
    for i := length - 1; i >= 0; i-- {
        counts[cs.elements[i] - min]--
        newElements[counts[cs.elements[i] - min]] = cs.elements[i]
    }
    for i := 0; i < length; i++ {
        cs.elements[i] =  newElements[i]
    }
}

// SortInt1
// 最简单的实现
// 不稳定、内存浪费极大、无法给负数排序
func (cs *CountingSort) SortInt1(elements []int) {
    cs.elements = elements
    length := len(cs.elements)
    max := cs.elements[0]
    
    // 找出最大值
    for i := 1; i < length; i++ {
        if cs.elements[i] > max {
            max = cs.elements[i]
        }
    }
    
    // 开辟内存空间，存储每个整数出现的次数
    counts := make([]int, max + 1)
    for i := 0; i < length; i++ {
        counts[cs.elements[i]]++
    }
    
    // 根据整数的出现次数，对整数进行排序
    index := 0
    countsLength := len(counts)
    for i := 0; i < countsLength; i++ {
        for ;counts[i] > 0; counts[i]-- {
            cs.elements[index] = i
            index++
        }
    }
}