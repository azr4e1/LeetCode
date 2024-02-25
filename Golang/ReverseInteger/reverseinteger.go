package reverseinteger

import "math"

// Given a signed 32-bit integer x, return x with its digits reversed. If
// reversing x causes the value to go outside the signed 32-bit integer range
// [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or
// unsigned).
func Reverse(x int) int {
	if x == 0 {
		return 0
	}
	var sign int = 1
	var reversed int

	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {

		rest := x % 10
		if (math.MaxInt32-rest)/10-reversed < 0 {
			return 0
		}
		if (math.MinInt32+rest)/10+reversed > 0 {
			return 0
		}
		reversed = 10*reversed + rest
		x /= 10
	}

	return sign * reversed
}
