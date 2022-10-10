package bstree

type E interface {
    CompareTo(e E) int
}

type Comparator interface {
    Compare(e1, e2 E) int
}
