/**
  @author: xinyulu
  @date: 2021/1/31 20:40
  @note:
**/
package set

import (
    "container/list"
)


// 使用链表实现
type ListSet struct {
    ls *list.List
}

func NewListSet() *ListSet {
    return &ListSet{
        ls: list.New(),
    }
}

func (lSet *ListSet) Size() int {
    return lSet.ls.Len()
}

func (lSet *ListSet) IsEmpty() bool {
    return lSet.Size() == 0
}

func (lSet *ListSet) Clear() {
    lSet.ls.Init()
}

func (lSet *ListSet) Contains(e interface{}) bool {
    cur := lSet.ls.Front()
    for cur != nil {
        if cur.Value == e {
            return true
        }
        cur = cur.Next()
    }
    return false
}

func (lSet *ListSet) Add(e interface{}) {
    if lSet.Contains(e) {
        return
    }
    lSet.ls.PushBack(e)
}

func (lSet *ListSet) Remove(e interface{}) {
    n := lSet.findNodeByElement(e)
    if n != nil {
        lSet.ls.Remove(n)
    }
}

func (lSet *ListSet) traversal(f func(e interface{})) {
    if f == nil {
        return
    }
    size := lSet.Size()
    node := lSet.ls.Front()
    for i := 0; i < size; i++ {
        f(node.Value)
        node = node.Next()
    }
}

func (lSet *ListSet) indexOf(e interface{}) int {
    if e == nil {
        node := lSet.ls.Front()
        for i := 0; i < lSet.Size(); i++ {
            if node.Value == nil {
                return i
            }
            node = node.Next()
        }
    } else {
        node := lSet.ls.Front()
        for i := 0; i < lSet.Size(); i++ {
            if node.Value == e {
                return i
            }
            node = node.Next()
        }
    }
    return -1
}

// 根据元素找到节点
func (lSet *ListSet) findNodeByElement(e interface{}) *list.Element {
    cur := lSet.ls.Front()
    for cur != nil {
        if cur.Value == e {
            return cur
        }
        cur = cur.Next()
    }
    return nil
}
