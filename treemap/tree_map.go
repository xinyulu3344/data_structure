package treemap

const red bool = false
const black bool = true

type rbNode struct {
    key Key
    value any
    left   *rbNode
    right  *rbNode
    parent *rbNode
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

// 节点是否有两个子节点
func (r *rbNode) hasTwoChildren() bool {
    return r.left != nil && r.right != nil
}

func newRbNode(key Key, value any, parent *rbNode) *rbNode {
    return &rbNode{
        key: key,
        value: value,
        parent: parent,
        color: red,
    }
}


type TreeMap struct {
    size       int
    root       *rbNode
    comparator Compare
}

func NewTreeMap() *TreeMap {
    return &TreeMap{}
}

func NewTreeMapWithComparator(comparator Compare) *TreeMap {
    return &TreeMap{
        comparator: comparator,
    }
}

func (tm *TreeMap) Size() int {
    return tm.size
}

func (tm *TreeMap) IsEmpty() bool {
    return tm.size == 0
}

func (tm *TreeMap) Clear() {
    tm.root = nil
    tm.size = 0
}

func (tm *TreeMap) Put(key Key, value any) {
    tm.elementNotNullCheck(key)
    
    if tm.root == nil { // 添加第一个节点
        tm.root = newRbNode(key, value, nil)
        tm.size++
        
        // 新添加节点之后的处理
        tm.afterPut(tm.root)
        return
    }
    // 添加的不是第一个节点
    // 找到父节点
    parent := tm.root // 保存添加节点的父节点
    n := tm.root
    cmp := 0
    for n != nil {
        cmp = tm.compare(key, n.key)
        parent = n
        if cmp > 0 {
            n = n.right
        } else if cmp < 0 {
            n = n.left
        } else {
            n.key = key
            n.value = value
            return
        }
    }
	newNode := newRbNode(key, value, parent)
    if cmp > 0 {
        parent.right = newNode
    } else {
        parent.left = newNode
    }
    tm.size++
    
    // 新添加节点之后的处理
    tm.afterPut(newNode)
}

func (tm *TreeMap) afterPut(n *rbNode) {
    parent := n.parent

    if parent == nil { // 添加的是根节点或者上溢到达了根节点
        tm.dyeBlack(n) // 将根节点染黑
        return
    }

    if tm.isBlack(parent) { // 父节点是Black，直接返回
        return
    }

    uncle := parent.sibling() // 叔父节点
    grand := parent.parent // 祖父节点
    if tm.isRed(uncle) { // 叔父节点是Red
        tm.dyeBlack(parent)
        tm.dyeBlack(uncle)
        // 把祖父节点当做是新添加的节点
        tm.dyeRed(grand)
        tm.afterPut(grand)
        return
    }

    // 叔父节点不是红色
    if parent.isLeftChild() {
        if n.isLeftChild() { // LL
            tm.dyeBlack(parent)
            tm.dyeRed(grand)
            tm.rotateRight(grand)
        } else { // LR
            tm.dyeBlack(n)
            tm.dyeRed(grand)
            tm.rotateLeft(parent)
            tm.rotateRight(grand)
        }
    } else {                 // parent在grand的右边
        if n.isLeftChild() { // RL
            tm.dyeBlack(n)
            tm.dyeRed(grand)
            tm.rotateRight(parent)
            tm.rotateLeft(grand)
        } else { // RR
            tm.dyeBlack(parent)
            tm.dyeRed(grand)
            tm.rotateLeft(grand)
        }
    }
}

func (tm *TreeMap) Get(key Key) any {
    n := tm.getNodeByElement(key)
    if n != nil {
        return n.value
    }
    return nil
}

// Remove 删除
func (tm *TreeMap) Remove(key Key) any {
    return tm.remove(tm.getNodeByElement(key))
}


// 删除节点
func (tm *TreeMap) remove(n *rbNode) any {
    if n == nil {
        return nil
    }
    tm.size--

    oldValue := n.value

    // 删除度为2的节点
    if n.hasTwoChildren() {
        // 找到待删除节点的后继节点
        s := tm.successor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.key = s.key
        n.value = s.value
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
            tm.root = replacement
        } else if n == n.parent.left {
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
		tm.afterRemove(replacement)
    } else if n.parent == nil { // n是叶子节点并且是根节点
        tm.root = nil
		tm.afterRemove(n)
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
		tm.afterRemove(n)
    }

    return oldValue
}

func (tm *TreeMap) ContainsKey(key Key) bool {
    return tm.getNodeByElement(key) != nil
}

func (tm *TreeMap) ContainsValue(value any, equals Equals) bool {
    if tm.root == nil {
        return false
    }
    queue := make([]*rbNode, 0)
    queue = append(queue, tm.root)
    for len(queue) != 0 {
        // 出队
        n := queue[0]
        queue = queue[1:]
        if tm.valEquals(value, n.value, equals) {
            return true
        }
        if n.left != nil {
            queue = append(queue, n.left)
        }
        if n.right != nil {
            queue = append(queue, n.right)
        }
    }
    return false
}

// 遍历TreeMap
func (tm *TreeMap) Traversal(visit Visit) {
    if visit == nil {
        return
    }
    stop := false
    tm.traversal(tm.root, visit, &stop)
}

func (tm *TreeMap) traversal(n *rbNode, visit Visit, stop *bool) {
    if n == nil || *stop {
        return
    }
    tm.traversal(n.left, visit, stop)
    if *stop {
        return
    }
    if visit(n.key, n.value) {
        *stop = true
        return
    }
    tm.traversal(n.right, visit, stop)
}


// 比较两个value是否相等
func (tm *TreeMap) valEquals(v1, v2 any, equals Equals) bool {
    return equals(v1, v2)
}

func (tm *TreeMap) elementNotNullCheck(key Key) {
    if key == nil {
        panic("key must not be null")
    }
}

func (tm *TreeMap) compare(k1, k2 Key) int {
    if tm.comparator != nil {
        return tm.comparator(k1, k2)
    }
    return k1.CompareTo(k2)
}


// 给节点染色
func (tm *TreeMap) dyeColor(n *rbNode, color bool) *rbNode {
	if n == nil {
		return n
	}
	n.color = color
	return n
}

// 将节点染成红色
func (tm *TreeMap) dyeRed(n *rbNode) *rbNode {
	return tm.dyeColor(n, red)
}

// 将节点染成黑色
func (tm *TreeMap) dyeBlack(n *rbNode) *rbNode {
	return tm.dyeColor(n, black)
}

// 判断节点是什么颜色
func (tm *TreeMap) colorOf(n *rbNode) bool {
	if n == nil {
        return black
    } else {
        return n.color
    }
}

// 判断节点是否是黑色
func (tm *TreeMap) isBlack(n *rbNode) bool {
    // return r.colorOf(n) == black
    return tm.colorOf(n)
}


// 判断节点是否是红色
func (tm *TreeMap) isRed(n *rbNode) bool {
    // return r.colorOf(n) == red
    return !tm.colorOf(n)
}

// 左旋
func (tm *TreeMap) rotateLeft(grand *rbNode) {
	parent := grand.right
	grand.right = parent.left
	parent.left = grand

	tm.afterRotate(grand, parent, grand.right)
}



// 右旋
func (tm *TreeMap) rotateRight(grand *rbNode) {
	parent := grand.left
	grand.left = parent.right
	parent.right = grand

	tm.afterRotate(grand, parent, grand.left)
}

// 旋转之后的维护操作
func (tm *TreeMap) afterRotate(grand, parent, child *rbNode) {
    // 让parent成为子树的根节点
    parent.parent = grand.parent
    if grand.isLeftChild() {
        grand.parent.left = parent
    } else if grand.isRightChild() {
        grand.parent.right = parent
    } else { // grand是root节点
        tm.root = parent
    }

    // 更新child的parent
    if child != nil {
        child.parent = grand
    }

    // 更新grand的parent
    grand.parent = parent
}

// 根据元素找到节点
func (tm *TreeMap) getNodeByElement(key Key) *rbNode {
    if key == nil {
        return nil
    }
    n := tm.root
    for n != nil {
        cmp := tm.compare(key, n.key)
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

// 获取后继结点
func (tm *TreeMap) successor(n *rbNode) *rbNode {
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
func (tm *TreeMap) afterRemove(n *rbNode) {

    
    if tm.isRed(n) { // 用以取代n的子节点是红色，或者用以取代删除节点的子节点是红色
        tm.dyeBlack(n)
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
        if tm.isRed(sibling) { // 兄弟节点是Red，需要先转换为兄弟节点为Black
            tm.dyeBlack(sibling)
            tm.dyeRed(parent)
            tm.rotateLeft(parent)
            // 更换sibling
            sibling = parent.right
        }

        // 来到这里，sibling必然是Black
        if tm.isBlack(sibling.left) && tm.isBlack(sibling.right) { // 兄弟节点没有Red子节点，父节点要向下和兄弟节点合并
            // 判断父节点颜色
            parentIsBlack := tm.isBlack(parent)
            tm.dyeBlack(parent)
            tm.dyeRed(sibling)
            if parentIsBlack { // 如果父节点是黑色，则父节点下来合并的时候，会导致父节点也下溢。
                // 将父节点再当做被删除节点递归处理
                tm.afterRemove(parent)
            }

        } else { // 兄弟节点至少有1个Red子节点
            if tm.isBlack(sibling.right) { // 兄弟右边是Black，符合LR条件，需要先对兄弟节点右旋转
                tm.rotateRight(sibling)
                sibling = parent.right
            }
            tm.dyeColor(sibling, tm.colorOf(parent))
            tm.dyeBlack(sibling.right)
            tm.dyeBlack(parent)
            tm.rotateLeft(parent)
        }
    } else { // 被删除的节点在右边
        if tm.isRed(sibling) { // 兄弟节点是Red，需要先转换为兄弟节点为Black
            tm.dyeBlack(sibling)
            tm.dyeRed(parent)
            tm.rotateRight(parent)
            // 更换sibling
            sibling = parent.left
        }

        // 来到这里，sibling必然是Black
        if tm.isBlack(sibling.left) && tm.isBlack(sibling.right) { // 兄弟节点没有Red子节点，父节点要向下和兄弟节点合并
            // 判断父节点颜色
            parentIsBlack := tm.isBlack(parent)
            tm.dyeBlack(parent)
            tm.dyeRed(sibling)
            if parentIsBlack { // 如果父节点是黑色，则父节点下来合并的时候，会导致父节点也下溢。
                // 将父节点再当做被删除节点递归处理
                tm.afterRemove(parent)
            }

        } else { // 兄弟节点至少有1个Red子节点
            if tm.isBlack(sibling.left) { // 兄弟左边是Black，符合LR条件，需要先对兄弟节点左旋转
                tm.rotateLeft(sibling)
                sibling = parent.left
            }
            tm.dyeColor(sibling, tm.colorOf(parent))
            tm.dyeBlack(sibling.left)
            tm.dyeBlack(parent)
            tm.rotateRight(parent)
        }
    }
}