package stack

import "fmt"

type IStack interface {
    Size() int
    IsEmpty() bool
    Push(e E)
    Pop() E
    Top() E
}

const (
    DEFAULT_CAPACITY = 10
)

type E interface {
    Equal(e E) bool
}

type Stack struct {
    size     int // 元素数量
    elements []E
}

// NewStackWithCap 创建一个动态数组对象
// capacity 动态数组起始容量， 如果小于DEFAULT_CAPACITY(10)，则设置为DEFAULT_CAPACITY
func NewStackWithCap(capacity int) *Stack {
    if capacity < 0 {
        capacity = DEFAULT_CAPACITY
    }
    return &Stack{
        elements: make([]E, capacity),
        size:     0,
    }
}

// NewStack 创建一个容量为DEFAULT_CAPACITY的动态数组
func NewStack() *Stack {
    return NewStackWithCap(DEFAULT_CAPACITY)
}

// Size 获取动态数组的元素数量
func (s *Stack) Size() int {
    return s.size
}

// IsEmpty 判断动态数组是否为空数组
func (s *Stack) IsEmpty() bool {
    return s.size == 0
}

// Push 向动态数组最后位置插入元素
func (s *Stack) Push(e E) {
    s.add(s.size, e)
}

func (s *Stack) Pop() E {
    return s.remove(s.size - 1)
}

func (s *Stack) Top() E {
    return s.get(s.size - 1)
}

// Get 获取指定索引位置的元素
func (s *Stack) get(index int) E {
    s.rangeCheck(index)
    return s.elements[index]
}

// Add 向指定索引位置插入元素
func (s *Stack) add(index int, e E) {
    s.rangeCheckForAdd(index)
    s.ensureCapacity(s.size + 1)
    for i := s.size; i > index; i-- {
        s.elements[i] = s.elements[i-1]
    }
    s.elements[index] = e
    s.size++
}

// Remove 移除指定索引位置的元素
func (s *Stack) remove(index int) E {
    s.rangeCheck(index)
    old := s.get(index)
    for i := index; i < s.size-1; i++ {
        s.elements[i] = s.elements[i+1]
    }
    s.size--
    s.elements[s.size] = nil
    s.trim()
    return old
}

// Clear 清空
func (s *Stack) Clear() {
    for i := 0; i < s.size; i++ {
        s.elements[i] = nil
    }
    s.size = 0
    if s.elements != nil && len(s.elements) > DEFAULT_CAPACITY {
        s.elements = make([]E, DEFAULT_CAPACITY)
    }
}

func (s *Stack) ensureCapacity(capacity int) {
    oldCapacity := len(s.elements)
    if oldCapacity >= capacity {
        return
    }
    // 新容量为旧容量的1.5倍
    newCapacity := oldCapacity + (oldCapacity >> 1)
    newElements := make([]E, newCapacity)
    for i := 0; i < s.size; i++ {
        newElements[i] = s.elements[i]
    }
    s.elements = newElements
}

func (s *Stack) trim() {
    oldCapacity := len(s.elements)
    if s.size >= (oldCapacity>>1) || oldCapacity <= DEFAULT_CAPACITY {
        return
    }
    // 剩余空间还很多
    newCapacity := oldCapacity >> 1
    newElements := make([]E, newCapacity)
    for i := 0; i < s.size; i++ {
        newElements[i] = s.elements[i]
    }
    s.elements = newElements
}

// 索引越界panic
func (s *Stack) outOfBounds(index int) {
    panic(fmt.Sprintf("Index: %d Size: %d", index, s.size))
}

// 索引越界检查
func (s *Stack) rangeCheck(index int) {
    if index < 0 || index >= s.size {
        s.outOfBounds(index)
    }
}

// 针对新增元素的索引越界检查
func (s *Stack) rangeCheckForAdd(index int) {
    if index < 0 || index > s.size {
        s.outOfBounds(index)
    }
}

// 打印数组时的字符串表示
// [xx, xx, xx, xx]
func (s *Stack) String() string {
    alStr := "["
    for i := 0; i < s.size; i++ {
        if i != 0 {
            alStr = alStr + ", "
        }
        alStr = alStr + fmt.Sprintf("%#v", s.elements[i])
    }
    return alStr + "]"
}
