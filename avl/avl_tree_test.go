package bstree

type Person struct {
    name string
    age  int
}

func (p *Person) CompareTo(e E) int {
    return p.age - e.(*Person).age
}

type Int int

func (i Int) CompareTo(e E) int {
    return int(i - e.(Int))
}
