package main

import "leetcode"
import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	list := leetcode.FromSlice(input)
	fmt.Println(list)

	fmt.Println(leetcode.FromList(list))
	newList := leetcode.SwapPairs(list)
	fmt.Println(newList)
}
