package learning_go

import (
	"errors"
	"fmt"
)

func Play() {
	fmt.Println("Functions:")
	fmt.Println("div()", div(1, 2))
	fmt.Println("addTo()", addTo(1, 2, 3, 4, 5))
	fmt.Println(multiReturn(0, ""))
	fmt.Println(multiReturn(38, "Greg"))
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func multiReturn(age int, name string) (err error, sentence string) {
	if name == "" || age == 0 {
		err := errors.New("Can't do it, man.")
		return err, sentence
	}
	return err, fmt.Sprintf("%s is %d years old", name, age)
}
