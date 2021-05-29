/**
  @author: xinyulu
  @date: 2021/2/1 0:02
  @note: 
**/
package treemap

type Visitor func(k, v interface{}) bool

type TreeMap interface {
    GetSize() int
    IsEmpty() bool
    Clear()
    Put(k, v interface{}) interface{}
    Get(k interface{}) interface{}
    Remove(k interface{}) interface{}
    ContainsKey(k interface{}) bool
    ContainsValue(v interface{}) bool
    Traversal(visitor Visitor) bool
}
