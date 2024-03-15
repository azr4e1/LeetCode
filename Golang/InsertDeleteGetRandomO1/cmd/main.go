package main

import (
	"fmt"
	"leetcode"
)

func main() {
	set := leetcode.Constructor()
	fmt.Println("Set:", set)
	fmt.Println("Remove 0:", set.Remove(0))
	fmt.Println("Set:", set)
	fmt.Println("Remove 0:", set.Remove(0))
	fmt.Println("Set:", set)
	fmt.Println("Insert 0:", set.Insert(0))
	fmt.Println("Set:", set)
	fmt.Println("Get Random:", set.GetRandom())
	fmt.Println("Remove 0:", set.Remove(0))
	fmt.Println("Set:", set)
	fmt.Println("Insert 0:", set.Insert(0))
	fmt.Println("Set:", set)
}
