package leetcode

func canJumpCache(jumps []int, cache *map[int]bool) bool {
	if val, ok := (*cache)[len(jumps)]; ok {
		return val
	}
	jump := jumps[0]
	if jump >= len(jumps)-1 {
		(*cache)[len(jumps)] = true
		return true
	}
	var result bool
	for i := min(jump, len(jumps)-1); i >= 1; i-- {
		result = canJumpCache(jumps[i:], cache)
		if result {
			(*cache)[len(jumps)] = true
			return true
		}
	}
	(*cache)[len(jumps)] = false
	return false
}

func CanJump(jumps []int) bool {
	cache := make(map[int]bool)
	result := canJumpCache(jumps, &cache)

	return result
}
