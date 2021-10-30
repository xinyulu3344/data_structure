package mymap

const RED bool = false
const BLACK bool = true
const DEFAULT_CAPACITY = 1 << 4

type Node struct {
	key   interface{}
	value interface{}
	color bool
	left *Node
    right *Node
    parent *Node
}

func NewNode(key interface{}, value interface{}, parent *Node) *Node {
    return &Node{
        key: key,
        value: value,
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
    size int
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
}