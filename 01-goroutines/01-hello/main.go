package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	ex1()
}

func ex1() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("go-routine-1")

	// goroutine with anonymous function
	go func() {
		fun("go routine-2")
	}()
	// goroutine with function value call
	fv := fun
	go fv("go routine-3")
	// wait for goroutines to end
	fmt.Println("waiting for go routine")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done..")
}
