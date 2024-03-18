// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
package leetcode

import "fmt"

func ContainsNearbyDuplicate(nums []int, k int) bool {
	valuesIndex := make(map[int]int)
	for i, v := range nums {
		val, ok := valuesIndex[v]
		if ok {
			fmt.Println(i, v)
			if Abs(val-i) <= k {
				return true
			}
		}
		valuesIndex[v] = i
	}
	return false
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
