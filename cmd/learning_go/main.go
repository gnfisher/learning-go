package main

import (
	"fmt"

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
}
