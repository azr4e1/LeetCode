package main

import (
	"fmt"
	"leetcode"
)

func main() {
	haystack := "mississipi"
	needle := "issip"
	result := leetcode.StrStr(haystack, needle)
	fmt.Println(result)
}
