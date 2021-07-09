package set

import (
    "data_structure/binary_tree/rbtree"
)

type TreeSet struct {
    tree *rbtree.RBTree
}

func NewTreeSet() *TreeSet {
    return &TreeSet{
        tree: rbtree.NewRBTree(),
    }
}

func NewTreeSetWithComparator(comparator rbtree.Comparator) *TreeSet {
    return &TreeSet{
        tree: rbtree.NewRBTreeWithComparator(comparator),
    }
}

// 获取元素个数
func (t *TreeSet) Size() int {
    return t.tree.GetSize()
}

// 判断是否为空
func (t *TreeSet) IsEmpty() bool {
    return t.tree.IsEmpty()
}

// 是否包含某元素
func (t *TreeSet) Contains(e interface{}) bool {
    return t.tree.Contains(e)
}

// 添加元素
func (t *TreeSet) Add(e interface{}) {
    if t.tree.Contains(e) {return}
    t.tree.Add(e)
}

// 移除元素
func (t *TreeSet) Remove(e interface{}) {
    t.tree.Remove(e)
}

type RBTreeVistor struct {
    visitor Visitor
}

func (r *RBTreeVistor) Visit(e interface{}, color bool) {
    r.visitor.Visit(e)
}

func (r *RBTreeVistor) Stop() bool {
    return false
}

// 遍历set
func (t *TreeSet) Traversal(visitor Visitor) {
    t.tree.InnerOrderTraversal(&RBTreeVistor{
        visitor: visitor,
    })
}