// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
package leetcode

import "strings"

func WordPattern(pattern string, s string) bool {
	var dict = make(map[byte]string)
	var chars = make(map[string]byte)
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		word, ok := dict[pattern[i]]
		char, ok2 := chars[words[i]]
		switch {
		case !ok && !ok2:
			dict[pattern[i]] = words[i]
			chars[words[i]] = pattern[i]
		case !ok && ok2:
			return false
		case ok && !ok2:
			return false
		default:
			if word != words[i] || char != pattern[i] {
				return false
			}
		}
	}
	return true
}
