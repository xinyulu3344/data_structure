package hashmap

const red bool = false
const black bool = true
const DEFAULT_CAPACITY = 1 << 4

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

type HashMap struct {
	size int
	table []*rbNode
}

func NewHashMap() *HashMap {
	return &HashMap{
		table: make([]*rbNode, DEFAULT_CAPACITY),
	}
}

func (hm *HashMap) Size() int {
	return hm.size
}

func (hm *HashMap) IsEmpty() bool {
	return hm.size == 0
}

// 清空
func (hm *HashMap) Clear() {
	if hm.size == 0 {
		return
	}
	hm.size = 0
	length := len(hm.table)
	for i := 0; i < length; i++ {
		hm.table[i] = nil
	}
}

func (hm *HashMap) Put(key Key, value any) {
	idx := hm.index(key)
	root := hm.table[idx]
	if root == nil {
		root = newRbNode(key, value, nil)
		hm.table[idx] = root
		hm.size++
		hm.afterPut(root)
		return
	}
	// root != nil，hash冲突
	// 添加新的节点到红黑树上
	// 添加的不是第一个节点
    // 找到父节点
    parent := root // 保存添加节点的父节点
    n := root
    cmp := 0
    for n != nil {
        cmp = hm.compare(key, n.key)
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
    hm.size++
    
    // 新添加节点之后的处理
    hm.afterPut(newNode)
}

func (hm *HashMap) Get(key Key) any {
	panic("not implemented") // TODO: Implement
}

func (hm *HashMap) Remove(key Key) any {
	panic("not implemented") // TODO: Implement
}

func (hm *HashMap) ContainsKey(key Key) bool {
	panic("not implemented") // TODO: Implement
}

func (hm *HashMap) ContainsValue(value any) bool {
	panic("not implemented") // TODO: Implement
}

func (hm *HashMap) Traversal(v Visit) {
	panic("not implemented") // TODO: Implement
}

// 根据Key生成对应的索引
func (hm *HashMap) index(key Key) int {
	if key == nil {
		return 0
	}
	hash := key.HashCode()
	return hash ^ (hash >> 16) & (len(hm.table) - 1)
}

func (hm *HashMap) afterPut(n *rbNode) {
    parent := n.parent

    if parent == nil { // 添加的是根节点或者上溢到达了根节点
        hm.dyeBlack(n) // 将根节点染黑
        return
    }

    if hm.isBlack(parent) { // 父节点是Black，直接返回
        return
    }

    uncle := parent.sibling() // 叔父节点
    grand := parent.parent // 祖父节点
    if hm.isRed(uncle) { // 叔父节点是Red
        hm.dyeBlack(parent)
        hm.dyeBlack(uncle)
        // 把祖父节点当做是新添加的节点
        hm.dyeRed(grand)
        hm.afterPut(grand)
        return
    }

    // 叔父节点不是红色
    if parent.isLeftChild() {
        if n.isLeftChild() { // LL
            hm.dyeBlack(parent)
            hm.dyeRed(grand)
            hm.rotateRight(grand)
        } else { // LR
            hm.dyeBlack(n)
            hm.dyeRed(grand)
            hm.rotateLeft(parent)
            hm.rotateRight(grand)
        }
    } else {                 // parent在grand的右边
        if n.isLeftChild() { // RL
            hm.dyeBlack(n)
            hm.dyeRed(grand)
            hm.rotateRight(parent)
            hm.rotateLeft(grand)
        } else { // RR
            hm.dyeBlack(parent)
            hm.dyeRed(grand)
            hm.rotateLeft(grand)
        }
    }
}

// 删除后的调整
// n 被删除的节点，或者用以取代被删除节点的子节点（当被删除节点的度为1）
func (hm *HashMap) afterRemove(n *rbNode) {

    
    if hm.isRed(n) { // 用以取代n的子节点是红色，或者用以取代删除节点的子节点是红色
        hm.dyeBlack(n)
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
        if hm.isRed(sibling) { // 兄弟节点是Red，需要先转换为兄弟节点为Black
            hm.dyeBlack(sibling)
            hm.dyeRed(parent)
            hm.rotateLeft(parent)
            // 更换sibling
            sibling = parent.right
        }

        // 来到这里，sibling必然是Black
        if hm.isBlack(sibling.left) && hm.isBlack(sibling.right) { // 兄弟节点没有Red子节点，父节点要向下和兄弟节点合并
            // 判断父节点颜色
            parentIsBlack := hm.isBlack(parent)
            hm.dyeBlack(parent)
            hm.dyeRed(sibling)
            if parentIsBlack { // 如果父节点是黑色，则父节点下来合并的时候，会导致父节点也下溢。
                // 将父节点再当做被删除节点递归处理
                hm.afterRemove(parent)
            }

        } else { // 兄弟节点至少有1个Red子节点
            if hm.isBlack(sibling.right) { // 兄弟右边是Black，符合LR条件，需要先对兄弟节点右旋转
                hm.rotateRight(sibling)
                sibling = parent.right
            }
            hm.dyeColor(sibling, hm.colorOf(parent))
            hm.dyeBlack(sibling.right)
            hm.dyeBlack(parent)
            hm.rotateLeft(parent)
        }
    } else { // 被删除的节点在右边
        if hm.isRed(sibling) { // 兄弟节点是Red，需要先转换为兄弟节点为Black
            hm.dyeBlack(sibling)
            hm.dyeRed(parent)
            hm.rotateRight(parent)
            // 更换sibling
            sibling = parent.left
        }

        // 来到这里，sibling必然是Black
        if hm.isBlack(sibling.left) && hm.isBlack(sibling.right) { // 兄弟节点没有Red子节点，父节点要向下和兄弟节点合并
            // 判断父节点颜色
            parentIsBlack := hm.isBlack(parent)
            hm.dyeBlack(parent)
            hm.dyeRed(sibling)
            if parentIsBlack { // 如果父节点是黑色，则父节点下来合并的时候，会导致父节点也下溢。
                // 将父节点再当做被删除节点递归处理
                hm.afterRemove(parent)
            }

        } else { // 兄弟节点至少有1个Red子节点
            if hm.isBlack(sibling.left) { // 兄弟左边是Black，符合LR条件，需要先对兄弟节点左旋转
                hm.rotateLeft(sibling)
                sibling = parent.left
            }
            hm.dyeColor(sibling, hm.colorOf(parent))
            hm.dyeBlack(sibling.left)
            hm.dyeBlack(parent)
            hm.rotateRight(parent)
        }
    }
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

func (hm *HashMap) compare(k1, k2 Key) int {
	h1 := 0
	h2 := 0
	if k1 == nil {

	}
}
