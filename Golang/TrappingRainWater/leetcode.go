// Given n non-negative integers representing an elevation map where the width of
// each bar is 1, compute how much water it can trap after raining.
package leetcode

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

func trap(height []int) int {
	maxHeight := slices.Max(height)
	waterPool := 0
	for i := 0; i < maxHeight; i++ {
		currentLevelWater := 0
		start, end := 0, len(height)-1
		for start < end {
			if height[start]-i > 0 {
				break
			} else {
				start++
			}
		}
		for start < end {
			if height[end]-i > 0 {
				break
			} else {
				end--
			}
		}
		if start == end {
			break
		}
		currentLevelWater += end - start - 1
		for j := start + 1; j < end; j++ {
			if height[j]-i > 0 {
				currentLevelWater--
			}
		}
		waterPool += currentLevelWater
	}
	return waterPool
}

func Print(height []int, out io.Writer) {
	maxHeight := slices.Max(height)
	output := new(strings.Builder)
	for i := maxHeight - 1; i >= 0; i-- {
		for j := 0; j < len(height); j++ {
			if diff := height[j] - i; diff > 0 {
				output.WriteString("#")
			} else {
				output.WriteString(" ")
			}
		}
		output.WriteString("\n")
	}
	fmt.Fprint(out, output.String())
}
