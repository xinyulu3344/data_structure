/**
  @author: xinyulu
  @date: 2021/1/23 21:03
  @note: 红黑树
**/
package rbtree

const red bool = false
const black bool = true

type Visitor interface {
    visit(e interface{}, color bool) // 操作遍历的数据
    stop() bool                      // 是否终止遍历
}

type rbNode struct {
    element interface{}
    left    *rbNode
    right   *rbNode
    parent  *rbNode
    color   bool
}

func newRBNode(element interface{}, parent *rbNode) *rbNode {
    return &rbNode{
        element: element,
        parent:  parent,
        color:   red,
    }
}

// 判断节点是否是叶子节点
func (r *rbNode) isLeaf() bool {
    return r.left == nil && r.right == nil
}

// 判断节点是否有两个子节点
func (r *rbNode) hasTwoChildren() bool {
    return r.left != nil && r.right != nil
}

// 判断节点是否是父节点的左节点
func (r *rbNode) isLeftChild() bool {
    return r.parent != nil && r == r.parent.left
}

// 判断节点是否是父节点的右节点
func (r *rbNode) isRightChild() bool {
    return r.parent != nil && r == r.parent.right
}

// 返回当前节点的兄弟节点
func (r *rbNode) sibling() *rbNode {
    if r.isLeftChild() {
        return r.parent.right
    }
    if r.isRightChild() {
        return r.parent.left
    }
    return nil
}

type RBTree struct {
    size       int
    root       *rbNode
    comparator Comparator
}

// 新建一个红黑树
func NewRBTree() *RBTree {
    return &RBTree{
        size:       0,
        root:       nil,
        comparator: nil,
    }
}

// 新建一个带比较器的红黑树
func NewRBTreeWithComparator(comparator Comparator) *RBTree {
    return &RBTree{
        size:       0,
        root:       nil,
        comparator: comparator,
    }
}

// 给节点染色
func dyeColor(node *rbNode, color bool) *rbNode {
    if node == nil {
        return nil
    }
    node.color = color
    return node
}

// 将节点染成红色
func dyeRed(node *rbNode) *rbNode {
    return dyeColor(node, red)
}

// 将节点染成黑色
func dyeBlack(node *rbNode) *rbNode {
    return dyeColor(node, black)
}

// 判断节点是什么颜色
func colorOf(node *rbNode) bool {
    if node == nil {
        return black
    } else {
        return node.color
    }
}

// 判断节点是否是黑色
func isBlack(node *rbNode) bool {
    return colorOf(node) == black
}

// 判断节点是否是红色
func isRed(node *rbNode) bool {
    return colorOf(node) == red
}

