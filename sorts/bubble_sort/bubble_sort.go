package bubble_sort

type Comparator interface {
	CompareTo(e1 interface{}, e2 interface{}) int
}

type Comparable interface {
	Compare(e interface{}) int
}

func BubbleSort(elements []interface{}) []interface{} {
	length := len(elements)
	for end := length; end > 0; end-- {
		sorted := true
		for begin := 1; begin < end; begin++ {
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
	for end := length; end > 0; end-- {
		sorted := true
		for begin := 1; begin < end; begin++ {
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