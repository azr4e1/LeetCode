// Given an integer array nums, return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k]
// == 0.
// Notice that the solution set must not contain duplicate triplets.
package leetcode

import (
	"fmt"
	"slices"
)

func ThreeSum(input []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(input)-1; i++ {
		pairs := TwoSum(input[i+1:], -input[i])
		for _, el := range pairs {
			finalSlice := append(el, input[i])
			slices.Sort(finalSlice)
			if !slices.ContainsFunc(result, func(sl []int) bool {
				if len(sl) != len(finalSlice) {
					return false
				}
				for j := 0; j < len(sl); j++ {
					if sl[j] != finalSlice[j] {
						return false
					}
				}
				return true
			},
			) {
				result = append(result, finalSlice)
			}
		}
	}
	return result
}

func TwoSum(input []int, sum int) [][]int {
	ledger := make(map[int]int)
	result := [][]int{}
	for _, num := range input {
		val, ok := ledger[num]
		if ok {
			result = append(result, []int{val, num})
			continue
		}
		ledger[sum-num] = num
	}

	return result
}
