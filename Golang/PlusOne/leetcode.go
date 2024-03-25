package leetcode

// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] <= 9 {
			return digits
		}
		digits[i] = 0
	}

	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}
