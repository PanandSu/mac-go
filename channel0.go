package main

import (
	"fmt"
)

func main() {
	// chr =>char
	// c,ch =>chan
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
	go func() {
		ch <- 2
		ch <- 3
	}()
	go func() {
		<-ch
		<-ch
	}()
	fmt.Println("the program is done")
}
