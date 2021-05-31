/**
  @author: xinyulu
  @date: 2021/1/31 20:39
  @note:
**/
package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
    listSet := NewListSet()
    listSet.Add(10)
    listSet.Add(11)
    listSet.Add(11)
    listSet.Add(12)
    listSet.traversal(func(e interface{}) {
        fmt.Println(e)
    })
}
