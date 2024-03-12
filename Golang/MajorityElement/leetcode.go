package leetcode

import "fmt"

func MajorityElement(nums []int) int {
	var maj int
	var greatest int
	var counts = make(map[int]int)
	for _, n := range nums {
		counts[n]++
		if counts[n] > greatest {
			maj = n
			greatest = counts[n]
		}
		if greatest >= len(nums)/2+1 {
			break
		}
	}
	fmt.Println(counts)
	return maj
}
