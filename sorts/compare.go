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
