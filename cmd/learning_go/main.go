package main

import (
	"fmt"

	"github.com/gnfisher/learning_go/internal/age"
	"github.com/gnfisher/learning_go/internal/learning_go"
)

func main() {
	err, a := age.FromInt(10)
	if err == nil {
		fmt.Println(a)
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
}
