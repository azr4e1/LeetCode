// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
package leetcode

func ProductExceptSelf(nums []int) []int {
	var totalProd = 1
	var numZeros int
	output := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			numZeros++
			continue
		}
		totalProd *= nums[i]
	}
	if numZeros > 1 {
		return output
	}
	return nil
}
