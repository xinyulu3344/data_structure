package sorts

import (
	"math/rand"
	"strconv"
	"testing"
)

type MyCountSort struct {
	name string
	score int
}

func NewMyCountSort(name string, score int) MyCountSort {
	return MyCountSort{
		name: name,
		score: score,
	}
}

func (mcs MyCountSort) GetNum() int {
	return mcs.score
}

func TestCountingSort_SortInt(t *testing.T) {
    cs := NewCountingSort()
    data := rand.Perm(100000)
    // t.Log(data)
    cs.SortInt(data)
    // t.Log(data)
    t.Log(IntsAreAsSorted(data))
}

func TestCountingSort_Sort(t *testing.T) {
	cs := NewCountingSort()
	data := make(Elements, 10)
	for i := 0; i < 10; i++ {
		data[i] = NewMyCountSort("xinyulu" + strconv.Itoa(i), rand.Intn(10))
	}
	t.Log(data)
	cs.Sort(data)
	t.Log(data)
}