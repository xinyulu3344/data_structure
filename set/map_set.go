package set

type MapSet map[E]struct{}

func NewMapSet() MapSet {
	return make(MapSet)
}

func (ms *MapSet) Size() int {
	return len(*ms)
}

func (ms *MapSet) IsEmpty() bool {
	return ms.Size() == 0
}

func (ms *MapSet) Clear() {
	*ms = NewMapSet()
}

func (ms *MapSet) Contains(v E) bool {
	_, ok := (*ms)[v]
	return ok
}

func (ms *MapSet) Add(v E) {
	(*ms)[v] = struct{}{}
}

func (ms *MapSet) Remove(v E) {
	delete(*ms, v)
}

func (ms *MapSet) Traversal(v Visit) {
	if v == nil {
		return
	}
	for k := range *ms {
		if v(k) {
			return
		}
	}
}