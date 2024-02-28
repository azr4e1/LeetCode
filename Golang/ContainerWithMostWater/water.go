package waterarea

import (
	"math"
)

func MaxArea(height []int) int {
	maxContainer := 0
	var curr int
	for index, el := range height {
		for index2, el2 := range height[index+1:] {
			curr = Area(el, el2, index2+1)
			if curr > maxContainer {
				maxContainer = curr
			}
		}
	}
	return maxContainer
}

func Area(Side1, Side2, Dist int) int {
	minHeight := int(math.Min(float64(Side1), float64(Side2)))

	return minHeight * Dist
}
