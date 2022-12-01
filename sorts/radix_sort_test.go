package sorts

import (
	"strconv"
    "math/rand"
    "testing"
)

type MyElement struct {
	name string
	score int
}

func NewMyElement(name string, score int) MyElement {
	return MyElement{
		name: name,
		score: score,
	}
}

func (me MyElement) GetNum() int {
	return me.score
}

type MyElements []MyElement

func (me MyElements) Len() int {
	return len(me)
}

func (me MyElements) GetElement(index int) Element {
	return me[index]
}

func (me MyElements) SetElement(index int, element Element) {
	me[index] = element.(MyElement)
}

func TestRadixSortInt(t *testing.T) {
    data := rand.Perm(10000)
    rs := NewRadixSort()
    rs.SortInt(data)
    t.Log(IntsAreAsSorted(data))
}

func TestRadixSort(t *testing.T) {
	rs := NewRadixSort()
	data := make(MyElements, 10)
	for i := 0; i < 10; i++ {
		data[i] = NewMyElement("xinyulu" + strconv.Itoa(i), rand.Intn(10))
	}
	t.Log(data)
	rs.Sort(data)
	t.Log(data)
}