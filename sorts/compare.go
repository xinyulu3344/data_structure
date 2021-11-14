package sorts


type Comparator interface {
    CompareTo(e1 interface{}, e2 interface{}) int
}

type Comparable interface {
    Compare(e interface{}) int
}
