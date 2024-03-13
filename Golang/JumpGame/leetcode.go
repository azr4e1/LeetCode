package leetcode

func CanJump(jumps []int) bool {
	jump := jumps[0]
	if jump >= len(jumps)-1 {
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
