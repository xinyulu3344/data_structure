package sorts

func IsAsSortedInts(elements []int) bool {
    length := len(elements)
    for end := length; end > 0; end-- {
        isSorted := true
        for begin := 1; begin < end; begin++ {
            if elements[begin] < elements[begin - 1] {
                isSorted = false
                return false
            }
        }
        if isSorted {
            break
        }
    }
    return true
}

func IntsAreAsSorted(elements []int) bool {
    return IsAsSorted(IntSlice(elements))
}

func IntsAreDsSorted(elements []int) bool {
    return IsDsSorted(IntSlice(elements))
}


// 判断数组是否降序
func AreDsSortedInts(elements []int) bool {
    length := len(elements)
    for end := length; end > 0; end-- {
        isSorted := true
        for begin := 1; begin < end; begin++ {
            if elements[begin] > elements[begin - 1] {
                isSorted = false
                return false
            }
        }
        if isSorted {
            break
        }
    }
    return true
}

// IndexOfInt
// elements: 元素数组
// v: 待查找元素
// return: 待查找元素在数组中的索引
func IndexOfInt(elements []int, v int) int {
    if elements == nil || len(elements) == 0 {
        return -1
    }
    begin := 0
    end := len(elements)
    for begin < end {
        mid := (end + begin) >> 1
        if v < elements[mid] {
            end = mid
        } else if v > elements[mid] {
            begin = mid + 1
        } else {
            return mid
        }
    }
    return -1
}

// Search 查找v在有序slice elements中的插入位置
func Search(elements []int, v int) int {
    if elements == nil || len(elements) == 0 {
        return -1
    }
    begin := 0
    end := len(elements)
    for begin < end {
        mid := (end + begin) >> 1
        if v < elements[mid] {
            end = mid
        } else {
            begin = mid + 1
        }
    }
    return begin
}

// IsAsSorted reports whether data is sorted.
func IsAsSorted(data Interface) bool {
    n := data.Len()
    for i := n - 1; i > 0; i-- {
        if data.Compare(i, i-1) < 0 {
            return false
        }
    }
    return true
}

// IsDsSorted reports whether data is sorted.
func IsDsSorted(data Interface) bool {
    n := data.Len()
    for i := n - 1; i > 0; i-- {
        if data.Compare(i, i-1) > 0 {
            return false
        }
    }
    return true
}
