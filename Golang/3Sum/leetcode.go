// Given an integer array nums, return all the triplets [nums[i], nums[j],
// nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k]
// == 0.
// Notice that the solution set must not contain duplicate triplets.
package leetcode

import "sort"
import "fmt"

func ThreeSum(nums []int) [][]int {
	var res [][]int

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target > 0 {
				right--
			} else {
				left++
			}
		}
	}

	fmt.Println(res)
	return res
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
