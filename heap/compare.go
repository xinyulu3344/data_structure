package heap

type Comparable interface {
    CompareTo(e any) int
}

type Compare func(e1, e2 any) int
