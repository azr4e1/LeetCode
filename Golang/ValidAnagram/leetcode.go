// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
package leetcode

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var dict = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		dict[t[i]]--
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}
