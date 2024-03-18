// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.
package leetcode

import "fmt"

func SearchInsert(nums []int, target int) int {
	n := len(nums)
	search := make([]int, n)
	copy(search, nums)
	a, b, m := 0, n-1, (n-1)/2
	if target > nums[b] {
		return n
	}
	if target <= nums[a] {
		return 0
	}
	for a+1 < b {
		fmt.Println(a, m, b)
		switch {
		case target == nums[m]:
			return m
		case target < nums[m]:
			b, m = m, a+(m-a)/2
		case target > nums[m]:
			a, m = m, m+(b-m)/2
		}
	}
	return m + 1
}
