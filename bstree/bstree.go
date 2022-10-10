package bstree

import "fmt"

type IBst interface {
    // 获取元素的数量
    Size() int
    
    // 是否为空
    IsEmpty() bool
    
    // 清空所有元素
    Clear()
    
    // 添加元素
    Add(e E)
    
    // 删除元素
    Remove(e E)
    
    // 是否包含某元素
    Contains(e E) bool
}

type node struct {
    e      E
    left   *node
    right  *node
    parent *node
}

func newNode(e E, parent *node) *node {
    return &node{
        e:      e,
        parent: parent,
    }
}

type Bstree struct {
    size       int
    root       *node
    comparator Comparator
}

func NewBstree() *Bstree {
    return &Bstree{}
}

func NewBstreeWithComparator(comparator Comparator) *Bstree {
    return &Bstree{
        comparator: comparator,
    }
}

func (b *Bstree) Size() int {
    return b.size
}

func (b *Bstree) IsEmpty() bool {
    return b.size == 0
}

func (b *Bstree) Clear() {
    //TODO implement me
    panic("implement me")
}

func (b *Bstree) Add(e E) {
    b.elementNotNullCheck(e)
    
    if b.root == nil { // 添加第一个节点
        b.root = newNode(e, nil)
        b.size++
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    parent := b.root // 保存添加节点的父节点
    n := b.root
    cmp := 0
    for n != nil {
        cmp = b.compare(e, n.e)
        parent = n
        if cmp > 0 {
            n = n.right
        } else if cmp < 0 {
            n = n.left
        } else {
            n.e = e
            return
        }
    }
    if cmp > 0 {
        parent.right = newNode(e, parent)
    } else {
        parent.left = newNode(e, parent)
    }
    b.size++
}

func (b *Bstree) Remove(e E) {
    //TODO implement me
    panic("implement me")
}

func (b *Bstree) Contains(e E) bool {
    //TODO implement me
    panic("implement me")
}

// 前序遍历
func (b *Bstree) PreorderTraversal() {
    b.preorderTraversal(b.root)
}

func (b *Bstree) preorderTraversal(n *node) {
    if n == nil {
        return
    }
    fmt.Println(n.e)
    b.preorderTraversal(n.left)
    b.preorderTraversal(n.right)
}

// 中序遍历
func (b *Bstree) InorderTraversal() {
    b.inorderTraversal(b.root)
}

func (b *Bstree) inorderTraversal(n *node) {
    if n == nil {
        return
    }
    b.inorderTraversal(n.left)
    fmt.Println(n.e)
    b.inorderTraversal(n.right)
}

// 后序遍历
func (b *Bstree) PostorderTraversal() {
    b.postorderTraversal(b.root)
}

func (b *Bstree) postorderTraversal(n *node) {
    if n == nil {
        return
    }
    b.postorderTraversal(n.left)
    b.postorderTraversal(n.right)
    fmt.Println(n.e)
}

// 层序遍历
func (b *Bstree) LevelOrderTraversal() {
    if b.root == nil {
        return
    }
    queue := make([]*node, 0)
    queue = append(queue, b.root)
    for len(queue) != 0 {
    
    }
}

func (b *Bstree) levelOrderTraversal() {}

func (b *Bstree) compare(e1, e2 E) int {
    if b.comparator != nil {
        return b.comparator.Compare(e1, e2)
    }
    return e1.CompareTo(e2)
}

func (b *Bstree) elementNotNullCheck(e E) {
    if e == nil {
        panic("element must not be null")
    }
}
