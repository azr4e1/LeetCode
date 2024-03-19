// You are given a sorted unique integer array nums.

// A range [a,b] is the set of all integers from a to b (inclusive).

// Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

// Each range [a,b] in the list should be output as:

//     "a->b" if a != b
//     "a" if a == b

package leetcode

import "strconv"

const SEPARATOR = "->"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	ranges := []string{}
	var start int = nums[0]
	var increment int
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			if increment == 0 {
				ranges = append(ranges, strconv.Itoa(start))
			} else {
				ranges = append(ranges, strconv.Itoa(start)+SEPARATOR+strconv.Itoa(start+increment))
			}
			start = nums[i]
			increment = 0
			continue
		}
		increment++
	}
	if increment == 0 {
		ranges = append(ranges, strconv.Itoa(start))
	} else {
		ranges = append(ranges, strconv.Itoa(start)+SEPARATOR+strconv.Itoa(start+increment))
	}

	return ranges
}
