package avl

type IAvl interface {
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

type avlNode struct {
    e      E
    left   *avlNode
    right  *avlNode
    parent *avlNode
    // 节点高度
    height int
}

func newAvlNode(e E, parent *avlNode) *avlNode {
    return &avlNode{
        e:      e,
        parent: parent,
        height: 1,
    }
}

func (a *avlNode) isLeaf() bool {
    return a.left == nil && a.right == nil
}

func (a *avlNode) hasTwoChildren() bool {
    return a.left != nil && a.right != nil
}

// 计算节点的平衡因子
func (a *avlNode) balanceFactor() int {
    return a.getLeftHeight() - a.getRightHeight()
}

func (a *avlNode) updateHeight() {
    a.height = 1 + max(a.getLeftHeight(), a.getRightHeight())
}

func (a *avlNode) getLeftHeight() int {
    var leftHeight int
    if a.left != nil {
        leftHeight = a.left.height
    }
    return leftHeight
}

func (a *avlNode) getRightHeight() int {
    var rightHeight int
    
    if a.right != nil {
        rightHeight = a.right.height
    }
    return rightHeight
}

// 返回高度更高的子节点
func (a *avlNode) tallerChild() *avlNode {
    leftHeight := a.getLeftHeight()
    rightHeight := a.getRightHeight()
    if leftHeight > rightHeight {
        return a.left
    } else if leftHeight < rightHeight {
        return a.right
    } else {
        if a.isLeftChild() {
            return a.left
        } else {
            return a.right
        }
    }
}

func (a *avlNode) isLeftChild() bool {
    return a.parent != nil && a == a.parent.left
}

func (a *avlNode) isRightChild() bool {
    return a.parent != nil && a == a.parent.right
}

type AVLTree struct {
    size       int
    root       *avlNode
    comparator Compare
}

func NewAVLTree() *AVLTree {
    return &AVLTree{}
}

func NewAVLTreeWithComparator(comparator Compare) *AVLTree {
    return &AVLTree{
        comparator: comparator,
    }
}

func (a *AVLTree) Size() int {
    return a.size
}

func (a *AVLTree) IsEmpty() bool {
    return a.size == 0
}

func (a *AVLTree) Clear() {
    a.root = nil
    a.size = 0
}

func (a *AVLTree) Add(e E) {
    a.elementNotNullCheck(e)
    
    if a.root == nil { // 添加第一个节点
        a.root = newAvlNode(e, nil)
        a.size++
        
        // 新添加节点之后的处理
        a.afterAdd(a.root)
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    parent := a.root // 保存添加节点的父节点
    n := a.root
    cmp := 0
    for n != nil {
        cmp = a.compare(e, n.e)
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
	newNode := newAvlNode(e, parent)
    if cmp > 0 {
        parent.right = newNode
    } else {
        parent.left = newNode
    }
    a.size++
    
    // 新添加节点之后的处理
    a.afterAdd(newNode)
}

// Remove 删除元素
func (a *AVLTree) Remove(e E) {
    a.remove(a.getNodeByElement(e))
}

// 删除节点
func (a *AVLTree) remove(n *avlNode) {
    if n == nil {
        return
    }
    a.size--
    // 删除度为2的节点
    if n.hasTwoChildren() {
        // 找到待删除节点的后继节点
        s := a.successor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.e = s.e
        // 让n指向后继节点，后续删除
        n = s
    }
    
    // 删除n节点，n的度必然为1或者0
    var replacement *avlNode
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
        } else if n == n.parent.left {
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
    } else if n.parent == nil { // n是叶子节点并且是根节点
        a.root = nil
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
    }
}

func (a *AVLTree) Contains(e E) bool {
    return a.getNodeByElement(e) != nil
}

// 前序遍历
func (a *AVLTree) PreorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    a.preorderTraversal(a.root, visit, &stop)
}

func (a *AVLTree) preorderTraversal(n *avlNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
    a.preorderTraversal(n.left, visit, stop)
    a.preorderTraversal(n.right, visit, stop)
}

// 中序遍历
func (a *AVLTree) InorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    a.inorderTraversal(a.root, visit, &stop)
}

