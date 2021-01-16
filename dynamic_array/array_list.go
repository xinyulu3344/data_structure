package dynamic_array

import (
    "bytes"
    "fmt"
)

type ArrayInterface interface {
    // 获取元素数量
    Size() int

    // 判断元素是否为空
    IsEmpty() bool

    // 是否包含某个元素
    Contains(e interface{}) bool

    // 追加元素
    Append(e interface{})

    // 获取指定索引的元素
    Get(i int) interface{}

    // 设置指定索引位置的元素的值
    Set(i int, e interface{})

    // 指定索引位置插入元素
    Insert(i int, e interface{})

    // 删除指定索引位置的元素
    Remove(i int) interface{}

    // 查看元素索引位置
    IndexOf(e interface{})

    // 清除所有元素
    Clear()
}


type DynamicArray struct {
    data []interface{}
    size int
}

func NewDynamicArray(capacity int) *DynamicArray {
    arr := &DynamicArray{}
    arr.data = make([]interface{}, capacity)
    arr.size = 0
    return arr
}


// 获取数组容量
func (d *DynamicArray) GetCapacity() int {
    return len(d.data)
}

// 获取数组元素个数
func (d *DynamicArray) GetSize() int {
    return d.size
}

// 判断数组是否为空
func (d *DynamicArray) IsEmpty() bool {
    return d.size == 0
}

// 容量调整
func (d *DynamicArray) resize(newCapacity int) {
    newArr := make([]interface{}, newCapacity)
    for i := 0; i < d.size; i++ {
        newArr[i] = d.data[i]
    }
    d.data = newArr
}

// 查找元素返回首个索引，不存在返回 -1
func (d *DynamicArray) Find(element interface{}) int {
    for i:= 0; i < d.size; i++ {
        if element == d.data[i] {
            return i
        }
    }
    return -1
}

// 查找元素返回所有索引的切片
func (d *DynamicArray) FindAll(element interface{}) (indexs []int) {
    for i:= 0; i < d.size; i++ {
        if element == d.data[i] {
            indexs = append(indexs, i)
        }
    }
    return
}


// 查看数组是否存在某个元素
func (d *DynamicArray) Contains(element interface{}) bool {
    if d.Find(element) == -1 {
        return false
    }
    return true
}

// 获取索引对应元素
func (d *DynamicArray) Get(index int) interface{} {
    if index < 0 || index > d.size - 1 {
        panic("Get failed, index out of range")
    }
    return d.data[index]
}

// 修改索引对应元素
func (d *DynamicArray) Set(index int, element interface{}) {
    if index < 0 || index > d.size - 1 {
        panic("Set failed, index out of range")
    }
    d.data[index] = element
}

// 添加元素
func (d *DynamicArray) Add(index int, element interface{}) {
    if index < 0 || index > d.GetCapacity() {
        panic("Add failed, index out of range")
    }
    // 如果数组已满则扩容
    if d.size == d.GetCapacity() {
        d.resize(2 * d.size)
    }
    // 元素后移，腾出插入位置
    for i := d.size - 1; i >= index; i-- {
        d.data[i + 1] = d.data[i]
    }
    d.data[index] = element
    d.size++
}

// 末尾追加元素
func (d *DynamicArray) AddInLast(element interface{}) {
    d.Add(d.size, element)
}

// 开头插入元素
func (d *DynamicArray) AddInFirst(element interface{}) {
    d.Add(0, element)
}

// 移除指定索引的元素
func (d *DynamicArray) Remove(index int) interface{} {
    if index < 0 || index >= d.size {
        panic("Remove failed, index out of range")
    }
    removeEle := d.data[index]
    // 从 index 之后的元素，都向前移动一个位置
    for i := index; i < d.size; i++ {
        d.data[i-1] = d.data[i]
    }
    d.size--
    // 清理最后一个元素
    d.data[d.size] = nil
    // 考虑边界情况，不能resize为0
    if d.size == d.GetCapacity() / 4 && d.GetCapacity() / 2 != 0{
        d.resize(d.GetCapacity() / 2)
    }
    return removeEle
}


func (d *DynamicArray) RemoveFirst() interface{} {
    return d.Remove(0)
}

func (d *DynamicArray) RemoveLast() interface{} {
    return d.Remove(d.size - 1)
}


func (d *DynamicArray) String() string {
    var buffer bytes.Buffer
    buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", d.size, d.GetCapacity()))
    buffer.WriteString("[")
    for i := 0; i < d.size; i++ {
        buffer.WriteString(fmt.Sprint(d.data[i]))
        if i != d.size - 1{
            buffer.WriteString(",")
        }
    }
    buffer.WriteString("]")
    return buffer.String()
}
