package ArrayList

import (
    "errors"
    "fmt"
)

type List interface {
    Size()                                      // 数组大小
    Get(index int) (interface{}, error)         // 抓取第几个元素
    Set(index int, newVal interface{}) error    // 修改数据
    Insert(index int, newVal interface{}) error //插入数据
    Append(newVal interface{})                  // 追加
    Clear()                                     // 清空
    Delete(index int) error                     // 删除
    String()                                    // 返回字符串
}

type ArrayList struct {
    dataStore []interface{} // 数组存储
    theSize   int           // 数组大小
}

func NewArrayList() *ArrayList {
    list := new(ArrayList) // 初始化结构体
    list.dataStore = make([]interface{}, 0, 10)
    list.theSize = 0
    return list
}

func (list *ArrayList) Size() int {
    return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
    if index < 0 || index >= list.theSize {
        return nil, errors.New("索引越界")
    }
    return list.dataStore[index], nil
}

// 追加数据
func (list *ArrayList) Append(newVal interface{}) {
    list.dataStore = append(list.dataStore, newVal)
    list.theSize++
}

func (list *ArrayList) String() string {
    return fmt.Sprint(list.dataStore)
}

