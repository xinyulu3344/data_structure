package sorts

import (
    "math/rand"
    "strconv"
    "testing"
)

type Person struct {
    name string
    age int
}

type ByAge []Person

func NewByAge(count int) ByAge {
    persons := make(ByAge, count)
    for i := 0; i < count; i++ {
        persons[i] = Person{
            name: "xinyulu" + strconv.Itoa(i),
            age: rand.Intn(count),
        }
    }
    return persons
}
func (b ByAge) Len() int { return len(b) }
func (b ByAge) Compare(i, j int) int { return b[i].age - b[j].age }
func (b ByAge) Swap(i, j int) { b[i], b[j] = b[j], b[i] }


func TestBubbleAsSort(t *testing.T) {
    s1 := NewByAge(10000)
    bs := NewBubbleSort(true)
    t.Log(s1)
    bs.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", bs.cmpCount)
    t.Log("swapCount: ", bs.swapCount)
    t.Log(IsAsSorted(s1))
}

func TestBubbleAsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    bs := NewBubbleSort(true)
    bs.SortInt(randInts)
    t.Log(randInts)
    t.Log("cmpCount: ", bs.cmpCount)
    t.Log("swapCount: ", bs.swapCount)
    t.Log(IntsAreAsSorted(randInts))
}

func TestBubbleSortDsSort(t *testing.T) {
    s1 := NewByAge(10000)
    bs := NewBubbleSort(false)
    t.Log(s1)
    bs.Sort(s1)
    t.Log(s1)
    t.Log("cmpCount: ", bs.cmpCount)
    t.Log("swapCount: ", bs.swapCount)
    t.Log(IsDsSorted(s1))
}

func TestBubbleSortDsSortInt(t *testing.T) {
    randInts := rand.Perm(10000)
    t.Log(randInts)
    
    bs := NewBubbleSort(false)
    bs.SortInt(randInts)
    t.Log(randInts)
    t.Log("cmpCount: ", bs.cmpCount)
    t.Log("swapCount: ", bs.swapCount)
    t.Log(IntsAreDsSorted(randInts))
}

func BenchmarkBubbleAsSortInt(b *testing.B) {
    bs := NewBubbleSort(true)
    for i := 0; i < b.N; i++ {
        randInts := rand.Perm(10000)
        bs.SortInt(randInts)
    }
}

func BenchmarkBubbleAsSort(b *testing.B) {
    bs := NewBubbleSort(true)
    for i := 0; i < b.N; i++ {
        s1 := NewByAge(10000)
        bs.Sort(s1)
    }
}