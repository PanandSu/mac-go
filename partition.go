package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 0, 1, 2, 0, 0, 1, 2, 1, 0, 2, 0, 1, 1, 0},
		{2, 2, 2, 1, 1, 0, 0, 0},
		{2, 2, 2, 2},
		{1, 1, 1},
		{0, 0},
		{0, 0, 1, 1, 2},
		{},
	}
	for _, test := range tests {
		partition(test)
		fmt.Println(test)
	}
}

func partition(nums []int) {
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++ //不能省略
		} else if nums[j] == 1 {
			j++
		} else if nums[j] == 2 {
			nums[k], nums[j] = nums[j], nums[k]
			k--
		}
	}
}
