package sorts

type BubbleSort struct {
    cmpCount  int
    swapCount int
}

func NewBubbleSort() *BubbleSort {
    return &BubbleSort{}
}

func (b *BubbleSort) AsSortInt(elements []int) {
    b.AsSort(IntSlice(elements))
}

func (b *BubbleSort) DsSortInt(elements []int) {
    b.DsSort(IntSlice(elements))
}

func (b *BubbleSort) AsSort(elements Interface) {
    b.Sort(elements, true)
}

func (b *BubbleSort) DsSort(elements Interface) {
    b.Sort(elements, false)
}

func (b *BubbleSort) Sort(elements Interface, ascend bool) {
    b.cmpCount = 0
    b.swapCount = 0
    length := elements.Len()
    if ascend {
        for end := length - 1; end > 0; end-- {
            //sorted := true
            sortedIndex := 1
            for begin := 1; begin <= end; begin++ {
                b.cmpCount++
                if elements.Compare(begin, begin-1) < 0 {
                    //b.swap(begin, begin - 1)
                    b.swapCount++
                    elements.Swap(begin, begin-1)
                    //sorted = false
                    sortedIndex = begin
                }
            }
            //if sorted {
            //	break
            //}
            end = sortedIndex
        }
        
    } else {
        for end := length - 1; end > 0; end-- {
            //sorted := true
            sortedIndex := 1
            for begin := 1; begin <= end; begin++ {
                b.cmpCount++
                if elements.Compare(begin, begin-1) > 0 {
                    //b.swap(begin, begin - 1)
                    b.swapCount++
                    elements.Swap(begin, begin-1)
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
}
