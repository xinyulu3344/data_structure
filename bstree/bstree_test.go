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
        return e1.(*Person).age - e2.(*Person).age
    })
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
        bstree2.Add(&Person{
            name: "张三" + fmt.Sprint(i),
            age:  int(data[i]),
        })
    }
    
    bstree1.PreorderTraversal(func(e E) {
        fmt.Println(e)
    })
    bstree2.PreorderTraversal(func(e E) {
        fmt.Println(e)
    })
}

func TestBstree_InorderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.InorderTraversal(func(e E) {
        fmt.Println(e)
    })
}

func TestBstree_PostorderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.PostorderTraversal(func(e E) {
        fmt.Println(e)
    })
}

func TestBstree_LevelOrderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.LevelOrderTraversal(func(e E) {
        fmt.Println(e)
    })
}

func TestBstree_Height(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    t.Log(bstree1.Height(), bstree1.Height2())
}

func TestBstree_IsComplete(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 9, 2, 5}
    
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    t.Log(bstree1.IsComplete())
}
