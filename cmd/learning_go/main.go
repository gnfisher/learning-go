package main

import (
	"fmt"
	"os"

	"github.com/gnfisher/learning_go/internal/age"
	"github.com/gnfisher/learning_go/internal/learning_go"
)

func main() {
	err, a := age.FromInt(10)
	if err == nil {
		fmt.Println(a.HowManyCandles())
	}

	err, b := age.FromInt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b.HowManyCandles())
	}

	learning_go.SayHello()
	learning_go.ChapterTwo()
	learning_go.SliceAppend()
	learning_go.SliceExpression()
	learning_go.SliceArrayExp()
	learning_go.StringSlices()
	learning_go.Maps()
	learning_go.Structs()

	learning_go.Play()

	learning_go.PointerPlay()

	var thing learning_go.IThing
	inner := learning_go.InnerA{Age: 38}
	thing = &learning_go.OuterA{
		InnerA: inner,
		Name:   "Greg",
	}

	fmt.Println(thing.HowOld())

	x := make([]interface{}, 3)
	x[0] = 1
	x[1] = "hello"
	fmt.Println(x)

	fmt.Println("Ch8...")
	fmt.Println(learning_go.InterfaceIsNil())
	learning_go.PlayWithErrors()

	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("There was an error opening the file")
		return
	}
	defer file.Close()

	m, err := learning_go.CountChars(file)
	if err != nil {
		fmt.Println("There was an error counting the characters")
		return
	}

	fmt.Println(m)
}
