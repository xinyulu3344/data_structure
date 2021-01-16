package binary_tree

import (
    "fmt"
    "testing"
)

// 实现了Comparable接口, 该类型可以被比较
type persion struct {
    name string
    age  int
}

func NewPersion(name string, age int) *persion {
    return &persion{
        name: name,
        age: age,
    }
}

func (p *persion) compareTo(e Comparable) int {
    v, ok := e.(*persion)
    if !ok {
        panic("传入的值不是*persion类型!")
    }
    return p.age - v.age
}

type PersionComparator1 struct {
}


func (p *PersionComparator1) compare(e1 interface{}, e2 interface{}) int {
    return e1.(*persion).age - e2.(*persion).age
}

type PersionComparator2 struct {
}

func (p *PersionComparator2) compare(e1 interface{}, e2 interface{}) int {
    return e2.(*persion).age - e1.(*persion).age
}


// 遍历二叉树元素的时候，何时终止遍历，以及如何处理二叉树的值
type myVistor struct {
    stoped bool
}

func (m *myVistor) visit(e interface{})  {
    element, ok := e.(*persion)
    if !ok {
        fmt.Println("断言错误")
    }
    fmt.Printf("%d,", element.age)
    // 如果符合条件，就将stopped置为true，表示要停止遍历
    if element.age == 3 {
       m.stoped = false
    }
}

func (m *myVistor) stop() bool {
    return m.stoped
}

func TestBST(t *testing.T) {

    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst1 := NewBinarySearchTree()
    bst2 := NewBinarySearchTreeWithComparator(&PersionComparator1{})
    bst3 := NewBinarySearchTreeWithComparator(&PersionComparator2{})
    //for _, v := range data {
    //   bst1.Add(&v)
    //}
    for i := 0; i < len(data); i++ {
        bst1.Add(&data[i])
    }
    for i := 0; i < len(data); i++ {
        bst2.Add(&data[i])
    }
    for i := 0; i < len(data); i++ {
        bst3.Add(&data[i])
    }
    fmt.Println(bst1)
    fmt.Println(bst2)
    fmt.Println(bst3)
}


// 测试前序遍历
func TestPreorderTraversal(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    bst.PreOrderTraversal(&myVistor{})
}

// 测试中序遍历
func TestInnerOrderTraversal(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    bst.InnerOrderTraversal(&myVistor{})
}

// 测试后序遍历
func TestPostOrderTraversal(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    bst.PostOrderTraversal(&myVistor{})
}

func TestBinarySearchTree_LevelOrder(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    bst.LevelOrder(&myVistor{})
}

func TestBinarySearchTree_GetTreeHeight(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    fmt.Println(bst.GetTreeHeight())
}

func TestAllTraversal(t *testing.T) {
    fmt.Println("前序遍历: ")
    TestPreorderTraversal(t)
    fmt.Println("\n中序遍历: ")
    TestInnerOrderTraversal(t)
    fmt.Println("\n后序遍历: ")
    TestPostOrderTraversal(t)
    fmt.Println("\n层序遍历: ")
    TestBinarySearchTree_LevelOrder(t)
    fmt.Printf("\n二叉树高度: ")
    TestBinarySearchTree_GetTreeHeight(t)
}


func TestBinarySearchTree_IsComplete(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 1},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    fmt.Println(bst.IsComplete())
}

func TestGetPredecessor(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    n := bst.getPredecessor(bst.root.left)
    fmt.Println("前驱结点: ", n.element)
}


func TestBinarySearchTree_Remove(t *testing.T) {
    data := []persion{
        {name: "a", age: 7},
        {name: "b", age: 4},
        {name: "c", age: 9},
        {name: "d", age: 2},
        {name: "e", age: 5},
        {name: "f", age: 8},
        {name: "g", age: 11},
        {name: "i", age: 8},
        {name: "h", age: 3},
    }
    bst := NewBinarySearchTree()
    for i := 0; i < len(data); i++ {
        bst.Add(&data[i])
    }
    bst.LevelOrder(&myVistor{})
    bst.Remove(&data[3])
    fmt.Println()
    bst.LevelOrder(&myVistor{})
}
