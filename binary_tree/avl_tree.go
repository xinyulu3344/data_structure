/**
  @author: xinyulu
  @date: 2020/12/8 23:46
  @note: 
**/
package binary_tree

import (
    "math"
)

type AVLTree struct {
    *BinarySearchTree
}

// 添加元素
func (a *AVLTree) Add(element interface{}) {
    a.elementNotNull(element) // 判断传入的元素是否为空，为空就panic

    // 如果根节点为空，添加根节点
    if a.root == nil {
        a.root = &node{element: element}
        a.size++
        // 添加根元素之后
        a.afterAdd(a.root)
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    tmpNode := a.root // 起始tmpNode指向根节点
    var parent *node  // 最终指向待添加元素的父节点
    cmp := 0          // 存储两个元素的比较结果
    for tmpNode != nil {
        cmp = a.compare(element, tmpNode.element)
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
    a.size++
    a.afterAdd(newNode)
}

// 添加节点后的操作
// n: 新增的节点
func (a *AVLTree) afterAdd(n *node) {
    for n.parent != nil {
        n = n.parent
       if a.isBalanced(n) {
           // 更新高度
           a.updateHeight(n)
       } else {
           // 恢复平衡
           a.rebalance(n)
           break
       }
    }
}

// 删除节点
func (a *AVLTree) Remove(element interface{}) {
    // 如果节点的度为2
    a.remove(a.getNodeByElement(element))
}

func (a *AVLTree) remove(n *node) {
    if n == nil {
        return
    }
    a.size--
    if n.hasTwoChildren() { // 度为2的节点
        // 找到后继节点
        s := a.getSuccessor(n)
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
            a.root = replacement
        } else if n == n.parent.left{
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
        // 删除节点之后的处理
        a.afterRemove(n)
    } else if n.parent == nil{ // n是叶子节点并且是根节点
        a.root = nil
        a.afterAdd(n)
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
        a.afterRemove(n)
    }
}

func (a *AVLTree) afterRemove(n *node) {
    for n.parent != nil {
        n = n.parent
        if a.isBalanced(n) {
            // 更新高度
            a.updateHeight(n)
        } else {
            // 恢复平衡
            a.rebalance(n)
        }
    }
}

func (a *AVLTree) isBalanced(n *node) bool {
    return math.Abs(float64(a.balanceFactor(n))) <= 1
}

// 恢复平衡
// grand: 高度最低的那个不平衡节点
func (a *AVLTree) rebalance(grand *node) {
    // 找到新增节点的父节点
    parent := a.tallerChild(grand)
    // 找到导致失衡的新增节点
    n := a.tallerChild(parent)
    if parent.isLeftChild() {
        if n.isLeftChild() { // LL
            a.rotateRight(grand)
        } else { // LR
            a.rotateLeft(parent)
            a.rotateRight(grand)
        }
    } else { // parent在grand的右边
        if n.isLeftChild() { // RL
            a.rotateRight(parent)
            a.rotateLeft(grand)
        } else { // RR
            a.rotateLeft(grand)
        }
    }
}

// 左旋转
func (a *AVLTree) rotateLeft(grand *node) {
    parent := grand.right
    child := parent.left
    grand.right = parent.left
    parent.left = grand

    a.afterRotate(grand, parent, child)
}

// 右旋转
func (a *AVLTree) rotateRight(grand *node) {
    parent := grand.left
    child := parent.right
    grand.left = child
    parent.right = grand

    a.afterRotate(grand, parent, child)
}

// 旋转之后的维护操作
func (a *AVLTree) afterRotate(grand, parent, child *node) {
    // 让parent成为子树的根节点
    parent.parent = grand.parent
    if grand.isLeftChild() {
        grand.parent.left = parent
    } else if grand.isRightChild() {
        grand.parent.right = parent
    } else { // grand是root节点
        a.root = parent
    }

    // 更新child的parent
    if child != nil {
        child.parent = grand
    }

    // 更新grand的parent
    grand.parent = parent

    // 更新高度
    a.updateHeight(grand)
    a.updateHeight(parent)
}

// 返回高度高的子树
func (a *AVLTree) tallerChild(n *node) *node {
    leftHeight := a.getLeftHeight(n)
    rightHeight := a.getRightHeight(n)
    if leftHeight > rightHeight {
        return n.left
    }else if leftHeight < rightHeight {
        return n.right
    } else {
        if n.isLeftChild() {
            return n.left
        } else {
            return n.right
        }
    }
}

// 更新子树高度
func (a *AVLTree) updateHeight(n *node) {
    n.height = 1 + int(math.Max(float64(a.getLeftHeight(n)), float64(a.getRightHeight(n))))
}

// 计算平衡因子
func (a *AVLTree) balanceFactor(n *node) int {
    return a.getLeftHeight(n) - a.getRightHeight(n)
}

// 获取左子树高度
func (a *AVLTree) getLeftHeight(n *node) (leftHeight int) {
    if n.left == nil {
        return 0
    } else {
        return n.left.height
    }
}

// 获取右子树高度
func (a *AVLTree) getRightHeight(n *node) (rightHeight int) {
    if n.right == nil {
        return 0
    } else {
        return n.right.height
    }
}

