package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator()
	// filter evens
	predicate := func(i int) bool {
		return i&1 == 0
	}
	out := filter(ch, predicate)
	fmt.Println("evens: ", readout(out))
	/*
		0
		2
		4
		6
		8
		10
		12
		14
		16
		18
		20
		22
		24
		26
		28
		evens:  [0 2 4 6 8 10 12 14 16 18 20 22 24 26 28]
	*/
}

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 30; i++ {
			ch <- i
			time.Sleep(time.Second / 10)
		}
	}()
	return ch
}

func filter(ch <-chan int, predicate func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range ch {
			if predicate(v) {
				out <- v
			}
		}
	}()
	return out
}

func readout(ch <-chan int) []int {
	var out []int
	for v := range ch {
		out = append(out, v)
		fmt.Println(v)
	}
	return out
}
