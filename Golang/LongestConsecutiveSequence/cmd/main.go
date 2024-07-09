package main

import (
	"fmt"
	"leetcode"
)

func main() {
	test := []int{100, 4, 200, 1, 3, 2}
	res := leetcode.LongestConsecutive(test)
	fmt.Println(res)
}
