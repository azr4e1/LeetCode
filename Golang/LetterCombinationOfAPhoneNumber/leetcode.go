// Given a string containing digits from 2-9 inclusive, return all possible letter
// combinations that the number could represent. Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons) is given
// below. Note that 1 does not map to any letters.
package leetcode

var Map = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	result := []string{}
	if digits == "" {
		return result
	}
	digit := digits[0]
	letters := Map[digit]
	if digits == string(digit) {
		for _, letter := range letters {
			result = append(result, string(letter))
		}
		return result
	}

	partialResults := LetterCombinations(digits[1:])
	for _, letter := range letters {
		for _, el := range partialResults {
			result = append(result, string(letter)+el)
		}
	}
	return result
}
