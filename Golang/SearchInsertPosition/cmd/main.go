package main

import (
	"fmt"
	"leetcode"
)

func main() {
	input := []int{3, 5, 7, 9, 10}
	target := 8
	fmt.Println(leetcode.SearchInsert(input, target))
}
