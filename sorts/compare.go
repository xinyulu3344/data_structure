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

