// Given a string s consisting of words and spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.
package leetcode

func lengthOfLastWord(s string) int {
	var counter int
	var inWord bool
	for i := 0; i < len(s); i++ {
		curr := s[i]
		if curr == ' ' || curr == '\n' || curr == '\t' {
			inWord = false
			continue
		} else {
			if !inWord {
				counter = 0
				inWord = true
			}
		}
		counter++
	}
	return counter
}
