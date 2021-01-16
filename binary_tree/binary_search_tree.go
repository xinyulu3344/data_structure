package binary_tree


type BinarySearchTree struct {
    BinaryTree
    comparator Comparator // 接收外部传入的实现了比较器的结构类型
}

// 返回一个二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
    return &BinarySearchTree{
        BinaryTree{
            size: 0,
            root: nil,
        },
        nil,
    }
}

// 返回一个自定义比较器的二叉搜索树
func NewBinarySearchTreeWithComparator(comparator Comparator) *BinarySearchTree {
    return &BinarySearchTree{
        BinaryTree{
            size:0,
            root:nil,
        },
        comparator,
    }
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
