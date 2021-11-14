package bubble_sort

type Comparator interface {
	CompareTo(e1 interface{}, e2 interface{}) int
}

type Comparable interface {
	Compare(e interface{}) int
}

func BubbleSortInt(elements []int) {
	length := len(elements)
	for end := length - 1; end > 0; end-- {
		sortedIndex := 1
		//sorted := true
		for begin := 1; begin <= end; begin++ {
			if elements[begin] < elements[begin - 1] {
				tmp := elements[begin]
				elements[begin] = elements[begin-1]
				elements[begin-1] = tmp
				sortedIndex = begin
				//sorted = false
			}
		}
		end = sortedIndex
		//if sorted {
		//	break
		//}
	}
}

func BubbleSort(elements []interface{}) []interface{} {
	length := len(elements)
	for end := length - 1; end > 0; end-- {
		sorted := true
		for begin := 1; begin <= end; begin++ {
			if elements[begin].(Comparable).Compare(elements[begin - 1]) < 0 {
				tmp := elements[begin]
				elements[begin] = elements[begin-1]
				elements[begin-1] = tmp
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
	return elements
}


func BubbleSortWithComparator(elements []interface{}, cmp Comparator) []interface{} {
	length := len(elements)
	if cmp == nil {
		panic("Comparator must be not empty!")
	}
	// 每轮比较，如果发现已经完全排好序了，就直接跳出循环
	for end := length - 1; end > 0; end-- {
		sorted := true
		for begin := 1; begin <= end; begin++ {
			if cmp.CompareTo(elements[begin], elements[begin-1]) < 0 {
				tmp := elements[begin]
				elements[begin] = elements[begin-1]
				elements[begin-1] = tmp
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
	return elements
}

func BubbleSortWithComparator1(elements []interface{}, cmp Comparator) []interface{} {
    length := len(elements)
    if cmp == nil {
        panic("Comparator must be not empty!")
    }
    for end := length - 1; end > 0; end-- {
        // sortedIndex的初始值在数组完全有序的时候有用
        sortedIndex := 1
        for begin := 1; begin <= end; begin++ {
            if cmp.CompareTo(elements[begin], elements[begin-1]) < 0 {
                tmp := elements[begin]
                elements[begin] = elements[begin-1]
                elements[begin-1] = tmp
                sortedIndex = begin
            }
        }
        end = sortedIndex
    }
    return elements
}
