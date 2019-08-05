package main

import (
	"react_go_wasm/gosrc/todo"
)

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized and changed")
	// register functions
	todo.RegisterCallbacks()
	<-c
}
