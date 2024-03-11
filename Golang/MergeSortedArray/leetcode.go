package leetcode

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
func Merge(nums1, nums2 []int, m, n int) {
	helper := make([]int, m+n)
	var i, j int
	for k := 0; i < m || j < n; k++ {
		if i == m {
			helper[k] = nums2[j]
			j++
			continue
		}
		if j == n {
			helper[k] = nums1[i]
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			helper[k] = nums1[i]
			i++
			continue
		}
		helper[k] = nums2[j]
		j++
	}

	for k := 0; k < m+n; k++ {
		nums1[k] = helper[k]
	}
}
