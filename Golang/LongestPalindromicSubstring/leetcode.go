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
