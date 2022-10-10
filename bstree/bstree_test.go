package bstree

import (
    "fmt"
    "testing"
)

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

func TestBstree(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    
    bstree2 := NewBstreeWithComparator(func(e1, e2 E) int {
		return int(e1.(Int) - e2.(Int))
	})
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
        bstree2.Add(&Person{
            name: "张三" + fmt.Sprint(i),
            age:  int(data[i]),
        })
    }
    
    bstree1.PreorderTraversal()
    bstree2.PreorderTraversal()
}

func TestBstree_InorderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.InorderTraversal()
}

func TestBstree_PostorderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.PostorderTraversal()
}

func TestBstree_LevelOrderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
	}
    bstree1.LevelOrderTraversal()
}