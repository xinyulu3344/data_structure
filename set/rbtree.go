package set

const red bool = false
const black bool = true

type rbNode struct {
    e      E
    left   *rbNode
    right  *rbNode
    parent *rbNode
    // 节点颜色
	color bool
}

func newRbNode(e E, parent *rbNode) *rbNode {
    return &rbNode{
        e:      e,
        parent: parent,
        color: red,
    }
}

func (r *rbNode) isLeaf() bool {
    return r.left == nil && r.right == nil
}

func (r *rbNode) isLeftChild() bool {
    return r.parent != nil && r == r.parent.left
}

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

// 节点是否有两个子节点
func (r *rbNode) hasTwoChildren() bool {
    return r.left != nil && r.right != nil
}

type RBTree struct {
    size       int
    root       *rbNode
    comparator Compare
}

func NewRBTree() *RBTree {
    return &RBTree{}
}

func NewRBTreeWithComparator(comparator Compare) *RBTree {
    return &RBTree{
        comparator: comparator,
    }
}

func (r *RBTree) Size() int {
    return r.size
}

func (r *RBTree) IsEmpty() bool {
    return r.size == 0
}

func (r *RBTree) Clear() {
    r.root = nil
    r.size = 0
}

func (r *RBTree) Contains(e E) bool {
    return r.getNodeByElement(e) != nil
}

// 前序遍历
func (r *RBTree) PreorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    r.preorderTraversal(r.root, visit, &stop)
}

func (r *RBTree) preorderTraversal(n *rbNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
    r.preorderTraversal(n.left, visit, stop)
    r.preorderTraversal(n.right, visit, stop)
}

// 中序遍历
func (r *RBTree) InorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    r.inorderTraversal(r.root, visit, &stop)
}

func (r *RBTree) inorderTraversal(n *rbNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    r.inorderTraversal(n.left, visit, stop)
    if *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
    r.inorderTraversal(n.right, visit, stop)
}

// 后序遍历
func (r *RBTree) PostorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    r.postorderTraversal(r.root, visit, &stop)
}

func (r *RBTree) postorderTraversal(n *rbNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    r.postorderTraversal(n.left, visit, stop)
    r.postorderTraversal(n.right, visit, stop)
    if *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
}

