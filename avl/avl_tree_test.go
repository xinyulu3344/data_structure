package avl

import "testing"

type Person struct {
    name string
    age  int
}

func (p *Person) CompareTo(e E) int {
    return p.age - e.(*Person).age
}

type Int int

func (i Int) CompareTo(e E) int {
    return int(i - e.(Int))
}

func TestXxx(t *testing.T) {
	avltree := NewAVLTree()
	data := []Int{7, 4, 2, 1, 3, 9, 8, 11, 10, 12}

	for _, v := range data {
		avltree.Add(v)
	}

	avltree.LevelOrderTraversal(func(e E) bool {
		n := avltree.getNodeByElement(e)
		if n.parent != nil {
			if n.left == nil || n.right == nil{
				t.Log(e, n.height, n.parent.e, nil, nil)
			} else {
				t.Log(e, n.height, n.parent.e, n.left.e, n.right.e)
			}
		} else {
			t.Log(e, n.height, nil, n.left, n.right)
		}
		return false
	})
}