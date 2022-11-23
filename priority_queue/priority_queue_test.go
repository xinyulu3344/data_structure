package priorityqueue

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	weight int
}

func NewPerson(name string, weight int) Person {
	return Person{
		name: name,
		weight: weight,
	}
}

func (p Person) CompareTo(e any) int {
	return e.(Person).weight - p.weight 
}

func TestXxx(t *testing.T) {
	q := NewPrioriryQueue()
	q.EnQueue(NewPerson("jack", 90))
	q.EnQueue(NewPerson("rose", 88))
	q.EnQueue(NewPerson("jake", 86))
	q.EnQueue(NewPerson("jim", 100))
	for !q.IsEmpty() {
		fmt.Printf("q.Dequeue(): %v\n", q.Dequeue())
	}
}