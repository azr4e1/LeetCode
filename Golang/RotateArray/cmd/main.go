package main

import (
	"fmt"
	"leetcode"
)

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := 3
	nums := []int{-1, -100, 3, 99}
	k := 2
	leetcode.Rotate3(nums, k)
	fmt.Println(nums)
}
