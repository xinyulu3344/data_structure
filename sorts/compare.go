package sorts


type Comparator interface {
    CompareTo(e1 interface{}, e2 interface{}) int
}

type Comparable interface {
    Compare(e interface{}) int
}

type IntComparator int

func (i *IntComparator) CompareTo(e1 interface{}, e2 interface{}) int {
    return e1.(int) - e2.(int)
}

type Integer int

func (i Integer) Compare(e interface{}) int {
    return int(i - e.(Integer))
}

func (i Integer) CompareTo(e1 interface{}, e2 interface{}) int {
    return int(e1.(Integer) - e2.(Integer))
}

type Interface interface {
    Len() int
    Compare(i, j int) int
    Swap(i, j int)
}

type IntSlice []int

func (is IntSlice) Len() int { return len(is) }
func (is IntSlice) Compare(i, j int) int { return is[i] - is[j] }
func (is IntSlice) Swap(i, j int) { is[i], is[j] = is[j], is[i] }


type Element interface {
    GetNum() int // 获取用来排序的整数值
}

type Elements interface{
    // 获取指定索引的元素
	GetElement(index int) Element
	SetElement(index int, element Element)
	Len() int
}

type ElementInt int

func (ei ElementInt) GetNum() int {
	return int(ei)
}

type ElementsInt []int

func (ei ElementsInt) Len() int {
	return len(ei)
}

func (ei ElementsInt) GetElement(index int) Element {
	return ElementInt(ei[index])
}

func (ei ElementsInt) SetElement(index int, element Element) {
	ei[index] = int(element.(ElementInt))
}