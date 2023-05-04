package learning_go

import "fmt"

type OuterA struct {
	InnerA
	Name string
}

type InnerA struct {
	Age int
}

type IThing interface {
	HowOld() string
}

func (i *InnerA) HowOld() string {
	return fmt.Sprintf("I am %d years old", i.Age)
}
