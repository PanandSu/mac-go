package main

import "fmt"

func main() {
	fmt.Println(splitN(4))
}

// 4
// [1,1,1,1]
// [1,1,2]
// [1,3]
// [2,2]
// [4]
func splitN(n int) int {
	if n <= 1 {
		return n
	}
	return process(1, n)
}

func process(pre, rest int) int {
	if rest == 0 {
		return 1
	}
	if pre > rest {
		return 0
	}
	ways := 0
	for i := pre; i <= rest; i++ {
		ways += process(i, rest-i)
	}
	return ways
}

func splitSequences(n int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	res := operate(1, n, make([][]int, 0))
	return res
}

func operate(pre, rest int, lists [][]int) [][]int {
	if rest == 0 {
		return lists
	}
	return [][]int{}
}
