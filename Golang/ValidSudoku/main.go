package main

import "fmt"

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//     Each row must contain the digits 1-9 without repetition.
//     Each column must contain the digits 1-9 without repetition.
//     Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
// Note:
//
//     A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//     Only the filled cells need to be validated according to the mentioned rules.
//

var board = [][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

func main() {
	fmt.Println(isValidSudoku(board))
}

type row []byte

type column []byte

type subBox []byte

func (r row) isValid() bool {
	return isValid(r)
}

func (c column) isValid() bool {
	return isValid(c)
}

func (s subBox) isValid() bool {
	return isValid(s)
}

func isValidSudoku(board [][]byte) bool {
	columns := make([]column, 9)
	subBoxes := make([]subBox, 3)
	for i, r := range board {
		// when completing a row of subboxes, check and re-initialize
		if i != 0 && i%3 == 0 {
			for _, sb := range subBoxes {
				if !sb.isValid() {
					return false
				}
			}
			subBoxes = make([]subBox, 3)
		}
		currRow := row(r)
		if !currRow.isValid() {
			return false
		}

		for j, c := range r {
			// create columns
			col := columns[j]
			col = append(col, c)
			columns[j] = col

			subBoxIndex := j / 3
			sb := subBoxes[subBoxIndex]
			sb = append(sb, c)
			subBoxes[subBoxIndex] = sb
		}
	}

	// last check for sub-boxes
	for _, sb := range subBoxes {
		if !sb.isValid() {
			return false
		}
	}

	// check for rows
	for _, c := range columns {
		if !c.isValid() {
			return false
		}
	}

	return true
}

func isValid(l []byte) bool {
	seen := make(map[byte]struct{})
	for _, e := range l {
		if e == '.' {
			continue
		}
		_, ok := seen[e]
		if ok {
			return false
		}
		seen[e] = struct{}{}
	}

	return true
}
