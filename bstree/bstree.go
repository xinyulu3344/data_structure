package bstree

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
    b.root = nil
    b.size = 0
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

// Remove 删除元素
func (b *Bstree) Remove(e E) {
    b.remove(b.getNodeByElement(e))
}

// 删除节点
func (b *Bstree) remove(n *node) {
    if n == nil {
        return
    }
    b.size--
    // 删除度为2的节点
    if n.hasTwoChildren() {
        // 找到待删除节点的后继节点
        s := b.successor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.e = s.e
        // 让n指向后继节点，后续删除
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
        } else if n == n.parent.left {
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
    } else if n.parent == nil { // n是叶子节点并且是根节点
        b.root = nil
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
    }
}

func (b *Bstree) Contains(e E) bool {
    return b.getNodeByElement(e) != nil
}

// 前序遍历
func (b *Bstree) PreorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    b.preorderTraversal(b.root, visit, &stop)
}

func (b *Bstree) preorderTraversal(n *node, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
    b.preorderTraversal(n.left, visit, stop)
    b.preorderTraversal(n.right, visit, stop)
}

func (b *Bstree) PreorderTraversalIter(visit Visit) {
	if visit == nil || b.root == nil {
		return
	}
	n := b.root
	stack := make([]*node, 0)
	for {
		if n != nil {
			if visit(n.e) {
				return
			}
			if n.right != nil {
				stack = append(stack, n.right)
			}
			n = n.left
		} else if len(stack) == 0 {
			return
		} else {
			n = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
	}
}
func (b *Bstree) PreorderTraversalIter2(visit Visit) {
	if visit == nil || b.root == nil {
		return
	}
	stack := make([]*node, 0)
    stack = append(stack, b.root)
	for len(stack) != 0 {
        n := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if visit(n.e) {
            return
        }
        if n.right != nil {
            stack = append(stack, n.right)
        }
        if n.left != nil {
            stack = append(stack, n.left)
        }
	}
}

// 中序遍历
func (b *Bstree) InorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    b.inorderTraversal(b.root, visit, &stop)
}

func (b *Bstree) inorderTraversal(n *node, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    b.inorderTraversal(n.left, visit, stop)
    if *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
    b.inorderTraversal(n.right, visit, stop)
}

func (b *Bstree) InorderTraversalIter(visit Visit) {
    if b.root == nil || visit == nil {
        return
    }

    n := b.root
    stack := make([]*node, 0)
    for {
        if n != nil {
            stack = append(stack, n)
            n = n.left
        } else if len(stack) == 0 {
            return
        } else {
            n = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            if visit(n.e) {
                return
            }
            // 让右节点进行中序遍历
            n = n.right
        }
    }
}

// 后序遍历
func (b *Bstree) PostorderTraversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    b.postorderTraversal(b.root, visit, &stop)
}

func (b *Bstree) postorderTraversal(n *node, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    b.postorderTraversal(n.left, visit, stop)
    b.postorderTraversal(n.right, visit, stop)
    if *stop {
        return
    }
    if visit(n.e) {
        *stop = true
        return
    }
}

func (b *Bstree) PostorderTraversalIter(visit Visit) {
    if visit == nil || b.root == nil {
        return
    }
    // 记录上次弹出访问的节点
    var prev *node
    stack := make([]*node, 0)
    stack = append(stack, b.root)
    for len(stack) != 0 {
        top := stack[len(stack) - 1]
        if top.isLeaf() || (prev != nil && prev.parent == top) {
            prev = stack[len(stack)-1]
            stack = stack[:len(stack) - 1]
            if visit(prev.e) {
                return
            }
        } else {
            if top.right != nil {
                stack = append(stack, top.right)
            }
            if top.left != nil {
                stack = append(stack, top.left)
            }
        }
    }
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
func (b *Bstree) successor(n *node) *node {
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
func (b *Bstree) getNodeByElement(e E) *node {
    if e == nil {
        return nil
    }
    n := b.root
    for n != nil {
        cmp := b.compare(e, n.e)
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
