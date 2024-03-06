package leetcode

func GenerateParenthesis(n int) []string {
	return genParHelper(n, 0)

}

func genParHelper(n int, openPar int) []string {
	if n == 0 {
		if openPar > 0 {
			finalPar := []byte{}
			for i := 0; i < openPar; i++ {
				finalPar = append(finalPar, ')')
			}
			return []string{string(finalPar)}
		}
		return []string{}
	}

	result := []string{}

	if openPar > 0 {
		partialResult := genParHelper(n, openPar-1)
		for _, res := range partialResult {
			result = append(result, ")"+res)
		}
	}
	partialResult := genParHelper(n-1, openPar+1)
	for _, res := range partialResult {
		result = append(result, "("+res)
	}

	return result
}
