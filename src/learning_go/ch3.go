package learning_go

import "fmt"

func SliceAppend() {
	x := []int{1, 2, 3}
	fmt.Println(x, len(x), cap(x))
	x = append(x, 4)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 5)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 6)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 7)
	fmt.Println(x, len(x), cap(x))

	// make
	y := make([]int, 5)
	fmt.Println(y, len(y), cap(y))
	y = make([]int, 5, 10)
	fmt.Println(y, len(y), cap(y))
	z := make([]int, 0, 10) // capacity of 10, non nil, zero length, initializing an empty one
	fmt.Println(z == nil)
	z = append(z, 1, 2, 3)
	fmt.Println(z)

	// slice expressions
	// OFFSETS not INDEXES
	a := x[1:2]
	fmt.Println(a) // [2]
	fmt.Println(x) // [1 2 3 4 5 6 7]
	b := x[1:7]
	fmt.Println(b) // [2 3 4 5 6 7]

	h := []int{1, 2, 3, 4}
	j := h[:2]
	fmt.Println(cap(h), cap(j))
	j = append(j, 30)
	fmt.Println("h:", h)
	fmt.Println("j:", j)
}

func SliceExpression() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(x, y, z)
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func SliceArrayExp() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 3)
	num := copy(y, x)
	fmt.Println(y, num)

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	n := copy(b, a)
	fmt.Println(b, n)
}

func StringSlices() {
	s := "Hello there"
	b := s[6]
	fmt.Println(b)

	c := s[0:5]
	fmt.Println(c)

	e := "Hello ☀️"
	fmt.Println(e[6])
	fmt.Println(e[4:])
	fmt.Println(len(e))
	var b2 byte = 'y'
	var r rune = 'x'
	fmt.Println(string(116))
	fmt.Println(b2)
	fmt.Println(string(b2))
	fmt.Println(r)
	fmt.Println(string(r))
}

func Maps() {
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams)

	for k, v := range teams {
		fmt.Println(k)
		for _, v := range v {
			fmt.Println(v)
		}
	}

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}

type person struct {
	name string
	age  int
	pet  string
}

func Structs() {
	var greg person
	daniela := person{}
	lucas := person{
		"Lucas",
		8,
		"fish",
	}
	mia := person{
		name: "Mia",
		age:  6,
		pet:  "Dog",
	}
	fmt.Println(greg)
	fmt.Println(daniela)
	fmt.Println(lucas, mia)
}