// 添加元素
func (r *RBTree) Add(element interface{}) {
    r.elementNotNull(element) // 判断传入的元素是否为空，为空就panic

    // 如果根节点为空，添加根节点
    if r.root == nil {
        r.root = &rbNode{element: element}
        r.size++
        // 添加根元素之后
        r.afterAdd(r.root)
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    tmpNode := r.root  // 起始tmpNode指向根节点
    var parent *rbNode // 最终指向待添加元素的父节点
    cmp := 0           // 存储两个元素的比较结果
    for tmpNode != nil {
        cmp = r.compare(element, tmpNode.element)
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
    newNode := newRBNode(element, parent)
    if cmp > 0 {
        parent.right = newNode
    } else {
        parent.left = newNode
    }
    r.size++
    r.afterAdd(newNode)
}

// 添加后的操作
func (r *RBTree) afterAdd(node *rbNode) {
    parent := node.parent

    // 如果添加的是根节点或者上溢到达了根节点
    if parent == nil {
        dyeBlack(node)
        return
    }

    // 如果父节点是黑色，直接返回，无需修复
    if isBlack(parent) {
        return
    }

    // 获取叔父节点
    uncle := parent.sibling()
    // 获取祖父节点
    grand := parent.parent

    // 叔父节点是红色
    // LL上溢、RR上溢、LR上溢、RL上溢
    if isRed(uncle) {
        dyeBlack(parent)
        dyeBlack(uncle)
        // 把祖父节点作为新添加节点
        r.afterAdd(dyeRed(grand))
        return
    }

    // 叔父节点不是红色
    // LL、RR、LR、RL
    if parent.isLeftChild() { // L
        dyeRed(grand)
        if node.isLeftChild() { // LL
            dyeBlack(parent)
        } else { // LR
            dyeBlack(node)
            r.rotateLeft(parent)
        }
        r.rotateRight(grand)
    } else { // R
        dyeRed(grand)
        if node.isLeftChild() { // RL
            dyeBlack(node)
            r.rotateRight(parent)
        } else { // RR
            dyeBlack(parent)
        }
        r.rotateLeft(grand)
    }
}

// 左旋转
func (r *RBTree) rotateLeft(grand *rbNode) {
    parent := grand.right
    child := parent.left
    grand.right = parent.left
    parent.left = grand

    r.afterRotate(grand, parent, child)
}

// 右旋转
func (r *RBTree) rotateRight(grand *rbNode) {
    parent := grand.left
    child := parent.right
    grand.left = child
    parent.right = grand

    r.afterRotate(grand, parent, child)
}

// 旋转之后的维护操作
func (r *RBTree) afterRotate(grand, parent, child *rbNode) {
    // 让parent成为子树的根节点
    parent.parent = grand.parent
    if grand.isLeftChild() {
        grand.parent.left = parent
    } else if grand.isRightChild() {
        grand.parent.right = parent
    } else { // grand是root节点
        r.root = parent
    }

    // 更新child的parent
    if child != nil {
        child.parent = grand
    }

    // 更新grand的parent
    grand.parent = parent
}

// 判断传入的元素是否为空，为空则panic
func (r *RBTree) elementNotNull(element interface{}) {
    if element == nil {
        panic("element must not be nil!")
    }
}

/**
 * @return 等于0: e1==e2; 大于0: e > e2; 小于0: e1 < e2
 */
func (r *RBTree) compare(e1 interface{}, e2 interface{}) int {
    // 如果比较器非空，表示外部有传入实现好的比较器
    if r.comparator != nil {
        return r.comparator.compare(e1, e2)
    }
    // 如果没有传入比较器，默认元素本身实现了可比较的接口
    return e1.(Comparable).compareTo(e2.(Comparable))
}

// 删除节点
func (r *RBTree) Remove(element interface{}) {
    // 如果节点的度为2
    r.remove(r.getNodeByElement(element))
}

func (r *RBTree) remove(n *rbNode) {
    if n == nil {
        return
    }
    r.size--
    if n.hasTwoChildren() { // 度为2的节点
        // 找到后继节点
        s := r.getSuccessor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.element = s.element
        // 删除后继节点
        n = s
    }
    // 删除n节点，n的度必然为1或者0
    var replacement *rbNode
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
            r.root = replacement
        } else if n == n.parent.left {
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
        // 删除节点之后的处理
        r.afterRemove(replacement)
    } else if n.parent == nil { // n是叶子节点并且是根节点
        r.root = nil
        r.afterAdd(n)
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
        r.afterRemove(n)
    }
}

// 删除之后的处理
// node: 被删除节点 或者 用以取代被删除节点的子节点(当被删除节点的度为1)
func (r *RBTree) afterRemove(node *rbNode) {

    // 删除的节点是红色 或者 用以替代删除节点的子节点是红色
    if isRed(node) {
        dyeBlack(node)
        return
    }

    // 删除的是根节点
    if node.parent == nil {
        return
    }

    // 删除的是黑色叶子节点
    // 判断被删除的节点是左还是右
    left := node.parent.left == nil || node.isLeftChild()
    var sibling *rbNode
    if left { // 被删除的节点在左边，兄弟在右边
        sibling = node.parent.right
    } else { // 被删除的节点在右边，兄弟在左边
        sibling = node.parent.left
    }

    if left { // 被删除的节点在左边，兄弟在右边
        if isRed(sibling) { // 兄弟节点是红色
            dyeBlack(sibling)
            dyeRed(node.parent)
            r.rotateLeft(node.parent)
            // 更换兄弟
            sibling = node.parent.right
        }

        // 兄弟节点必然是黑色
        if isBlack(sibling.left) && isBlack(sibling.right) {
            // 兄弟节点没有一个红色子节点，父节点要向下和兄弟节点合并
            parentIsBlack := isBlack(node.parent)
            dyeBlack(node.parent)
            dyeRed(sibling)
            if parentIsBlack {
                r.afterRemove(node.parent)
            }

        } else { // 兄弟节点至少有一个红色子节点，向兄弟借
            // 兄弟节点的左边是黑色，兄弟节点要先旋转
            if isBlack(sibling.right) {
                r.rotateRight(sibling)
                sibling = node.parent.right
            }
            dyeColor(sibling, colorOf(node.parent))
            dyeBlack(sibling.right)
            dyeBlack(node.parent)
            r.rotateLeft(node.parent)
        }
    } else { // 被删除的节点在右边，兄弟节点在左边
        if isRed(sibling) { // 兄弟节点是红色
            dyeBlack(sibling)
            dyeRed(node.parent)
            r.rotateRight(node.parent)
            // 更换兄弟
            sibling = node.parent.left
        }

        // 兄弟节点必然是黑色
        if isBlack(sibling.left) && isBlack(sibling.right) {
            // 兄弟节点没有一个红色子节点，父节点要向下和兄弟节点合并
            parentIsBlack := isBlack(node.parent)
            dyeBlack(node.parent)
            dyeRed(sibling)
            if parentIsBlack {
                r.afterRemove(node.parent)
            }

        } else { // 兄弟节点至少有一个红色子节点，向兄弟借
            // 兄弟节点的左边是黑色，兄弟节点要先旋转
            if isBlack(sibling.left) {
                r.rotateLeft(sibling)
                sibling = node.parent.left
            }
            dyeColor(sibling, colorOf(node.parent))
            dyeBlack(sibling.left)
            dyeBlack(node.parent)
            r.rotateRight(node.parent)
        }

    }

}

// 根据元素找到节点
func (r *RBTree) getNodeByElement(element interface{}) *rbNode {
    if element == nil {
        return nil
    }
    n := r.root
    for n != nil {
        cmp := r.compare(element, n.element)
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

// 获取前驱结点
func (r *RBTree) getPredecessor(n *rbNode) *rbNode {
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
func (r *RBTree) getSuccessor(n *rbNode) *rbNode {
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

// 中序遍历
func (r *RBTree) InnerOrderTraversal(visitor Visitor) {
    if visitor == nil {
        return
    }
    r.innerOrderTraversal(r.root, visitor)
}

func (r *RBTree) innerOrderTraversal(n *rbNode, visitor Visitor) {
    if n == nil || visitor.stop() {
        return
    }
    r.innerOrderTraversal(n.left, visitor)
    if visitor.stop() {
        return
    }
    visitor.visit(n.element, n.color)
    r.innerOrderTraversal(n.right, visitor)
}
