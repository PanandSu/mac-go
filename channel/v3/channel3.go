package main

import (
	"fmt"
	"sync"
)

var token = struct{}{}
var chs = []chan struct{}{
	make(chan struct{}, 1),
	make(chan struct{}, 1),
	make(chan struct{}, 1),
}
var wg sync.WaitGroup

func main() {
	// printDownpers,printUppers,printNumbers顺序无关紧要
	go printDownpers()
	go printUppers()
	go printNumbers()
	wg.Add(3)
	chs[0] <- token
	wg.Wait()

	/*
		1
		A
		a
		2
		B
		b
		3
		C
		c
		4
		D
		d
		5
		E
		e
		6
		F
		f
		7
		G
		g
		8
		H
		h
		9
		I
		i
		10
		J
		j
		11
		K
		k
		12
		L
		l
		13
		M
		m
		14
		N
		n
		15
		O
		o
		16
		P
		p
		17
		Q
		q
		18
		R
		r
		19
		S
		s
		20
		T
		t
		21
		U
		u
		22
		V
		v
		23
		W
		w
		24
		X
		x
		25
		Y
		y
		26
		Z
		z
	*/
}

func printNumbers() {
	defer wg.Done()
	for i := 1; i <= 26; i++ {
		<-chs[0]
		fmt.Println(i)
		chs[1] <- token
	}
}

func printUppers() {
	defer wg.Done()
	for up := 'A'; up <= 'Z'; up++ {
		<-chs[1]
		fmt.Println(string(up))
		chs[2] <- token
	}
}

func printDownpers() {
	defer wg.Done()
	for down := 'a'; down <= 'z'; down++ {
		<-chs[2]
		fmt.Println(string(down))
		chs[0] <- token
	}
}
