package leetcode

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

// You must not use any built-in exponent function or operator.
func mySqrt(x int) int {
	var i int
	for i = 0; i*i <= x; i++ {
	}
	return i
}
