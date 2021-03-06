package rbtree

// 元素可比较接口
type Comparable interface {
    CompareTo(e Comparable) int
}

// 比较器
type Comparator interface {
    // 等于0: e1==e2; 大于0: e1 > e2; 小于0: e1 < e2
    Compare(e1 interface{}, e2 interface{}) int
}
