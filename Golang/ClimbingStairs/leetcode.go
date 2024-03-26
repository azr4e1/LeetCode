// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
package leetcode

var memoize = make(map[int]int)

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	val, ok := memoize[n]
	if ok {
		return val
	}

	var ways int
	ways += climbStairs(n - 1)
	ways += climbStairs(n - 2)
	memoize[n] = ways

	return ways
}
