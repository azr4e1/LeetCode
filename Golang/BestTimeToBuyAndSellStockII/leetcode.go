package leetcode

func MaxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var maxProfit int
	buy := prices[0]
	for _, v := range prices[1:] {
		if v > buy {
			maxProfit += v - buy
		}
		buy = v
	}
	return maxProfit
}
