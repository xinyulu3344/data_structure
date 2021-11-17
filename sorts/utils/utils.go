package utils

// 判断数组是否升序
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


// 判断数组是否降序
func IsDesSortedInts(elements []int) bool {
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