package leetcode

func RemoveDuplicates(nums []int) int {
	var i, counter int = 1, 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			counter = 0
		}
		if counter < 2 {
			nums[i] = nums[j]
			i++
		}
		counter++
	}
	return i
}
