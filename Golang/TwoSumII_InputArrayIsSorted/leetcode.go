package leetcode

func TwoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum < target {
			start++
			continue
		}
		if sum > target {
			end--
		}
	}
	return nil
}
