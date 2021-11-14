package sorts

type BubbleSort struct {
	elements []interface{}
	elementsInt []int
	cmpCount int
	swapCount int
	cmp Comparator
}

func NewBubbleSort() *BubbleSort {
	return &BubbleSort{
	}
}

func NewBubbleSortWithComparator(cmp Comparator) *BubbleSort {
	return &BubbleSort{
		cmp: cmp,
	}
}

func (b *BubbleSort) AsSortInt(elements []int) {
	b.cmpCount = 0
	b.swapCount = 0
	b.elementsInt = elements
	length := len(b.elementsInt)
	for end := length - 1; end > 0; end-- {
		//sorted := true
		sortedIndex := 1
		for begin := 1; begin <= end; begin++ {
			if b.compareInt(begin, begin - 1) < 0 {
				b.swapInt(begin, begin - 1)
				//sorted = false
				sortedIndex = begin
			}
		}
		//if sorted {
		//	break
		//}
		end = sortedIndex
	}
}

func (b *BubbleSort) AsSort(elements []interface{}) {
	b.cmpCount = 0
	b.swapCount = 0
	b.elements = elements
	length := len(b.elements)
	for end := length - 1; end > 0; end-- {
		//sorted := true
		sortedIndex := 1
		for begin := 1; begin <= end; begin++ {
			if b.compare(begin, begin - 1) < 0 {
				b.swap(begin, begin - 1)
				//sorted = false
				sortedIndex = begin
			}
		}
		//if sorted {
		//	break
		//}
		end = sortedIndex
	}
}

// 交换元素位置
func (b *BubbleSort) swap(index1, index2 int) {
	b.swapCount++
	tmp := b.elements[index1]
	b.elements[index1] = b.elements[index2]
	b.elements[index2] = tmp
}

func (b *BubbleSort) swapInt(index1, index2 int) {
	b.swapCount++
	tmp := b.elementsInt[index1]
	b.elementsInt[index1] = b.elementsInt[index2]
	b.elementsInt[index2] = tmp
}

// 比较指定索引位置元素大小
func (b *BubbleSort) compare(index1, index2 int) int {
	b.cmpCount++
	if b.cmp != nil {
		return b.cmp.CompareTo(b.elements[index1], b.elements[index2])
	}
	return b.elements[index1].(Comparable).Compare(b.elements[index2])
}

func (b *BubbleSort) compareInt(index1, index2 int) int {
	b.cmpCount++
	return b.elementsInt[index1] - b.elementsInt[index2]
}