package treemap

import (
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

func TestTreeMap(t *testing.T) {
    tm := NewTreeMap()
    tm.Put(Str("aaa"), 11)
    tm.Put(Str("bbb"), 22)
    tm.Put(Str("aaa"), 33)

    if tm.Size() != 2 {
        t.Errorf("tm.Size(): %v\n", tm.Size())
    }

    if !tm.ContainsKey(Str("aaa")) {
        t.Errorf("tm.ContainsKey(Str(\"aaa\")): %v\n", tm.ContainsKey(Str("aaa")))
    }

    containsValue := tm.ContainsValue(33, func(v1, v2 any) bool {
        if v1 == v2 {
            return true
        } else {
            return false
        }
    })

    if !containsValue {
        t.Errorf("tm.ContainsValue(33): %v\n", containsValue)
    }

    if tm.Get(Str("aaa")) != 33 {
        t.Errorf("tm.Get(Str(\"aaa\")): %v\n", tm.Get(Str("aaa")))
    }

    tm.Traversal(func(key Key, value any) bool {
        t.Log(key, value)
        return true
    })

    tm.Remove(Str("bbb"))
    if tm.Get(Str("bbb")) != nil {
        t.Errorf("tm.Get(Str(\"bbb\")): %v\n", tm.Get(Str("bbb")))
    }

    tm.Clear()

    if !tm.IsEmpty() {
        t.Errorf("tm.IsEmpty(): %v\n", tm.IsEmpty())
    }
}