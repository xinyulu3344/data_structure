package sorts

type InsertionSort struct {
    //elements []interface{}
    ascend bool
    elements Interface
    cmpCount int
    swapCount int
}

func NewInsertionSort(ascend bool) *InsertionSort {
    return &InsertionSort{
        ascend: ascend,
    }
}

// SetAscend 设置是否升序排序
func (is *InsertionSort) SetAscend(ascend bool) {
    is.ascend = ascend
}

// SortInt 二分查找法给整数排序
func (is *InsertionSort) SortInt(elements []int) {
    is.Sort(IntSlice(elements))
}

// SortSwapInt 交换法给整数排序
func (is *InsertionSort) SortSwapInt(elements []int) {
    is.SortSwap(IntSlice(elements))
}

// Sort 二分查找法排序
func (is *InsertionSort) Sort(elements Interface) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elements = elements
    length := elements.Len()
    for begin := 1; begin < length; begin++ {
        // 将查找到的元素，插入到begin位置
        is.insert(begin, is.search(begin))
    }
}

// SortSwap 交换法排序
func (is *InsertionSort) SortSwap(elements Interface) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elements = elements
    length := elements.Len()
    if is.ascend {
        for begin := 1; begin < length; begin++ {
            cur := begin
            for cur > 0 && is.compare(cur, cur - 1) < 0 {
                is.swap(cur, cur - 1)
                cur--
            }
        }
    } else {
        for begin := 1; begin < length; begin++ {
            cur := begin
            for cur > 0 && is.compare(cur, cur - 1) > 0 {
                is.swap(cur, cur - 1)
                cur--
            }
        }
    }
}

// swap 交换指定索引元素的位置
func (is *InsertionSort) swap(index1, index2 int) {
    is.swapCount++
    is.elements.Swap(index1, index2)
}

// compare 比较指定索引元素的大小
func (is *InsertionSort) compare(index1, index2 int) int {
    is.cmpCount++
    return is.elements.Compare(index1, index2)
}


// 将 src 位置的元素插入到 dest 位置
func (is *InsertionSort) insert(src int, dest int) {
    // 先交换待插入元素和插入位置元素的位置
    is.swap(src, dest)
    lastIndex := dest + 1
    // 将交换后的插入位置元素，和前面的元素依次交换位置，来实现元素后移
    for i := src; i > lastIndex; i-- {
        is.swap(i, i - 1)
    }
}

// 利用二分搜索，找到 index 位置元素的待插入位置
// 已经排好序 slice 的区间范围是 [0, index)
func (is *InsertionSort) search(index int) int {
    begin := 0
    end := index
    if is.ascend {
        for begin < end {
            mid := (begin + end) >> 1
            if is.compare(index, mid) < 0 {
                end = mid
            } else {
                begin = mid + 1
            }
        }
        
    } else {
        for begin < end {
            mid := (begin + end) >> 1
            if is.compare(index, mid) > 0 {
                end = mid
            } else {
                begin = mid + 1
            }
        }
    }
    return begin
}