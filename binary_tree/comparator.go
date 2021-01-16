package binary_tree

// 元素可比较接口
type Comparable interface {
    compareTo(e Comparable) int
}

// 比较器
type Comparator interface {
    // 等于0: e1==e2; 大于0: e > e2; 小于0: e1 < e2
    compare(e1 interface{}, e2 interface{}) int
}