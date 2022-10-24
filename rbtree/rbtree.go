package rbtree

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