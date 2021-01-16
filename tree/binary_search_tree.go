package tree

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
    left    *node       // 左子节点
    right   *node       // 右子节点
    parent  *node       // 父节点
}

func newNode(element interface{}, parent *node) *node {
    return &node{
        element: element,
        parent:  parent,
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

type BinarySearchTree struct {
    size       int        // 二叉树中节点个数
    root       *node      // 根节点
    comparator Comparator // 接收外部传入的实现了比较器的结构类型
}

// 返回一个二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
    return &BinarySearchTree{
        size: 0,
        root: nil,
    }
}

// 返回一个自定义比较器的二叉搜索树
func NewBinarySearchTreeWithComparator(comparator Comparator) *BinarySearchTree {
    return &BinarySearchTree{
        size:       0,
        root:       nil,
        comparator: comparator,
    }
}

// 获取二叉树中节点个数
func (b *BinarySearchTree) GetSize() int {
    return b.size
}

// 判断二叉树中节点数是否为0
func (b *BinarySearchTree) IsEmpty() bool {
    return b.size == 0
}

// 清空二叉树
func (b *BinarySearchTree) Clear() {
    b.root = nil
    b.size = 0
}

// 添加元素
func (b *BinarySearchTree) Add(element interface{}) {
    b.elementNotNull(element) // 判断传入的元素是否为空，为空就panic

    // 如果根节点为空，添加根节点
    if b.root == nil {
        b.root = &node{element: element}
        b.size++
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    tmpNode := b.root // 起始tmpNode指向根节点
    var parent *node  // 最终指向待添加元素的父节点
    cmp := 0          // 存储两个元素的比较结果
    for tmpNode != nil {
        cmp = b.compare(element, tmpNode.element)
        parent = tmpNode
        // 如果传入元素大，则tmpNode指向右子节点
        // 如果传入元素小，则tmpNode指向左子节点
        // 如果一样大，则新元素覆盖旧元素
        if cmp > 0 {
            tmpNode = tmpNode.right
        } else if cmp < 0 {
            tmpNode = tmpNode.left
        } else {
            tmpNode.element = element
            return
        }
    }
    // 新建一个节点，并插入该节点
    newNode := newNode(element, parent)
    if cmp > 0 {
        parent.right = newNode
    } else {
        parent.left = newNode
    }
    b.size++
}

// 删除节点
func (b *BinarySearchTree) Remove(element interface{}) {
    // 如果节点的度为2
    b.remove(b.getNodeByElement(element))
}

func (b *BinarySearchTree) remove(n *node) {
    if n == nil {
        return
    }
    b.size--
    if n.hasTwoChildren() { // 度为2的节点
        // 找到后继节点
        s := b.getSuccessor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.element = s.element
        // 删除后继节点
        n = s
    }
    // 删除n节点，n的度必然为1或者0
    var replacement *node
    if n.left != nil {
        replacement = n.left
    } else if n.right != nil {
        replacement = n.right
    } else {
        replacement = nil
    }

    if replacement != nil { // n是度为1的节点
        replacement.parent = n.parent
        if n.parent == nil { // n是度为1的节点并且是根节点
            b.root = replacement
        } else if n == n.parent.left{
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
    } else if n.parent == nil{ // n是叶子节点并且是根节点
        b.root = nil
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
    }
}

// 根据元素找到节点
func (b *BinarySearchTree) getNodeByElement(element interface{}) *node {
    if element == nil {
        return nil
    }
    n := b.root
    for n != nil {
        cmp := b.compare(element, n.element)
        if cmp == 0 {
            return n
        }
        if cmp > 0 {
            n = n.right
        } else {
            n = n.left
        }
    }
    return nil
}

// 是否包含节点
func (b *BinarySearchTree) Contains(element interface{}) bool {
    return b.getNodeByElement(element) != nil
}

// 判断传入的元素是否为空，为空则panic
func (b *BinarySearchTree) elementNotNull(element interface{}) {
    if element == nil {
        panic("element must not be nil!")
    }
}

/**
 * @return 等于0: e1==e2; 大于0: e > e2; 小于0: e1 < e2
 */
func (b *BinarySearchTree) compare(e1 interface{}, e2 interface{}) int {
    // 如果比较器非空，表示外部有传入实现好的比较器
    if b.comparator != nil {
        return b.comparator.compare(e1, e2)
    }
    // 如果没有传入比较器，默认元素本身实现了可比较的接口
    return e1.(Comparable).compareTo(e2.(Comparable))
}

// 前序遍历
func (b *BinarySearchTree) PreOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    b.preOrderTraversal(b.root, visitor)
}

func (b *BinarySearchTree) preOrderTraversal(n *node, visitor Visitor) {
    if n == nil || visitor.stop() {
        return
    }
    visitor.visit(n.element)
    b.preOrderTraversal(n.left, visitor)
    b.preOrderTraversal(n.right, visitor)
}

// 中序遍历
func (b *BinarySearchTree) InnerOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    b.innerOrderTraversal(b.root, visitor)
}

func (b *BinarySearchTree) innerOrderTraversal(n *node, visitor Visitor) {
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
func (b *BinarySearchTree) PostOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    b.postOrderTraversal(b.root, visitor)
}
func (b *BinarySearchTree) postOrderTraversal(n *node, visitor Visitor) {
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
func (b *BinarySearchTree) LevelOrder(vistor Visitor) {
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
    }
}

// 获取树的高度, 也就是获取根节点的高度
func (b *BinarySearchTree) GetTreeHeight() int {
    return b.getNodeHeightByIteration(b.root)
}

// 通过递归方式, 获取某个节点的高度
func (b *BinarySearchTree) getNodeHeightByRecursion(n *node) int {
    if n == nil {
        return 0
    }
    return int(math.Max(float64(b.getNodeHeightByRecursion(n.left)), float64(b.getNodeHeightByRecursion(n.right)))) + 1
}

// 通过层序遍历的方式, 获取某个节点的高度
func (b *BinarySearchTree) getNodeHeightByIteration(n *node) int {
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
func (b *BinarySearchTree) IsComplete() bool {
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
func (b *BinarySearchTree) getPredecessor(n *node) *node {
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
func (b *BinarySearchTree) getSuccessor(n *node) *node {
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
