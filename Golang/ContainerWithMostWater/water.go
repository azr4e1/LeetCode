package waterarea

func calculateArea(low, high, first, second int) int {
	if second > first {
		return first * (high - low)
	} else {
		return second * (high - low)
	}
}

func MaxArea(height []int) int {
	low_pointer := 0
	high_pointer := len(height) - 1
	max_area := 0
	for low_pointer < high_pointer {
		newSum := calculateArea(low_pointer, high_pointer, height[low_pointer], height[high_pointer])
		if newSum > max_area {
			max_area = newSum
		}
		if height[low_pointer] < height[high_pointer] {
			low_pointer++
		} else {
			high_pointer--
		}

	}
	return max_area
}
