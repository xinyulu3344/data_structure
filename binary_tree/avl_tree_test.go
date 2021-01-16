/**
  @author: xinyulu
  @date: 2020/12/13 1:57
  @note: 
**/
package binary_tree

import (
    "fmt"
    "testing"
)

type avlElement int

func (p avlElement) compareTo(e Comparable) int {
    v, ok := e.(avlElement)
    if !ok {
        panic("传入的值不是avlElement类型!")
    }
    return int(p - v)
}

// 遍历二叉树元素的时候，何时终止遍历，以及如何处理二叉树的值
type myAvlVistor struct {
    stoped bool
}

func (m *myAvlVistor) visit(e interface{})  {
    element, ok := e.(avlElement)
    if !ok {
        fmt.Println("断言错误")
    }
    fmt.Printf("%d,", element)
    // 如果符合条件，就将stopped置为true，表示要停止遍历
    if element == 3 {
        m.stoped = false
    }
}

func (m *myAvlVistor) stop() bool {
    return m.stoped
}

// 打印avl树
func TestPrintAvlTree(t *testing.T) {
    data := []avlElement {28, 47, 9, 66, 11, 52, 54, 25, 5, 64, 88, 34, 96, 100, 12}
    bst := NewBinarySearchTree()
    avlTree := &AVLTree{
        bst,
    }
    for _, e := range data {
        avlTree.Add(e)
    }
    avlTree.LevelOrder(&myAvlVistor{})
    fmt.Println()
    avlTree.Remove(data[3])
    avlTree.LevelOrder(&myAvlVistor{})
}
