package learning_go

import "fmt"

type names struct {
	FirstName string
	LastName  *string
}

type human struct {
	FirstName string
	LastName  string
}

func PointerPlay() {
	fmt.Println("Pointers")

	var y int
	z := &y
	fmt.Println(*z)

	f := "Fisher"
	ns := names{
		FirstName: "Greg",
		LastName:  &f,
	}
	fmt.Println(ns)

	g := human{
		FirstName: "Greg",
		LastName:  "Fisher",
	}
	pointerTest(&g)
	fmt.Println(g)
}

func pointerTest(p *human) {
	fmt.Println("pointerTest:")
	p.FirstName = "Gregorio"
	fmt.Println(*p)

	// This will just error
	// p = human{
	// 	FirstName: "Daniela",
	// 	LastName:  "Aguel",
	// }
}

