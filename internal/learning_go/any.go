package learning_go

import (
	"fmt"
	"net/http"
)

type Any struct {
	Value interface{}
}

func main() {
	fmt.Println("Hello, World!")
	h := HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	h.ServeHTTP(nil, nil)
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

type MyHandler struct {
	Data string
}

func (mh MyHandler) HelperFunction() string {
	return "Helped"
}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	callMyHelper := mh.HelperFunction()
	useMyData := mh.Data
	// Do server stuff
	fmt.Println(callMyHelper, useMyData)
}

func requiresHandlerInterface(h Handler) {
	fmt.Println(h)
}

func anyFuncLikeThis(f func(w http.ResponseWriter, r *http.Request)) {
}

func doesItWork() {
	h := MyHandler{
		Data: "Hello, World!",
	}
	requiresHandlerInterface(h)
	anyFuncLikeThis(h.ServeHTTP)
}
