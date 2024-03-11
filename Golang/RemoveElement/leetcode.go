package leetcode

func RemoveElement(nums []int, val int) int {
	var offset, length int = 0, len(nums)
	if length == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			offset++
			length--
			continue
		}
		nums[i-offset] = nums[i]
	}
	return length
}
