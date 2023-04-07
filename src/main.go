package main

// - in ./src run `go mod init main`
// - in ./src/learning_go `go mod init learning_go`
// Because this is a local file not a published package, I had to then do:
// - in ./src/main `go mod edit -replace learning_go=./learning_go`
// - in ./src/main `go mod tidy` - to "syncronize the dependencies from learning_go module"
// Then, the makefile commands to build and run the `src/main.go` file (or binary after building)
// To set this up I had to do the following.
// Source: https://go.dev/doc/tutorial/call-module-code
import "learning_go"

func main() {
	learning_go.SayHello()
	learning_go.ChapterTwo()
	learning_go.SliceAppend()
	learning_go.SliceExpression()
	learning_go.SliceArrayExp()
	learning_go.StringSlices()
	learning_go.Maps()
	learning_go.Structs()
}
