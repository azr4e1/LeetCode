package leetcode

import "bytes"

func IsPalindrome(s string) bool {
	newS := bytes.ToLower([]byte(s))
	var start int
	var end int = len(newS) - 1
	for start < end {
		startEl := newS[start]
		endEl := newS[end]
		if isSpecial(startEl) {
			start++
			continue
		}
		if isSpecial(endEl) {
			end--
			continue
		}
		if startEl != endEl {
			return false
		}
		start++
		end--
	}
	return true
}

func isSpecial(b byte) bool {
	if (b < '0' || b > '9') && (b < 'a' || b > 'z') {
		return true
	}
	return false
}
