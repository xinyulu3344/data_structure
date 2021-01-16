/**
  @author: xinyulu
  @date: 2020/12/8 22:08
  @note: 
**/
package binary_tree

import (
    "data_structure/queue"
    "math"
)

type Visitor interface {
    visit(e interface{}) // 操作遍历的数据
    stop() bool          // 是否终止遍历
}

type node struct {
    element interface{} // 节点元素
    height  int          // 节点高度
    left    *node       // 左子节点
    right   *node       // 右子节点
    parent  *node       // 父节点
}

func newNode(element interface{}, parent *node) *node {
    return &node{
        element: element,
        parent:  parent,
        height: 1,
    }
}

// 判断节点是否是叶子节点
func (n *node) isLeaf() bool {
    return n.left == nil && n.right == nil
}

// 判断节点是否有两个子节点
func (n *node) hasTwoChildren() bool {
    return n.left != nil && n.right != nil
}

// 判断节点是否是父节点的左节点
func (n *node) isLeftChild() bool {
    return n.parent != nil && n == n.parent.left
}

// 判断节点是否是父节点的右节点
func (n *node) isRightChild() bool {
    return n.parent != nil && n == n.parent.right
}


type BinaryTree struct {
    size int   // 二叉树中节点个数
    root *node // 根节点
}

// 获取二叉树中节点个数
func (b *BinaryTree) GetSize() int {
    return b.size
}

// 判断二叉树中节点数是否为0
func (b *BinaryTree) IsEmpty() bool {
    return b.size == 0
}

// 清空二叉树
func (b *BinaryTree) Clear() {
    b.root = nil
    b.size = 0
}

// 前序遍历
func (b *BinaryTree) PreOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    b.preOrderTraversal(b.root, visitor)
}

func (b *BinaryTree) preOrderTraversal(n *node, visitor Visitor) {
    if n == nil || visitor.stop() {
        return
    }
    visitor.visit(n.element)
    b.preOrderTraversal(n.left, visitor)
    b.preOrderTraversal(n.right, visitor)
}

// 中序遍历
func (b *BinaryTree) InnerOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    b.innerOrderTraversal(b.root, visitor)
}

func (b *BinaryTree) innerOrderTraversal(n *node, visitor Visitor) {
    if n == nil || visitor.stop() {
        return
    }
    b.innerOrderTraversal(n.left, visitor)
    if visitor.stop() {
        return
    }
    visitor.visit(n.element)
    b.innerOrderTraversal(n.right, visitor)
}

// 后序遍历
func (b *BinaryTree) PostOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    b.postOrderTraversal(b.root, visitor)
}
func (b *BinaryTree) postOrderTraversal(n *node, visitor Visitor) {
    if n == nil || visitor.stop() {
        return
    }
    b.postOrderTraversal(n.left, visitor)
    b.postOrderTraversal(n.right, visitor)
    if visitor.stop() {
        return
    }
    visitor.visit(n.element)
}

// 层序遍历: 接收外部传入的对象，调用外部对象的操作元素的方法
func (b *BinaryTree) LevelOrder(vistor Visitor) {
    if b.root == nil || vistor == nil {
        return
    }
    q := queue.NewQueue()
    q.EnQueue(b.root)
    for !q.IsEmpty() {
        if vistor.stop() {
            return
        }
        n, _ := q.DeQueue().(*node)
        vistor.visit(n.element)
        if n.left != nil {
            q.EnQueue(n.left)
        }
        if n.right != nil {
            q.EnQueue(n.right)
        }
        //fmt.Println(n.element, n.height)
    }
}

// 获取树的高度, 也就是获取根节点的高度
func (b *BinaryTree) GetTreeHeight() int {
    return b.getNodeHeightByIteration(b.root)
}

// 通过递归方式, 获取某个节点的高度
func (b *BinaryTree) getNodeHeightByRecursion(n *node) int {
    if n == nil {
        return 0
    }
    return int(math.Max(float64(b.getNodeHeightByRecursion(n.left)), float64(b.getNodeHeightByRecursion(n.right)))) + 1
}

// 通过层序遍历的方式, 获取某个节点的高度
func (b *BinaryTree) getNodeHeightByIteration(n *node) int {
    if b.root == nil {
        return 0
    }
    // 树的高度
    height := 0
    // 存储每一层的节点数量
    levelSize := 1
    q := queue.NewQueue()
    q.EnQueue(b.root)
    for !q.IsEmpty() {
        n, _ := q.DeQueue().(*node)
        levelSize-- // 每出队一个节点，这一层未遍历的节点数减1
        if n.left != nil {
            q.EnQueue(n.left)
        }
        if n.right != nil {
            q.EnQueue(n.right)
        }
        if levelSize == 0 { // 如果levelSize减到0，意味着将要访问下一层
            levelSize = q.GetSize() // 此时队列中的元素个数，就是下一层要遍历的节点数
            height++                // 进入下一层，高度加1
        }
    }
    return height
}

//判断是否是完全二叉树
func (b *BinaryTree) IsComplete() bool {
    if b.root == nil {
        return false
    }
    q := queue.NewQueue()
    q.EnQueue(b.root)

    // 只要leaf被置为true，表示后面遍历的所有节点，都必须是叶子节点
    leaf := false
    for !q.IsEmpty() {
        n, _ := q.DeQueue().(*node)
        if leaf && !n.isLeaf() { // 如果该节点应该是叶子节点，但是发现它不是叶子节点，说明这棵树不是完全二叉树
            return false
        }
        if n.left != nil { // 如果左子节点非空，左子节点入队
            q.EnQueue(n.left)
        } else if n.right != nil { // 如果左子节点为空，右子节点非空，判断为非完全二叉树
            return false
        }
        if n.right != nil { // 如果右子节点非空，右子节点入队
            q.EnQueue(n.right)
        } else { // 右子节点为空，则后续遍历的所有节点，都必须是叶子节点
            leaf = true
        }
    }
    return true
}

// 获取前驱结点
func (b *BinaryTree) getPredecessor(n *node) *node {
    if n == nil {
        return nil
    }
    p := n.left
    // 如果左子树不为空, 遍历左子树的右子节点，找出最右子节点
    if p != nil {
        for p.right != nil {
            p = p.right
        }
        return p
    }
    // 从父节点、祖父节点...中寻找前驱结点
    // 直到当前节点的父节点为空或者当前节点是父节点左子节点，返回当前节点的父节点
    for n.parent != nil && n == n.parent.left {
        n = n.parent
    }
    // 到这里，要么n是根节点，父节点为空，要么n是其父节点的左子节点
    // n.parent == nil || n == n.parent.right
    return n.parent
}

// 获取后继结点
func (b *BinaryTree) getSuccessor(n *node) *node {
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

