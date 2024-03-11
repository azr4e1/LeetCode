package leetcode

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	idx := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[i-2] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
