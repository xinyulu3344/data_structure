package avl

type E interface {
    CompareTo(e E) int
}

type Compare func(e1, e2 E) int