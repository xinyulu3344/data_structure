package bstree

import (
	"fmt"
)

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
    comparator Compare
}

func NewBstree() *Bstree {
    return &Bstree{}
}

func NewBstreeWithComparator(comparator Compare) *Bstree {
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
func (b *Bstree) PreorderTraversal(visit Visit) {
	if visit == nil {
		return
	}
    b.preorderTraversal(b.root, visit)
}

func (b *Bstree) preorderTraversal(n *node, visit Visit) {
    if n == nil {
        return
    }
    fmt.Println(n.e)
    b.preorderTraversal(n.left, visit)
    b.preorderTraversal(n.right, visit)
}

// 中序遍历
func (b *Bstree) InorderTraversal(visit Visit) {
	if visit == nil {
		return
	}
    b.inorderTraversal(b.root, visit)
}

func (b *Bstree) inorderTraversal(n *node, visit Visit) {
    if n == nil {
        return
    }
    b.inorderTraversal(n.left, visit)
    fmt.Println(n.e)
    b.inorderTraversal(n.right, visit)
}

// 后序遍历
func (b *Bstree) PostorderTraversal(visit Visit) {
	if visit == nil {
		return
	}
    b.postorderTraversal(b.root, visit)
}

func (b *Bstree) postorderTraversal(n *node, visit Visit) {
    if n == nil {
        return
    }
    b.postorderTraversal(n.left, visit)
    b.postorderTraversal(n.right, visit)
    fmt.Println(n.e)
}

// 层序遍历
func (b *Bstree) LevelOrderTraversal(visit Visit) {
    if b.root == nil || visit == nil {
        return
    }
    queue := make([]*node, 0)
    queue = append(queue, b.root)
    for len(queue) != 0 {
		// 出队
		n := queue[0]
		queue = queue[1:]
		visit(n.e)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
    }
}


func (b *Bstree) compare(e1, e2 E) int {
    if b.comparator != nil {
        return b.comparator(e1, e2)
    }
    return e1.CompareTo(e2)
}

func (b *Bstree) elementNotNullCheck(e E) {
    if e == nil {
        panic("element must not be null")
    }
}
