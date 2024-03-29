package heap

import (
    "runtime/debug"
    "testing"
)

func TestAdd(t *testing.T) {
    heap := NewBinaryHeapWithComparator(func(e1, e2 any) int {
        return e1.(int) - e2.(int)
    })
    data := []int{68, 72, 43, 50, 38}
    for _, v := range data {
        heap.Add(v)
    }
    // t.Log(heap.elements) // [72 68 43 50 38 <nil> <nil> <nil> <nil> <nil>]
    assert(t, heap.Size() == 5)
    assert(t, heap.Get() == 72)
    assert(t, heap.elements[1] == 68)
    assert(t, heap.elements[2] == 43)
    assert(t, heap.elements[3] == 50)
    assert(t, heap.elements[4] == 38)
}

func TestRemove(t *testing.T) {
    heap := NewBinaryHeapWithComparator(func(e1, e2 any) int {
        return e1.(int) - e2.(int)
    })
    data := []int{68, 72, 43, 50, 38, 10, 90, 65}
    for _, v := range data {
        heap.Add(v)
    }
    // t.Log(heap.elements) // [90 68 72 65 38 10 43 50 <nil> <nil>]
    assert(t, heap.Remove() == 90)
    assert(t, heap.Size() == 7)
    // t.Log(heap.elements) // [72 68 50 65 38 10 43 <nil> <nil> <nil>]
    out := []any{72, 68, 50, 65, 38, 10, 43, nil, nil, nil}
    for i := 0; i < heap.size; i++ {
        if out[i] != heap.elements[i] {
            assert(t, false)
        }
    }
}

func TestReplace(t *testing.T) {
    heap := NewBinaryHeapWithComparator(func(e1, e2 any) int {
        return e1.(int) - e2.(int)
    })
    data := []int{68, 72, 43, 50, 38, 10, 90, 65}
    for _, v := range data {
        heap.Add(v)
    }
    root, _ := heap.Replace(70)
    assert(t, root == 90)
    assert(t, heap.Size() == 8)

    out := []any{72, 68, 70, 65, 38, 10, 43, 50}
    for i := 0; i < heap.size; i++ {
        if out[i] != heap.elements[i] {
            assert(t, false)
        }
    }
}

func TestHeapify(t *testing.T) {
    data := []any{68, 72, 43, 50, 38, 10, 90, 65}
    heap := NewBinaryHeapify(data, func(e1, e2 any) int {
        return e1.(int) - e2.(int)
    })
    out := []any{90, 72, 68, 65, 38, 10, 43, 50}
    for i := 0; i < heap.size; i++ {
        if out[i] != heap.elements[i] {
            assert(t, false)
        }
    }
}

func TestTopK(t *testing.T) {
	data := []any{8,2,7,4,5,18,8,99,10}
	// 构建小顶堆
    heap := NewBinaryHeapWithComparator(func(e1, e2 any) int {
        return e2.(int) - e1.(int)
    })
	k := 5
	for i :=0 ;i < len(data); i++ {
		if heap.size < k {
			heap.Add(data[i])
		} else if data[i].(int) > heap.Get().(int) {
			heap.Replace(data[i])
		} 
	}
	t.Log(heap.elements...)
}

func assert(t *testing.T, ok bool) {
    if !ok {
        debug.PrintStack()
        t.FailNow()
    }
}