// 层序遍历
func (r *RBTree) LevelOrderTraversal(visit Visit) {
    if r.root == nil || visit == nil {
        return
    }
    queue := make([]*rbNode, 0)
    queue = append(queue, r.root)
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
func (r *RBTree) IsComplete() bool {
    if r.root == nil {
        return false
    }
    queue := make([]*rbNode, 0)
    queue = append(queue, r.root)
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
func (r *RBTree) Height() int {
    if r.root == nil {
        return 0
    }
    
    // 树的高度
    height := 0
    // 存储每一层的元素数量
    levelSize := 1
    
    queue := make([]*rbNode, 0)
    queue = append(queue, r.root)
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
func (r *RBTree) Height2() int {
    return r.height(r.root)
}

func (r *RBTree) height(n *rbNode) int {
    if n == nil {
        return 0
    }
    return 1 + max(r.height(n.left), r.height(n.right))
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

// Add添加元素
func (r *RBTree) Add(e E) {
    r.elementNotNullCheck(e)
    
    if r.root == nil { // 添加第一个节点
        r.root = newRbNode(e, nil)
        r.size++
        
        // 新添加节点之后的处理
        r.afterAdd(r.root)
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    parent := r.root // 保存添加节点的父节点
    n := r.root
    cmp := 0
    for n != nil {
        cmp = r.compare(e, n.e)
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
	newNode := newRbNode(e, parent)
    if cmp > 0 {
        parent.right = newNode
    } else {
        parent.left = newNode
    }
    r.size++
    
    // 新添加节点之后的处理
    r.afterAdd(newNode)
}


func (r *RBTree) afterAdd(n *rbNode) {
    parent := n.parent

    if parent == nil { // 添加的是根节点或者上溢到达了根节点
        r.dyeBlack(n) // 将根节点染黑
        return
    }

    if r.isBlack(parent) { // 父节点是Black，直接返回
        return
    }

    uncle := parent.sibling() // 叔父节点
    grand := parent.parent // 祖父节点
    if r.isRed(uncle) { // 叔父节点是Red
        r.dyeBlack(parent)
        r.dyeBlack(uncle)
        // 把祖父节点当做是新添加的节点
        r.dyeRed(grand)
        r.afterAdd(grand)
        return
    }

    // 叔父节点不是红色
    if parent.isLeftChild() {
        if n.isLeftChild() { // LL
            r.dyeBlack(parent)
            r.dyeRed(grand)
            r.rotateRight(grand)
        } else { // LR
            r.dyeBlack(n)
            r.dyeRed(grand)
            r.rotateLeft(parent)
            r.rotateRight(grand)
        }
    } else {                 // parent在grand的右边
        if n.isLeftChild() { // RL
            r.dyeBlack(n)
            r.dyeRed(grand)
            r.rotateRight(parent)
            r.rotateLeft(grand)
        } else { // RR
            r.dyeBlack(parent)
            r.dyeRed(grand)
            r.rotateLeft(grand)
        }
    }
}

// 给节点染色
func (r *RBTree) dyeColor(n *rbNode, color bool) *rbNode {
	if n == nil {
		return n
	}
	n.color = color
	return n
}

// 将节点染成红色
func (r *RBTree) dyeRed(n *rbNode) *rbNode {
	return r.dyeColor(n, red)
}

// 将节点染成黑色
func (r *RBTree) dyeBlack(n *rbNode) *rbNode {
	return r.dyeColor(n, black)
}

// 判断节点是什么颜色
func (r *RBTree) colorOf(n *rbNode) bool {
	if n == nil {
        return black
    } else {
        return n.color
    }
}

// 判断节点是否是黑色
func (r *RBTree) isBlack(n *rbNode) bool {
    // return r.colorOf(n) == black
    return r.colorOf(n)
}


// 判断节点是否是红色
func (r *RBTree) isRed(n *rbNode) bool {
    // return r.colorOf(n) == red
    return !r.colorOf(n)
}

func (r *RBTree) compare(e1, e2 E) int {
    if r.comparator != nil {
        return r.comparator(e1, e2)
    }
    return e1.CompareTo(e2)
}

func (r *RBTree) elementNotNullCheck(e E) {
    if e == nil {
        panic("element must not be null")
    }
}

// 左旋
func (r *RBTree) rotateLeft(grand *rbNode) {
	parent := grand.right
	grand.right = parent.left
	parent.left = grand

	r.afterRotate(grand, parent, grand.right)
}



// 右旋
func (r *RBTree) rotateRight(grand *rbNode) {
	parent := grand.left
	grand.left = parent.right
	parent.right = grand

	r.afterRotate(grand, parent, grand.left)
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

// 根据元素找到节点
func (r *RBTree) getNodeByElement(e E) *rbNode {
    if e == nil {
        return nil
    }
    n := r.root
    for n != nil {
        cmp := r.compare(e, n.e)
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

// Remove 删除元素
func (r *RBTree) Remove(e E) {
    r.remove(r.getNodeByElement(e))
}

// 删除节点
func (r *RBTree) remove(n *rbNode) {
    if n == nil {
        return
    }
    r.size--
    // 删除度为2的节点
    if n.hasTwoChildren() {
        // 找到待删除节点的后继节点
        s := r.successor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.e = s.e
        // 让n指向后继节点，后续删除
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
		r.afterRemove(replacement)
    } else if n.parent == nil { // n是叶子节点并且是根节点
        r.root = nil
		r.afterRemove(n)
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
		r.afterRemove(n)
    }
}

// 获取后继结点
func (a *RBTree) successor(n *rbNode) *rbNode {
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

// 删除后的调整
// n 被删除的节点，或者用以取代被删除节点的子节点（当被删除节点的度为1）
func (r *RBTree) afterRemove(n *rbNode) {

    
    if r.isRed(n) { // 用以取代n的子节点是红色，或者用以取代删除节点的子节点是红色
        r.dyeBlack(n)
        return
    }

    // 获取被删除节点的父节点
    parent := n.parent
    // 删除的是黑色叶子节点
    if parent == nil { // 1. 删除的是根节点
        return
    }

    // 判断被删除的n是parent的left还是right
    left := (parent.left == nil) || n.isLeftChild()
    var sibling *rbNode
    if left {
        sibling = parent.right
    } else {
        sibling = parent.left
    }

    if left { // 被删除的节点在左边
        if r.isRed(sibling) { // 兄弟节点是Red，需要先转换为兄弟节点为Black
            r.dyeBlack(sibling)
            r.dyeRed(parent)
            r.rotateLeft(parent)
            // 更换sibling
            sibling = parent.right
        }

        // 来到这里，sibling必然是Black
        if r.isBlack(sibling.left) && r.isBlack(sibling.right) { // 兄弟节点没有Red子节点，父节点要向下和兄弟节点合并
            // 判断父节点颜色
            parentIsBlack := r.isBlack(parent)
            r.dyeBlack(parent)
            r.dyeRed(sibling)
            if parentIsBlack { // 如果父节点是黑色，则父节点下来合并的时候，会导致父节点也下溢。
                // 将父节点再当做被删除节点递归处理
                r.afterRemove(parent)
            }

        } else { // 兄弟节点至少有1个Red子节点
            if r.isBlack(sibling.right) { // 兄弟右边是Black，符合LR条件，需要先对兄弟节点右旋转
                r.rotateRight(sibling)
                sibling = parent.right
            }
            r.dyeColor(sibling, r.colorOf(parent))
            r.dyeBlack(sibling.right)
            r.dyeBlack(parent)
            r.rotateLeft(parent)
        }
    } else { // 被删除的节点在右边
        if r.isRed(sibling) { // 兄弟节点是Red，需要先转换为兄弟节点为Black
            r.dyeBlack(sibling)
            r.dyeRed(parent)
            r.rotateRight(parent)
            // 更换sibling
            sibling = parent.left
        }

        // 来到这里，sibling必然是Black
        if r.isBlack(sibling.left) && r.isBlack(sibling.right) { // 兄弟节点没有Red子节点，父节点要向下和兄弟节点合并
            // 判断父节点颜色
            parentIsBlack := r.isBlack(parent)
            r.dyeBlack(parent)
            r.dyeRed(sibling)
            if parentIsBlack { // 如果父节点是黑色，则父节点下来合并的时候，会导致父节点也下溢。
                // 将父节点再当做被删除节点递归处理
                r.afterRemove(parent)
            }

        } else { // 兄弟节点至少有1个Red子节点
            if r.isBlack(sibling.left) { // 兄弟左边是Black，符合LR条件，需要先对兄弟节点左旋转
                r.rotateLeft(sibling)
                sibling = parent.left
            }
            r.dyeColor(sibling, r.colorOf(parent))
            r.dyeBlack(sibling.left)
            r.dyeBlack(parent)
            r.rotateRight(parent)
        }
    }
}