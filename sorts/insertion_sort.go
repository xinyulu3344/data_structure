package sorts

type InsertionSort struct {
    elements []interface{}
    elementsInt []int
    cmp Comparator
    cmpCount int
    swapCount int
}

func NewInsertionSort() *InsertionSort {
    return &InsertionSort{}
}

func NewInsertionSortWithComparator(cmp Comparator) *InsertionSort {
    return &InsertionSort{
        cmp: cmp,
    }
}

// 二分法插入排序，优化插入位置查找
func (is *InsertionSort) AsSortInt(elements []int) {
   is.elementsInt = elements
   length := len(elements)
   for begin := 1; begin < length; begin++ {
       is.insertInt(begin, is.searchInt(begin))
   }
}

// 将待插入元素不断和前面的元素比较，如果比前面元素小，就将前面的元素往后挪，直到比前面元素大为止
func (is *InsertionSort) AsSortIntMove(elements []int) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elementsInt = elements
    length := len(elements)
    // 遍历元素切片
    for begin := 1; begin < length; begin++ {
        // 保存待插入元素索引
        cur := begin
        // 保存待插入元素
        v := is.elementsInt[cur]
        // 循环比较待插入元素和前面的元素的大小，如果比前面的小，就将前面的元素往后挪
        for cur > 0 && is.compareElementsInt(v, is.elementsInt[cur - 1]) < 0 {
            is.elementsInt[cur] = is.elementsInt[cur - 1]
            cur--
        }
        // 已经找到合适的位置插入元素
        is.elementsInt[cur] = v
    }
}

// 将待插入元素不断和前面的元素交换位置，直到比前面的元素大为止
func (is *InsertionSort) AsSortIntSwap(elements []int) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elementsInt = elements
    length := len(elements)
    // 遍历元素切片
    for begin := 1; begin < length; begin++ {
        // 记录待插入元素索引
        cur := begin
        // 循环比较待插入元素和前面的元素的大小，如果比前面的元素小，就和前面的元素交换位置
        for cur > 0 && is.compareInt(cur, cur - 1) < 0 {
            is.swapInt(cur, cur - 1)
            cur--
        }
    }
}

// 二分法插入排序，优化插入位置查找
func (is *InsertionSort) AsSort(elements []interface{}) {
    is.elements = elements
    length := len(elements)
    for begin := 1; begin < length; begin++ {
        is.insert(begin, is.search(begin))
    }
}

func (is *InsertionSort) AsSortMove(elements []interface{}) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elements = elements
    length := len(elements)
    for begin := 1; begin < length; begin++ {
        cur := begin
        v := is.elements[cur]
        for cur > 0 && is.compareElements(v, is.elements[cur - 1]) < 0 {
            is.elements[cur] = is.elements[cur - 1]
            cur--
        }
        is.elements[cur] = v
    }
}

func (is *InsertionSort) AsSortSwap(elements []interface{}) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elements = elements
    length := len(elements)
    for begin := 1; begin < length; begin++ {
        cur := begin
        for cur > 0 && is.compare(cur, cur - 1) < 0 {
            is.swap(cur, cur - 1)
            cur--
        }
    }
}

func (is *InsertionSort) swap(index1, index2 int) {
    is.swapCount++
    tmp := is.elements[index1]
    is.elements[index1] = is.elements[index2]
    is.elements[index2] = tmp
}

func (is *InsertionSort) swapInt(index1, index2 int) {
    is.swapCount++
    tmp := is.elementsInt[index1]
    is.elementsInt[index1] = is.elementsInt[index2]
    is.elementsInt[index2] = tmp
}

func (is *InsertionSort) compare(index1, index2 int) int {
    is.cmpCount++
    if is.cmp != nil {
        return is.cmp.CompareTo(is.elements[index1], is.elements[index2])
    }
    return is.elements[index1].(Comparable).Compare(is.elements[index2])
}

func (is *InsertionSort) compareElements(e1, e2 interface{}) int {
    is.cmpCount++
    if is.cmp != nil {
        return is.cmp.CompareTo(e1, e2)
    }
    return e1.(Comparable).Compare(e2)
}

func (is *InsertionSort) compareInt(index1, index2 int) int {
    is.cmpCount++
    return is.elementsInt[index1] - is.elementsInt[index2]
}

func (is *InsertionSort) compareElementsInt(e1, e2 int) int {
    is.cmpCount++
    return e1 - e2
}

// 利用二分搜索，找到 index 位置元素的待插入位置
// 已经排好序 slice 的区间范围是 [0, index)
func (is *InsertionSort) searchInt(index int) int {
    begin := 0
    end := index
    for begin < end {
        mid := (begin + end) >> 1
        if is.compareInt(index, mid) < 0 {
            end = mid
        } else {
            begin = mid + 1
        }
    }
    return begin
}

// 将 dest 位置的元素插入到 dest 位置
func (is *InsertionSort) insertInt(src int, dest int) {
    v := is.elementsInt[src]
    for i := src; i > dest; i-- {
        is.elementsInt[i] = is.elementsInt[i - 1]
    }
    is.elementsInt[dest] = v
}

// 将 dest 位置的元素插入到 dest 位置
func (is *InsertionSort) insert(src int, dest int) {
    v := is.elements[src]
    for i := src; i > dest; i-- {
        is.elements[i] = is.elements[i - 1]
    }
    is.elements[dest] = v
}

// 利用二分搜索，找到 index 位置元素的待插入位置
// 已经排好序 slice 的区间范围是 [0, index)
func (is *InsertionSort) search(index int) int {
    begin := 0
    end := index
    for begin < end {
        mid := (begin + end) >> 1
        if is.compare(index, mid) < 0 {
            end = mid
        } else {
            begin = mid + 1
        }
    }
    return begin
}
