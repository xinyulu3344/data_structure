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

type RBTree struct {
    size       int
    root       *avlNode
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


func (r *RBTree) afterAdd(n *rbNode) {}

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
	return r.colorOf(n) == black
}


// 判断节点是否是红色
func (r *RBTree) isRed(n *rbNode) bool {
	return r.colorOf(n) == red
}
