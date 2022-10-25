package set

type MapSet[T comparable] map[T]struct{}

func NewMapSet[T comparable]() MapSet[T] {
	return make(MapSet[T])
}

func (ms *MapSet[T]) Size() int {
	return len(*ms)
}

func (ms *MapSet[T]) IsEmpty() bool {
	return ms.Size() == 0
}

func (ms *MapSet[T]) Clear() {
	*ms = NewMapSet[T]()
}

func (ms *MapSet[T]) Contains(v T) bool {
	_, ok := (*ms)[v]
	return ok
}

func (ms *MapSet[T]) Add(v T) {
	(*ms)[v] = struct{}{}
}

func (ms *MapSet[T]) Remove(v T) {
	delete(*ms, v)
}

func (ms *MapSet[T]) Traversal(v Visit[T]) {
	if v == nil {
		return
	}
	for k := range *ms {
		if v(k) {
			return
		}
	}
}