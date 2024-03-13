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
		if i%l == 0 {
			prev = nums[l-1]
		}
		prev, nums[i%l] = nums[i%l], prev
	}
}

func Rotate4(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
