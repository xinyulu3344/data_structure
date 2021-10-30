package treemap

type Node struct {
    key interface{}
    value interface{}
    color bool
    left *Node
    right *Node
    parent *Node
}

// 判断节点是否是叶子节点
func (r *Node) isLeaf() bool {
    return r.left == nil && r.right == nil
}

// 判断节点是否有两个子节点
func (r *Node) hasTwoChildren() bool {
    return r.left != nil && r.right != nil
}

// 判断节点是否是父节点的左节点
func (r *Node) isLeftChild() bool {
    return r.parent != nil && r == r.parent.left
}

// 判断节点是否是父节点的右节点
func (r *Node) isRightChild() bool {
    return r.parent != nil && r == r.parent.right
}

// 返回当前节点的兄弟节点
func (r *Node) sibling() *Node {
    if r.isLeftChild() {
        return r.parent.right
    }
    if r.isRightChild() {
        return r.parent.left
    }
    return nil
}

type TreeMap struct {
}

func (t TreeMap) GetSize() int {
    panic("implement me")
}

func (t TreeMap) IsEmpty() bool {
    panic("implement me")
}

func (t TreeMap) Clear() {
    panic("implement me")
}

func (t TreeMap) Put(k, v interface{}) interface{} {
    panic("implement me")
}

func (t TreeMap) Get(k interface{}) interface{} {
    panic("implement me")
}

func (t TreeMap) Remove(k interface{}) interface{} {
    panic("implement me")
}

func (t TreeMap) ContainsKey(k interface{}) bool {
    panic("implement me")
}

func (t TreeMap) ContainsValue(v interface{}) bool {
    panic("implement me")
}

func (t TreeMap) Traversal(visitor Visitor) bool {
    panic("implement me")
}


