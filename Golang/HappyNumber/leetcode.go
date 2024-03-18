// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
package leetcode

func IsHappy(n int) bool {
	if n == 1 {
		return true
	}
	seen := map[int]bool{n: true}
	for {
		n = replaceBySquareSum(n)
		if _, ok := seen[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}
		seen[n] = true
	}
}

func replaceBySquareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += digit * digit
	}

	return sum
}