func (a *AVLTree) inorderTraversal(n *avlNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    a.inorderTraversal(n.left, visit, stop)
    if *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
    a.inorderTraversal(n.right, visit, stop)
}

// 后序遍历
func (a *AVLTree) PostorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    a.postorderTraversal(a.root, visit, &stop)
}

func (a *AVLTree) postorderTraversal(n *avlNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    a.postorderTraversal(n.left, visit, stop)
    a.postorderTraversal(n.right, visit, stop)
    if *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
}

// 层序遍历
func (a *AVLTree) LevelOrderTraversal(visit Visit) {
    if a.root == nil || visit == nil {
        return
    }
    queue := make([]*avlNode, 0)
    queue = append(queue, a.root)
    for len(queue) != 0 {
        // 出队
        n := queue[0]
        queue = queue[1:]
        if visit(n.e) {
            return
        }
        if n.left != nil {
            queue = append(queue, n.left)
        }
        if n.right != nil {
            queue = append(queue, n.right)
        }
    }
}

// IsComplete 利用层序遍历判断是否是完全二叉树
func (a *AVLTree) IsComplete() bool {
    if a.root == nil {
        return false
    }
    queue := make([]*avlNode, 0)
    queue = append(queue, a.root)
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
func (a *AVLTree) Height() int {
    if a.root == nil {
        return 0
    }
    
    // 树的高度
    height := 0
    // 存储每一层的元素数量
    levelSize := 1
    
    queue := make([]*avlNode, 0)
    queue = append(queue, a.root)
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
func (a *AVLTree) Height2() int {
    return a.height(a.root)
}

func (a *AVLTree) height(n *avlNode) int {
    if n == nil {
        return 0
    }
    return 1 + max(a.height(n.left), a.height(n.right))
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func (a *AVLTree) compare(e1, e2 E) int {
    if a.comparator != nil {
        return a.comparator(e1, e2)
    }
    return e1.CompareTo(e2)
}

func (a *AVLTree) elementNotNullCheck(e E) {
    if e == nil {
        panic("element must not be null")
    }
}

// 获取前驱结点
func (a *AVLTree) predecessor(n *avlNode) *avlNode {
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
func (a *AVLTree) successor(n *avlNode) *avlNode {
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

// 根据元素找到节点
func (a *AVLTree) getNodeByElement(e E) *avlNode {
    if e == nil {
        return nil
    }
    n := a.root
    for n != nil {
        cmp := a.compare(e, n.e)
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

// 添加节点后的处理
func (a *AVLTree) afterAdd(n *avlNode) {
    // 向上遍历父节点和祖先节点
	// 如果节点平衡，就更新高度
	// 如果节点不平衡，则恢复平衡，只需要处理高度最低的不平衡节点。
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

// 判断节点是否平衡
func (a *AVLTree) isBalanced(n *avlNode) bool {
    return abs(n.balanceFactor()) <= 1
}

// 计算绝对值
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// 更新节点高度
func (a *AVLTree) updateHeight(n *avlNode) {
    n.updateHeight()
}

// 恢复平衡
// grand: 高度最低的那个不平衡节点
func (a *AVLTree) rebalance(grand *avlNode) {
    // 找到新增节点的父节点
    parent := grand.tallerChild()
    // 找到导致失衡的新增节点
    n := parent.tallerChild()
    
    if parent.isLeftChild() {
        if n.isLeftChild() { // LL
            a.rotateRight(grand)
        } else { // LR
            a.rotateLeft(parent)
            a.rotateRight(grand)
        }
    } else {                 // parent在grand的右边
        if n.isLeftChild() { // RL
            a.rotateRight(parent)
            a.rotateLeft(grand)
        } else { // RR
            a.rotateLeft(grand)
        }
    }
}

// 左旋
func (a *AVLTree) rotateLeft(grand *avlNode) {
	parent := grand.right
	grand.right = parent.left
	parent.left = grand

	a.afterRotate(grand, parent, grand.right)
}



// 右旋
func (a *AVLTree) rotateRight(grand *avlNode) {
	parent := grand.left
	grand.left = parent.right
	parent.right = grand

	a.afterRotate(grand, parent, grand.left)
}

// 旋转之后的维护操作
func (a *AVLTree) afterRotate(grand, parent, child *avlNode) {
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