package main

import (
	"fmt"
	"leetcode"
)

func main() {
	input := []int{6, 4, 54, 23, 65, 3, 4, 56, 324, 5, 7, 87, 6, 6, 6, 5, 3, 654, 7}
	fmt.Println(leetcode.NewList(input))
	fmt.Println(leetcode.InsertionSortList(leetcode.NewList(input)))
}
