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

func (is *InsertionSort) AsSortInt(elements []int) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elementsInt = elements
    length := len(elements)
    for begin := 1; begin < length; begin++ {
        cur := begin
        v := is.elementsInt[cur]
        for cur > 0 && is.compareElementsInt(v, is.elementsInt[cur - 1]) < 0 {
            is.elementsInt[cur] = is.elementsInt[cur - 1]
            cur--
        }
        is.elementsInt[cur] = v
    }
}

func (is *InsertionSort) AsSortInt1(elements []int) {
    is.cmpCount = 0
    is.swapCount = 0
    is.elementsInt = elements
    length := len(elements)
    for begin := 1; begin < length; begin++ {
        cur := begin
        for cur > 0 && is.compareInt(cur, cur - 1) < 0 {
            is.swapInt(cur, cur - 1)
            cur--
        }
    }
}

func (is *InsertionSort) AsSort(elements []interface{}) {
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

func (is *InsertionSort) AsSort1(elements []interface{}) {
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