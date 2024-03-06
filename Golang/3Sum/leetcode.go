// Given an integer array nums, return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k]
// == 0.
// Notice that the solution set must not contain duplicate triplets.
package leetcode

import "slices"
import "fmt"

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	tracker := make(map[[3]int]bool)

	for i := 0; i < len(nums)-2; i++ {
		ledger := make(map[int]int)
		el := nums[i]
		for _, num := range nums[i+1:] {
			val, ok := ledger[num]
			if ok {
				finalArray := [3]int{el, val, num}
				finalSlice := finalArray[:]
				slices.Sort(finalSlice)
				_, ok2 := tracker[finalArray]
				if ok2 {
					continue
				}
				result = append(result, finalSlice)
				tracker[finalArray] = true
				continue
			}
			ledger[-el-num] = num
		}
	}
	fmt.Println(result)
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
