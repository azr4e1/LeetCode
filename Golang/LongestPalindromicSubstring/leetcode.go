// Given a string s, return the longest palindromic substring in s.
package leetcode

import "strings"

// func LongestPalindrome(s string) string {
// 	if len(s) <= 1 {
// 		return s
// 	}
// 	var substring string
// 	var reverseSubstring string
// 	var longestPal string
// 	for i := 0; i < len(s); i++ {
// 		substring = string(s[i])
// 		reverseSubstring = substring
// 		for j := i + 1; j < len(s); j++ {
// 			substring = substring + string(s[j])
// 			reverseSubstring = string(s[j]) + reverseSubstring
// 			if substring == reverseSubstring && len(substring) > len(longestPal) {
// 				longestPal = substring
// 			}
// 		}
// 	}

// 	if longestPal == "" {
// 		longestPal = string(s[0])
// 	}
// 	return longestPal
// }

// Manacher's algorithm (not mine)

// Approach

// Manacher's Algorithm is a powerful technique that allows us to find the longest palindromic substring in a given string in linear time. Here's a detailed breakdown of the algorithm's approach:
// 1. String Transformation

// We first transform the original string to simplify the algorithm. This transformation achieves two things:

//     It ensures that every potential center of a palindrome is surrounded by identical characters (#), which simplifies the process of expanding around a center.
//     It adds special characters ^ at the beginning and $ at the end of the string to avoid any boundary checks during the palindrome expansion. For instance, the string "babad" is transformed into "^#b#a#b#a#d#$".

// 2. Initialization

// We maintain an array P with the same length as the transformed string. Each entry P[i] denotes the radius (half-length) of the palindrome centered at position i.

// We also introduce two critical pointers:

//     C: The center of the palindrome that has the rightmost boundary.
//     R: The right boundary of this palindrome.

// Both C and R start at the beginning of the string.
// 3. Iterating Through the String

// For every character in the transformed string, we consider it as a potential center for a palindrome.

// a. Using Previously Computed Information:
// If the current position is to the left of R, its mirror position about the center C might have information about a palindrome centered at the current position. We can leverage this to avoid unnecessary calculations.

// b. Expanding Around the Center:
// Starting from the current radius at position i (which might be derived from its mirror or initialized to 0), we attempt to expand around i and check if the characters are the same.

// c. Updating C and R:
// If the palindrome centered at i extends beyond R, we update C to i and R to the new boundary.
// 4. Extracting the Result

// Once we've computed the palindromic radii for all positions in the transformed string, we find the position with the largest radius in P. This position represents the center of the longest palindromic substring. We then extract and return this palindrome from the original string.
// Complexity

//     Time complexity: O(n)O(n)O(n)
//     Manacher's algorithm processes each character in the transformed string once, making the time complexity linear.

// Space complexity: O(n)O(n)O(n)
// We use an array P to store the palindrome radii, making the space complexity linear as well.
func LongestPalindrome(s string) string {
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)
	P := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		if R > i {
			P[i] = min(R-i, P[2*C-i])
		}
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
