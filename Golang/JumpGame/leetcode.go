package leetcode

func CanJump(jumps []int) bool {
	jump := jumps[0]
	if len(jumps) == 1 {
		return true
	}
	if jump == 0 {
		return false
	}
	if jump >= len(jumps) {
		return true
	}
	var result bool
	for i := 1; i <= jump && i < len(jumps); i++ {
		result = CanJump(jumps[i:])
		if result {
			return true
		}
	}
	return false
}
