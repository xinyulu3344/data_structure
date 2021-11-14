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
