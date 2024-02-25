package commpref

import "math"

// Write a function to find the longest common prefix string amongst an array of st
// rings.
// If there is no common prefix, return an empty string "".
func LongestCommonPrefix(strs []string) string {
	finalStr := []rune(strs[0])
	maxLength := len(finalStr)
	for _, str := range strs {
		if string(finalStr) == "" {
			return ""
		}
		maxLength = int(math.Min(float64(maxLength), float64(len(str))))
		finalStr = finalStr[:maxLength]
		for index, r := range str {
			if index >= maxLength {
				break
			}
			if r != finalStr[index] {
				if index == 0 {
					return ""
				}
				finalStr = finalStr[:index]
				maxLength = index
				break
			}
		}
	}

	return string(finalStr)
}
