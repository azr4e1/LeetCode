package leetcode

func LengthOfLongestSubstring(s string) int {
	alphabet := make(map[byte]int)
	var offset int
	var substring = make([]byte, 0, len(s))
	var maxLength int
	for i := 0; i < len(s); i++ {
		el := s[i]
		val, ok := alphabet[el]
		if ok {
			if len(substring) > maxLength {
				maxLength = len(substring)
			}
			alphabet = make(map[byte]int)
			substring = append(substring, el)
			substring = substring[val+1:]
			for j := 0; j < len(substring); j++ {
				alphabet[substring[j]] = j
			}
			offset = i - len(substring) + 1
			continue
		}
		substring = append(substring, el)
		alphabet[el] = i - offset
	}
	if len(substring) > maxLength {
		return len(substring)
	}

	return maxLength
}
