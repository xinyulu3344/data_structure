package dynamic_array

import (
    "fmt"
    "testing"
)

func TestDynamicArray_FindAll(t *testing.T) {
    array := NewDynamicArray(3)
    array.AddInLast("呵呵")
    array.AddInLast("哈哈")
    array.AddInLast("啦啦")
    array.Add(2, 1)
    fmt.Println(array)
}

//func TestDynamicArray_String(t *testing.T) {
//    array := NewDynamicArray(10)
//    array.AddInLast("aa")
//    array.AddInLast("bb")
//    array.AddInLast("cc")
//    fmt.Println(array)
//}
