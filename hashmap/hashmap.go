package hashmap

import (
	"fmt"
	"strconv"
)

const red bool = false
const black bool = true
const DEFAULT_CAPACITY = 1 << 4
// 默认装填因子
const DEFAULT_LOAD_FACTOR = 0.75

type rbNode struct {
	hash   int
	key    Key
	value  any
	left   *rbNode
	right  *rbNode
	parent *rbNode
	color  bool
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
	n := &rbNode{
		key:    key,
		value:  value,
		parent: parent,
		color:  red,
	}
	var h int
	if key == nil {
		h = 0
	} else {
		h = key.HashCode()
	}
	n.hash = h ^ (h >> 16)
	return n
}

type HashMap struct {
	size  int
	table []*rbNode
    comparator Compare
}

func NewHashMap() *HashMap {
	return &HashMap{
		table: make([]*rbNode, DEFAULT_CAPACITY),
	}
}

func NewHashMapWithComparator(comparator Compare) *HashMap {
    return &HashMap{
		table: make([]*rbNode, DEFAULT_CAPACITY),
        comparator: comparator,
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
	hm.resize()
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
	n := root // 保存遍历红黑树的游标，起始指向root
	cmp := 0 // 保存比较结果
    k1 := key // 将传入的key赋值给k1，增强可读性
	h1 := hm.hash(k1) // 保存传入key的哈希值
    var result *rbNode
	var searched bool // 用于扫描红黑树判断key是否存在时，保证只扫描一次
	for n != nil {
		parent = n
        k2 := n.key
        h2 := n.hash
		tempCmp := 0 // 存储具有可比较性的key的比较的值。避免一个Equals不等的key，但是cmp为0的情况。
		if hm.comparator != nil {
			tempCmp = hm.comparator(k1, k2)
		}
		if h1 > h2 { // 先比较哈希值，哈希值大的，放在红黑树右边
            cmp = 1
        } else if h1 < h2 { // 哈希值小的，放在红黑树左边
            cmp = -1
        } else if k1.Equals(k2) {
            cmp = 0
        } else if k1 != nil && k2 != nil && hm.comparator != nil && tempCmp != 0 { 
			cmp = tempCmp
        } else if searched { // key已经扫描了
			cmp = hm.getMemAdd(&k1) - hm.getMemAdd(&k2)
        } else { // searched == false，key还没有扫描。后面根据内存地址大小决定左右
            if n.left != nil || n.right != nil { // 如果左边不空
                if result = hm.getNode(n.left, k1); result != nil {
                    n = result
                    cmp = 0
                } else if result = hm.getNode(n.right, k1); result != nil {
                    n = result
                    cmp = 0
				}else {
					searched = true
					cmp = hm.getMemAdd(&k1) - hm.getMemAdd(&k2)
                }
            } else {
				searched = true
				cmp = hm.getMemAdd(&k1) - hm.getMemAdd(&k2)
            }
		}

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
    n := hm.getNodeByKey(key)
    if n != nil {
        return n.value
    }
    return nil
}

func (hm *HashMap) Remove(key Key) any {
    return hm.remove(hm.getNodeByKey(key))
}

func (hm *HashMap) ContainsKey(key Key) bool {
    return hm.getNodeByKey(key) != nil
}

func (hm *HashMap) ContainsValue(value any, equals Equals) bool {
    if hm.size == 0 {
        return false
    }
    length := len(hm.table)
    queue := make([]*rbNode, 0)
    for i := 0; i < length; i++ {
        if hm.table[i] == nil {
            continue
        }
        queue = append(queue, hm.table[i])
        for len(queue) != 0 {
            n := queue[0]
            queue = queue[1:]
            if hm.valEquals(value, n.value, equals) {
                return true
            }
            if n.left != nil {
                queue = append(queue, n.left)
            }
            if n.right != nil {
                queue = append(queue, n.right)
            }
        }
    }
    return false
}

// 比较两个value是否相等
func (hm *HashMap) valEquals(v1, v2 any, equals Equals) bool {
    return equals(v1, v2)
}

func (hm *HashMap) Traversal(visit Visit) {
    if hm.size == 0 || visit == nil {
        return
    }
    queue := make([]*rbNode, 0)
    length := len(hm.table)
    for i := 0; i < length; i++ {
        if hm.table[i] == nil {
            continue
        }
        queue = append(queue, hm.table[i])
        for len(queue) != 0 {
            n := queue[0]
            queue = queue[1:]
            if visit(n.key, n.value) {
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
}

// 根据Key生成对应的索引
func (hm *HashMap) index(key Key) int {
	return hm.hash(key) & (len(hm.table)-1)
}

// 获取节点的索引
func (hm *HashMap) indexByNode(n *rbNode) int {
	return n.hash & (len(hm.table)-1)
}

// 扰动计算
func (hm *HashMap) hash(key Key) int {
	if key == nil {
		return 0
	}
	h := key.HashCode()
	return h ^ (h >> 16)
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
	grand := parent.parent    // 祖父节点
	if hm.isRed(uncle) {      // 叔父节点是Red
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
	} else { // parent在grand的右边
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
func (hm *HashMap) dyeColor(n *rbNode, color bool) *rbNode {
	if n == nil {
		return n
	}
	n.color = color
	return n
}

// 将节点染成红色
func (hm *HashMap) dyeRed(n *rbNode) *rbNode {
	return hm.dyeColor(n, red)
}

// 将节点染成黑色
func (hm *HashMap) dyeBlack(n *rbNode) *rbNode {
	return hm.dyeColor(n, black)
}

// 判断节点是什么颜色
func (hm *HashMap) colorOf(n *rbNode) bool {
	if n == nil {
		return black
	} else {
		return n.color
	}
}

// 判断节点是否是黑色
func (hm *HashMap) isBlack(n *rbNode) bool {
	// return r.colorOf(n) == black
	return hm.colorOf(n)
}

// 判断节点是否是红色
func (hm *HashMap) isRed(n *rbNode) bool {
	// return r.colorOf(n) == red
	return !hm.colorOf(n)
}

// 左旋
func (hm *HashMap) rotateLeft(grand *rbNode) {
	parent := grand.right
	grand.right = parent.left
	parent.left = grand

	hm.afterRotate(grand, parent, grand.right)
}

// 右旋
func (hm *HashMap) rotateRight(grand *rbNode) {
	parent := grand.left
	grand.left = parent.right
	parent.right = grand

	hm.afterRotate(grand, parent, grand.left)
}

// 旋转之后的维护操作
func (hm *HashMap) afterRotate(grand, parent, child *rbNode) {
	// 让parent成为子树的根节点
	parent.parent = grand.parent
	if grand.isLeftChild() {
		grand.parent.left = parent
	} else if grand.isRightChild() {
		grand.parent.right = parent
	} else { // grand是root节点
		hm.table[hm.indexByNode(grand)] = parent
	}

	// 更新child的parent
	if child != nil {
		child.parent = grand
	}

	// 更新grand的parent
	grand.parent = parent
}

func (hm *HashMap) getNodeByKey(key Key) *rbNode {
    root := hm.table[hm.index(key)]
    if root == nil {
        return nil
    } else {
        return hm.getNode(root, key)
    }
}

// 根据key查找Node
// 1. 比较哈希值
// 2. key是否Equals
// 3. key是否具备可比较性
// 4. 递归扫描红黑树
func (hm *HashMap) getNode(n *rbNode, k1 Key) *rbNode {
    h1 := hm.hash(k1)
	cmp := 0
    for n != nil {
        k2 := n.key
        h2 := n.hash

		if hm.comparator != nil {
			cmp = hm.comparator(k1, k2)
		}

        // 比较哈希值
        if h1 > h2 {
            n = n.right
        } else if h1 < h2 {
            n = n.left
        } else if k1.Equals(k2) {
            return n
        } else if k1 != nil && k2 != nil && hm.comparator != nil && cmp != 0 {
            if cmp > 0 {
                n = n.right
            } else {
                n = n.left
            }
        } else {
            if n.right != nil {
                if result := hm.getNode(n.right, k1); result != nil {
                    return result
                } else {
					n = n.left
                }
            } else {
				n = n.left
            }
        }
    }
    return nil
}

func (hm *HashMap) remove(n *rbNode) any {
    if n == nil {
        return nil
    }
    hm.size--

    oldValue := n.value

    // 删除度为2的节点
    if n.hasTwoChildren() {
        // 找到待删除节点的后继节点
        s := hm.successor(n)
        // 用后继节点的值覆盖传入的n节点的值
        n.key = s.key
        n.value = s.value
		n.hash = s.hash
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
    
	// 获取红黑树所在的数组索引
    idx := hm.indexByNode(n)
    if replacement != nil { // n是度为1的节点
        replacement.parent = n.parent
        if n.parent == nil { // n是度为1的节点并且是根节点
            hm.table[idx] = replacement
        } else if n == n.parent.left {
            n.parent.left = replacement
        } else {
            n.parent.right = replacement
        }
		hm.afterRemove(replacement)
    } else if n.parent == nil { // n是叶子节点并且是根节点
        hm.table[idx] = nil
		hm.afterRemove(n)
    } else { // n是叶子节点并且不是根节点
        if n == n.parent.left {
            n.parent.left = nil
        } else {
            n.parent.right = nil
        }
		hm.afterRemove(n)
    }

    return oldValue
}

// 获取后继结点
func (hm *HashMap) successor(n *rbNode) *rbNode {
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

func (hm *HashMap) getMemAdd(key *Key) int {
	spk := fmt.Sprintf("%d", key)
	ipk, _ := strconv.Atoi(spk)
	return ipk
}

func (hm *HashMap) resize() {
	if float32(hm.size) / float32(len(hm.table)) <= DEFAULT_LOAD_FACTOR {
		return
	}
	oldTable := hm.table
	hm.table = make([]*rbNode, len(hm.table) << 1)

	length := len(oldTable)
    queue := make([]*rbNode, 0)
    for i := 0; i < length; i++ {
        if oldTable[i] == nil {
            continue
        }
        queue = append(queue, oldTable[i])
        for len(queue) != 0 {
            n := queue[0]
            queue = queue[1:]
            if n.left != nil {
                queue = append(queue, n.left)
            }
            if n.right != nil {
                queue = append(queue, n.right)
            }
			hm.moveNode(n)
        }
    }
}

func (hm *HashMap) moveNode(newNode *rbNode) {
	// 重置
	newNode.parent = nil
	newNode.left = nil
	newNode.right = nil
	newNode.color = red

	idx := hm.indexByNode(newNode)
	root := hm.table[idx]
	if root == nil {
		root = newNode
		hm.table[idx] = root
		hm.afterPut(root)
		return
	}
	// root != nil，hash冲突
	// 添加新的节点到红黑树上
	// 添加的不是第一个节点
	// 找到父节点
	parent := root // 保存添加节点的父节点
	n := root // 保存遍历红黑树的游标，起始指向root
	cmp := 0 // 保存比较结果
    k1 := newNode.key // 将传入的key赋值给k1，增强可读性
	h1 := newNode.hash // 保存传入key的哈希值
	for n != nil {
		parent = n
        k2 := n.key
        h2 := n.hash
		tempCmp := 0 // 存储具有可比较性的key的比较的值。避免一个Equals不等的key，但是cmp为0的情况。
		if hm.comparator != nil {
			tempCmp = hm.comparator(k1, k2)
		}
		if h1 > h2 { // 先比较哈希值，哈希值大的，放在红黑树右边
            cmp = 1
        } else if h1 < h2 { // 哈希值小的，放在红黑树左边
            cmp = -1
        } else if k1 != nil && k2 != nil && hm.comparator != nil && tempCmp != 0 { 
			cmp = tempCmp
        } else {
			cmp = hm.getMemAdd(&k1) - hm.getMemAdd(&k2)
		}

		if cmp > 0 {
			n = n.right
		} else if cmp < 0 {
			n = n.left
		}
	}

	newNode.parent = parent
	if cmp > 0 {
		parent.right = newNode
	} else {
		parent.left = newNode
	}

	// 新添加节点之后的处理
	hm.afterPut(newNode)
}