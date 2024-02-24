package palindrome

import (
	"errors"
	"fmt"
	"math"
)

// Given an integer x, return true if x is a palindrome , and false otherwise.
func IsPalindromeV1(x int) bool {
	if x < 0 {
		return false
	}
	decomposition := Decompose(x)
	length := len(decomposition)
	for i := 0; i < length; i++ {
		if decomposition[i] != decomposition[length-1-i] {
			return false
		}
	}

	return true
}

func IsPalindrome(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	if x%10 == 0 || x < 0 {
		return false
	}
	inverted := 0
	original := x

	for x > 0 {
		curr := x % 10
		inverted = 10*inverted + curr
		x = (x - curr) / 10
	}

	return original == inverted
}

func Decompose(x int) []int {
	if x == 0 {
		return []int{0}
	}

	result := []int{}
	for x > 0 {
		curr := x % 10
		result = append(result, curr)
		x = (x - curr) / 10
	}
	return result
}

func Compose(s []int) (int, error) {
	var result int
	for i, el := range s {
		if el >= 10 || el < 0 {
			return 0, errors.New("Invalid element in slice")
		}
		result += int(math.Pow(10, float64(i))) * el
	}

	return result, nil
}
