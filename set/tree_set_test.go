package set

import (
    "fmt"
    "testing"
)

type treeSetVistor struct {
	stopped bool
	stoppedElement int
}

func (t *treeSetVistor) Visit(e interface{}) {
	fmt.Println(e)
}
func (t *treeSetVistor) Stop() bool {
	return t.stopped
}

type num struct {}

func (n *num) Compare(num1 interface{}, num2 interface{}) int {
	n1, _ := num1.(int)
	n2, _ := num2.(int)
	return n1 - n2
}


func TestTreeSet(t *testing.T) {
	treeSet := NewTreeSetWithComparator(&num{})
	treeSet.Add(1)
	treeSet.Add(3)
	treeSet.Add(2)
	treeSet.Add(3)
	treeSet.Traversal(&treeSetVistor{})
}