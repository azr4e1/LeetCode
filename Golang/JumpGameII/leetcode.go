package leetcode

func jumpCache(nums []int, cache *map[int]int) int {
	if len(nums) == 1 {
		return 0
	}
	if val, ok := (*cache)[len(nums)]; ok {
		return val
	}
	firstJump := nums[0]
	if firstJump >= len(nums)-1 {
		(*cache)[len(nums)] = 1
		return 1
	}
	var minJumps int = len(nums)
	for i := 1; i < len(nums) && i <= firstJump; i++ {
		result := jumpCache(nums[i:], cache) + 1
		if result < minJumps {
			minJumps = result
		}
	}
	(*cache)[len(nums)] = minJumps
	return minJumps
}

func Jump(nums []int) int {
	cache := make(map[int]int)
	return jumpCache(nums, &cache)
}
