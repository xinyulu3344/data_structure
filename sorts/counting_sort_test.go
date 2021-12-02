package sorts

import (
	"math/rand"
	"strconv"
	"testing"
)


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
	data := make(MyElements, 10)
	for i := 0; i < 10; i++ {
		data[i] = NewMyElement("xinyulu" + strconv.Itoa(i), rand.Intn(10))
	}
	t.Log(data)
	cs.Sort(data)
	t.Log(data)
}