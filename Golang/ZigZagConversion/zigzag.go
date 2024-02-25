package zigzag

import "strings"

func Convert(s string, numRows int) string {
	matrix := ToMatrix(s, numRows)
	toString := []string{}
	for _, el := range matrix {
		toString = append(toString, string(el))
	}

	return strings.Join(toString, "")
}

func ToMatrix(s string, numRows int) [][]byte {
	matrix := make([][]byte, numRows)
	if numRows == 1 {
		matrix[0] = []byte(s)
		return matrix
	}
	jumpChar := 2 * (numRows - 1)

	for i := 0; i < len(s); i++ {
		el := s[i]
		index := i % jumpChar
		if index > numRows-1 {
			index = jumpChar - index
		}

		row := append(matrix[index], el)
		matrix[index] = row

	}
	return matrix
}
