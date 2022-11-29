package bstree

import (
	"fmt"
	"runtime/debug"
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
    
    bstree1.PreorderTraversal(func(e E) bool {
        fmt.Println(e)
        return false
    })
    bstree2.PreorderTraversal(func(e E) bool {
        fmt.Println(e)
        return false
    })
}

func TestBstree_InorderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.InorderTraversal(func(e E) bool {
        fmt.Println(e)
        return false
    })
}

func TestBstree_PostorderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.PostorderTraversal(func(e E) bool {
        fmt.Println(e)
        return false
    })
}

func TestBstree_LevelOrderTraversal(t *testing.T) {
    bstree1 := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    for i := 0; i < len(data); i++ {
        bstree1.Add(data[i])
    }
    bstree1.LevelOrderTraversal(func(e E) bool {
        fmt.Printf("%d ", e)
        return false
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

func TestBstree_visit(t *testing.T) {
    bstree := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
    data1 := []Int{7, 4, 2, 1, 3}
    data2 := []Int{1, 2, 3, 4, 5}
    data3 := []Int{1, 3, 2, 5, 4}
    data4 := []Int{7, 4, 9, 2, 5}
    for i := 0; i < len(data); i++ {
        bstree.Add(data[i])
    }
    
    index := 0
    bstree.PreorderTraversal(func(e E) bool {
        if e.(Int) == Int(5) {
            return true
        }
        if e.(Int) != data1[index] {
            t.Fail()
        }
        index++
        return false
    })
    
    index = 0
    bstree.InorderTraversal(func(e E) bool {
        if e.(Int) == Int(7) {
            return true
        }
        if e.(Int) != data2[index] {
            t.Fail()
        }
        index++
        return false
    })
    
    index = 0
    bstree.PostorderTraversal(func(e E) bool {
        if e.(Int) == Int(8) {
            return true
        }
        if e.(Int) != data3[index] {
            t.Fail()
        }
        index++
        return false
    })
    
    index = 0
    bstree.LevelOrderTraversal(func(e E) bool {
        if e.(Int) == Int(8) {
            return true
        }
        if e.(Int) != data4[index] {
            t.Fail()
        }
        index++
        return false
    })
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
    b.PreorderTraversal(func(e E) bool {
        data1 = append(data1, e.(Int))
        return false
    })
    b.InorderTraversal(func(e E) bool {
        data2 = append(data2, e.(Int))
        return false
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

func TestPreorderTraversalIter(t *testing.T) {
	bstree := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 9, 8, 11, 10, 12}
    for _, v := range data {
        bstree.Add(v)
    }
    bstree.PreorderTraversalIter(func(e E) bool {
        fmt.Printf("%v ", e) // 7 4 2 1 3 9 8 11 10 12
        return false
    })
    fmt.Println()
    bstree.PreorderTraversalIter2(func(e E) bool {
        fmt.Printf("%v ", e) // 7 4 2 1 3 9 8 11 10 12
        return false
    })
}

func TestInorderTraversalIter(t *testing.T) {
	bstree := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 9, 8, 11, 10, 12}
    for _, v := range data {
        bstree.Add(v)
    }
    bstree.InorderTraversalIter(func(e E) bool {
        fmt.Printf("%v ", e) // 1 2 3 4 7 8 9 10 11 12 
        return false
    })
}

func TestPostorderTraversalIter(t *testing.T) {
    bstree := NewBstree()
    data := []Int{7, 4, 2, 1, 3, 9, 8, 11, 10, 12}
    for _, v := range data {
        bstree.Add(v)
    }
    bstree.PostorderTraversalIter(func(e E) bool {
        fmt.Printf("%v ", e) // 1 3 2 4 8 10 12 11 9 7
        return false
    })
}

func test(t *testing.T, result bool) {
    if result {
        return
    }
	debug.PrintStack()
    t.Fail()
}
