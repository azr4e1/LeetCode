package leetcode

const MAP_LENGTH = 256

func LengthOfLongestSubstring(substring string) int {
	currentSubstringHashMap := [MAP_LENGTH]bool{}
	currentSubstringHashMapPositions := [MAP_LENGTH]int{}
	sequenceLengthMax := 0
	sequenceLengthCurrent := 0
	substringLength := len(substring)
	for currentSymbolIndex := 0; currentSymbolIndex < substringLength; currentSymbolIndex++ {
		symbol := substring[currentSymbolIndex]
		isSymbolExists := currentSubstringHashMap[symbol]
		if isSymbolExists {
			if sequenceLengthCurrent > sequenceLengthMax {
				sequenceLengthMax = sequenceLengthCurrent
			}
			currentSymbolIndex = currentSubstringHashMapPositions[symbol]

			sequenceLengthCurrent = 0
			currentSubstringHashMap = [MAP_LENGTH]bool{}
			continue
		}
		currentSubstringHashMap[symbol] = true
		currentSubstringHashMapPositions[symbol] = currentSymbolIndex

		sequenceLengthCurrent++
		if sequenceLengthCurrent > sequenceLengthMax {
			sequenceLengthMax = sequenceLengthCurrent
		}
	}
	return sequenceLengthMax
}
