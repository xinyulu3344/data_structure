package treemap

import (
	"fmt"
	"testing"
)

type Str string

func (sk Str) CompareTo(key Key) int {
    k1 := string(sk)
    k2 := string(key.(Str))
    if k1 > k2 {
        return -1
    } else if k1 == k2 {
        return 0
    } else {
        return 1
    }
}

func TestXxx(t *testing.T) {
    tm := NewTreeMap()
    tm.Put(Str("aaa"), 11)
    tm.Put(Str("bbb"), 22)
    tm.Put(Str("aaa"), 33)

    fmt.Printf("tm.ContainsKey(Str(\"aaa\")): %v\n", tm.ContainsKey(Str("aaa")))

    fmt.Printf("tm.ContainsValue(33): %v\n", tm.ContainsValue(33, func(v1, v2 any) bool {
        if v1 == v2 {
            return true
        } else {
            return false
        }
    }))

    fmt.Printf("tm.Get(Str(\"aaa\")): %v\n", tm.Get(Str("aaa")))
}