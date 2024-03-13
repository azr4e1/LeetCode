package leetcode

func Rotate1(nums []int, k int) {
	copyNums := make([]int, len(nums))
	for i, v := range nums {
		copyNums[(i+k)%len(nums)] = v
	}
	copy(nums, copyNums)
}

func Rotate2(nums []int, k int) {
	k %= len(nums)
	copyNums := make([]int, k)
	copy(copyNums, nums[len(nums)-k:])
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = copyNums[i]
	}
}

func Rotate3(nums []int, k int) {
	l := len(nums)
	k %= l
	var prev int
	for i := 0; i < k*l; i++ {
		// prev = nums[(i+1)%l]
		if i%l == 0 {
			prev = nums[l-1]
		}
		prev, nums[i%l] = nums[i%l], prev
	}
}
