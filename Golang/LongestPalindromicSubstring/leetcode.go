// Given a string s, return the longest palindromic substring in s.
package leetcode

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var substring string
	var reverseSubstring string
	var longestPal string
	for i := 0; i < len(s); i++ {
		substring = string(s[i])
		reverseSubstring = substring
		for j := i + 1; j < len(s); j++ {
			substring = substring + string(s[j])
			reverseSubstring = string(s[j]) + reverseSubstring
			if substring == reverseSubstring && len(substring) > len(longestPal) {
				longestPal = substring
			}
		}
	}

	if longestPal == "" {
		longestPal = string(s[0])
	}
	return longestPal
}
