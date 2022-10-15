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

func (n *node) isLeaf() bool {
    return n.left == nil && n.right == nil
}

func (n *node) hasTwoChildren() bool {
    return n.left != nil && n.right != nil
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

// IsComplete 利用层序遍历判断是否是完全二叉树
func (b *Bstree) IsComplete() bool {
    if b.root == nil {
        return false
    }
    queue := make([]*node, 0)
    queue = append(queue, b.root)
    // 只要leaf被置为true，表示后面遍历的所有节点，都必须是叶子节点
    leaf := false
    for len(queue) != 0 {
        n := queue[0]
        queue = queue[1:]
        
        if leaf && !n.isLeaf() { // 如果该节点应该是叶子节点，但是发现它不是叶子节点，说明这棵树不是完全二叉树
            return false
        }
        
        if n.left != nil { // 如果左子节点非空，左子节点入队
            queue = append(queue, n.left)
        } else if n.right != nil { // 如果左子节点为空，右子节点非空，判断为非完全二叉树
            return false
        }
        
        if n.right != nil { // 如果右子节点非空，右子节点入队
            queue = append(queue, n.right)
        } else { // 意味着后面所有的节点都必须是叶子节点
            leaf = true
        }
    }
    return true
}

// Height 利用层序遍历计算二叉树高度
func (b *Bstree) Height() int {
    if b.root == nil {
        return 0
    }
    
    // 树的高度
    height := 0
    // 存储每一层的元素数量
    levelSize := 1
    
    queue := make([]*node, 0)
    queue = append(queue, b.root)
    for len(queue) != 0 {
        // 出队
        n := queue[0]
        queue = queue[1:]
        levelSize--
        
        if n.left != nil {
            queue = append(queue, n.left)
        }
        if n.right != nil {
            queue = append(queue, n.right)
        }
        if levelSize == 0 { // 意味着即将要访问下一层
            levelSize = len(queue)
            height++
            
        }
    }
    return height
}

// Height2 递归的方式获取二叉树高度
func (b *Bstree) Height2() int {
    return b.height(b.root)
}

func (b *Bstree) height(n *node) int {
    if n == nil {
        return 0
    }
    return 1 + b.max(b.height(n.left), b.height(n.right))
}

func (b *Bstree) max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
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

// 获取前驱结点
func (b *Bstree) predecessor(n *node) *node {
    if n == nil {
        return nil
    }
    
    // 如果左子树不为空, 遍历左子树的右子节点，找出最右子节点
    p := n.left
    if p != nil {
        for p.right != nil {
            p = p.right
        }
        return p
    }
    
    // 从父节点、祖父节点...中寻找前驱结点
    // 直到当前节点的父节点为空并且当前节点是父节点左子节点，返回当前节点的父节点
    for n.parent != nil && n == n.parent.left {
        n = n.parent
    }
    return n.parent
}

// 获取后继结点
func (b *Bstree) getSuccessor(n *node) *node {
    if n == nil {
        return nil
    }
    p := n.right
    if p != nil {
        for p.left != nil {
            p = p.left
        }
        return p
    }
    for n.parent != nil && n == n.parent.right {
        n = n.parent
    }
    // 到这里，要么n是根节点，父节点为空，要么n是其父节点的右子节点
    // n.parent == nil || n == n.parent.left
    return n.parent
}
