// 选择排序
// 1. 从序列中找出最大的元素，然后与最末尾的元素交换位置
// 2. 忽略 1 中曾经找到的最大元素，重复执行步骤 1

// 最好、最坏、平均时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 属于稳定排序

package selection_sort


func SelectionSortInt(elements []int) {
    length := len(elements)
    for end := length - 1; end > 0; end-- {
        maxIndex := 0
        // 找出最大值
        for begin := 1; begin <= end; begin++ {
            if elements[maxIndex] <= elements[begin] {
                maxIndex = begin
            }
        }
        // 交换
        tmp := elements[maxIndex]
        elements[maxIndex] = elements[end]
        elements[end] = tmp
    }
}