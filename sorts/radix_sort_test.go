package sorts

import (
	"strconv"
    "math/rand"
    "testing"
)

type MyRadixSort struct {
	name string
	score int
}

func NewMyRadixSort(name string, score int) MyCountSort {
	return MyCountSort{
		name: name,
		score: score,
	}
}

func (mrs MyRadixSort) GetNum() int {
	return mrs.score
}

func TestRadixSortInt(t *testing.T) {
    data := rand.Perm(10000)
    rs := NewRadixSort()
    rs.SortInt(data)
    t.Log(IntsAreAsSorted(data))
}

func TestRadixSort(t *testing.T) {
	rs := NewRadixSort()
	data := make(Elements, 10)
	for i := 0; i < 10; i++ {
		data[i] = NewMyRadixSort("xinyulu" + strconv.Itoa(i), rand.Intn(10))
	}
	t.Log(data)
	rs.Sort(data)
	t.Log(data)
}