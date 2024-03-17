// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.
package leetcode

func FindMaxLength(nums []int) int {
	maxLength := 0
	n := len(nums)
	sums := make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	for i := 0; i < n-1; i++ {
		initialPos := sums[i]
		for j := maxLength + i + 1; j < n; j += 2 {
			length := j - i + 1
			if (sums[j]-initialPos+nums[i]) == length/2 && length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}
