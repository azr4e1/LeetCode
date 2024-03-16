// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
package leetcode

func StrStr(haystack string, needle string) int {
	var index int = -1
outer:
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			index = i
			for j, v := range needle[1:] {
				if i+j+1 >= len(haystack) {
					index = -1
					continue outer
				}
				if rune(haystack[i+j+1]) != v {
					index = -1
					continue outer
				}
			}
			return index
		}
	}
	return index
}

// better version
func strStr(h string, n string) int {
	l := len(h)
	m := len(n)

	if m > l {
		return -1
	}

	if h == n || m == 0 {
		return 0
	}

	for i := 0; i <= l-m; i++ {
		if h[i:i+m] == n {
			return i
		}
	}
	return -1
}
