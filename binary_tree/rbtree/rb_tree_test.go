package rbtree

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 实现了Comparable接口, 该类型可以被比较
type persion struct {
    name string
    age  int
}

func NewPersion(name string, age int) *persion {
    return &persion{
        name: name,
        age:  age,
    }
}

func (p *persion) CompareTo(e Comparable) int {
    v, ok := e.(*persion)
    if !ok {
        panic("传入的值不是*persion类型!")
    }
    return p.age - v.age
}

type myVistor struct {
    stoppedElement int
    stopped bool
}

func (m *myVistor) Visit(e interface{}, color bool) {
    element, ok := e.(*persion)
    if !ok {
        fmt.Println("断言错误")
    }
    if color {
        fmt.Printf("%d_%s,", element.age, "black")
    } else {
        fmt.Printf("%d_%s,", element.age, "red")
    }
    // 如果符合条件，就将stopped置为true，表示要停止遍历
    if element.age == m.stoppedElement {
        m.stopped = true
    }
}

func (m *myVistor) Stop() bool {
    return m.stopped
}

type SearchVistor struct {
    searchIndex int
    stopped bool
}

func (s *SearchVistor) Visit(e interface{}, color bool) {
    element, ok := e.(*persion)
    if !ok {
        panic("传入e的类型不是*persion")
    }
    // 如果符合条件，就将stopped置为true，表示要停止遍历
    if element.age == s.searchIndex {
        s.stopped = true
        fmt.Println(element)
    }
}

func (s *SearchVistor) Stop() bool {
    return s.stopped
}

func TestInnerOrderTraversal(t *testing.T) {
    data := []persion{
        {name: "a", age: 55},
        {name: "b", age: 87},
        {name: "c", age: 56},
        {name: "d", age: 74},
        {name: "e", age: 96},
        {name: "f", age: 22},
        {name: "g", age: 62},
        {name: "h", age: 20},
        {name: "i", age: 70},
        {name: "j", age: 68},
        {name: "k", age: 90},
        {name: "l", age: 50},
    }
    rbTree := NewRBTree()
    for i := 0; i < len(data); i++ {
        rbTree.Add(&data[i])
    }
    rbTree.InnerOrderTraversal(&myVistor{stoppedElement: 74})
}

func TestRemove(t *testing.T) {
    data := make([]persion, 0)
    for i := 0; i < 10; i++ {
        data = append(data, persion{name: "a" + strconv.Itoa(i), age: i})
    }
    rbTree := NewRBTree()
    for i := 0; i < len(data); i++ {
        rbTree.Add(&data[i])
    }
    rbTree.InnerOrderTraversal(&myVistor{stoppedElement: 100})
    fmt.Println()
    for i := 0; i < len(data); i++ {
        rbTree.Remove(&data[i])
        rbTree.InnerOrderTraversal(&myVistor{stoppedElement: 100})
        fmt.Println()
    }
    fmt.Println("==========删除结束============")
    rbTree.InnerOrderTraversal(&myVistor{})
}


func TestSearch(t *testing.T) {
    data := make([]persion, 0)
    for i := 0; i < 10000000; i++ {
        data = append(data, persion{name: "a" + strconv.Itoa(i), age: i})
    }
    rbTree := NewRBTree()
    for i := 0; i < len(data); i++ {
        rbTree.Add(&data[i])
    }
    startTime := time.Now()
    rbTree.InnerOrderTraversal(&SearchVistor{searchIndex: 1000000})
    fmt.Println("查询耗时: ", time.Since(startTime).Seconds())
}