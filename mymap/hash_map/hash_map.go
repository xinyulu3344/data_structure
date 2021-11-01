package hash_map

import (
    "reflect"
    "strings"
)

const RED bool = false
const BLACK bool = true
const DEFAULT_CAPACITY = 1 << 4

type Comparable interface {
    CompareTo(k interface{}) int
}

type Node struct {
    key    interface{}
    hash   int // key的hash值
    value  interface{}
    color  bool
    left   *Node
    right  *Node
    parent *Node
}

func NewNode(key interface{}, value interface{}, parent *Node) *Node {
    var hash int
    if key == nil {
        hash = 0
    } else {
        hash = hashCode(key)
    }
    return &Node{
        key:    key,
        hash:   hash,
        value:  value,
        color:  RED,
        parent: parent,
    }
}

func (n *Node) isLeaf() bool {
    return n.left == nil && n.right == nil
}

func (n *Node) hasTwoChildren() bool {
    return n.left != nil && n.right != nil
}

func (n *Node) isLeftChild() bool {
    return n.parent != nil && n == n.parent.left
}

func (n *Node) isRightChild() bool {
    return n.parent != nil && n == n.parent.right
}

func (n *Node) sibling() *Node {
    if n.isLeftChild() {
        return n.parent.right
    }
    if n.isRightChild() {
        return n.parent.left
    }
    return nil
}

type HashMap struct {
    size  int
    table []*Node
}

func NewHashMap() *HashMap {
    return &HashMap{
        table: make([]*Node, DEFAULT_CAPACITY),
    }
}

func (h *HashMap) GetSize() int {
    return h.size
}

func (h *HashMap) IsEmpty() bool {
    return h.size == 0
}

func (h *HashMap) Clear() {
    if h.size == 0 {
        return
    }
    h.size = 0
    for i := 0; i < len(h.table); i++ {
        h.table[i] = nil
    }
}

func (h *HashMap) Put(k interface{}, v interface{}) interface{} {
    idx := h.index(k)
    // 取出index位置红黑树根节点
    root := h.table[idx]
    if root == nil {
        root = NewNode(k, v, nil)
        h.table[idx] = root
        h.size++
        // 修复红黑树性质
        h.afterPut(root)
        return nil
    }
    // 添加新的节点到红黑树上
    parent := root
    node := root
    cmp := 0
    for node != nil {
        cmp = h.compare()
        parent = node
        if cmp > 0 {
            node = node.right
        } else if cmp < 0 {
            node = node.left
        } else {
            node.key = k
            oldValue := node.value
            node.value = v
            return oldValue
        }
    }
    
    // 看看插入到父节点哪个位置
    newNode := NewNode(k, v, parent)
    if cmp > 0 {
        parent.right = newNode
    } else {
        parent.left = newNode
    }
    h.size++
    
    return nil
}

func (h *HashMap) Get(k interface{}) interface{} {
    return nil
}

func (h *HashMap) Remove(k interface{}) interface{} {
    return nil
}

func (h *HashMap) ContainsKey(k interface{}) bool {
    return false
}

func (h *HashMap) ContainsValue(v interface{}) bool {
    return false
}

func (h *HashMap) Traversal(func(k, v interface{})) {

}

// 根据key生成对应的索引，在桶数组中的位置
func (h *HashMap) index(k interface{}) int {
    if k == nil {
        return 0
    }
    hash := hashCode(k)
    hash = hash ^ (hash >> 16)
    return hash & (len(h.table) - 1)
}

func (h *HashMap) afterPut(node *Node) {

}

// 用hash值来比较
func (h *HashMap) compare(k1 interface{}, h1 int, k2 interface{}, h2 int) int {
    // 比较哈希值
    result := h1 - h2
    // 哈希值不等
    if result != 0 { return result }
    
    // 哈希值相等，equals相等
    if equals(k1, k2) { return 0 }
    
    // 哈希值相等，equals()不等
    // 比较类名
    if k1 != nil && k2 != nil {
        k1Type := reflect.TypeOf(k1).Name()
        k2Type := reflect.TypeOf(k2).Name()
        result = strings.Compare(k1Type, k2Type)
        if result != 0 { return result }
        // 同一种类型并且具备可比较性
        if key1, ok := k1.(Comparable); ok {
            return key1.CompareTo(k2)
        }
    }
    
    // 哈希值相等，同一种类型，但是不具备可比较性
    // k1不为nil，k2为nil
    // k1为nil，k2不为nil
    return getAddress(k1) - getAddress(k2)
}

func hashCode(k interface{}) int {
    // 未实现
    return 0
}

func equals(k1 interface{}, k2 interface{}) bool {
    // 未实现
    return false
}

func getAddress(k interface{}) int {
    // 未实现
    return 0
}