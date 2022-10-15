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

func TestBstree_Remove(t *testing.T) {
    bstree := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 9, 8, 11, 10, 12}
    data1 := []Int{7, 2, 1, 3, 9, 8, 11, 10, 12}
    data2 := []Int{1, 2, 3, 7, 8, 9, 10, 11, 12}
    
    data3 := []Int{7, 2, 1, 3, 10, 8, 11, 12}
    data4 := []Int{1, 2, 3, 7, 8, 10, 11, 12}
    
    for i := 0; i < len(data); i++ {
        bstree.Add(data[i])
    }
    bstree.Remove(Int(4))
    test(t, judgeBst(bstree, data1, data2))
    test(t, bstree.Size() == 9)
    bstree.Remove(Int(9))
    test(t, judgeBst(bstree, data3, data4))
    test(t, bstree.Size() == 8)
}

// 前序遍历、中序遍历二叉树，与传入的正确结果比对
func judgeBst(b *Bstree, preorderData []Int, inorderData []Int) bool {
    var data1 []Int
    var data2 []Int
    b.PreorderTraversal(func(e E) {
        data1 = append(data1, e.(Int))
    })
    b.InorderTraversal(func(e E) {
        data2 = append(data2, e.(Int))
    })
    for i, v := range data1 {
        if v != preorderData[i] {
            return false
        }
    }
    for i, v := range data2 {
        if v != inorderData[i] {
            return false
        }
    }
    return true
}

func test(t *testing.T, result bool) {
    if result {
        return
    }
    t.Fail()
}
