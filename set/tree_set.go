package set


type TreeSet struct {
	tree *RBTree
}

func NewTreeSet() *TreeSet {
	return &TreeSet{
		tree: NewRBTree(),
	}
}

func NewTreeSetWithComparator(comparator Compare) *TreeSet {
	return &TreeSet{
		tree: NewRBTreeWithComparator(comparator),
	}
}

func (ts *TreeSet) Size() int {
	return ts.tree.Size()
}

func (ts *TreeSet) IsEmpty() bool {
	return ts.tree.IsEmpty()
}

func (ts *TreeSet) Clear() {
	ts.tree.Clear()
}

func (ts *TreeSet) Contains(e E) bool {
	return ts.tree.Contains(e)
}

func (ts *TreeSet) Add(e E) {
	ts.tree.Add(e)
}

func (ts *TreeSet) Remove(e E) {
	ts.tree.Remove(e)
}

func (ts *TreeSet) Traversal(v Visit) {
	ts.tree.InorderTraversal(v)
}

