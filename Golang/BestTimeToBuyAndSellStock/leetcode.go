package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var buy = prices[0]
	var maxRevenue int
	for _, v := range prices[1:] {
		if v < buy {
			buy = v
			continue
		}
		value := v - buy
		if value > maxRevenue {
			maxRevenue = value
		}
	}
	return maxRevenue
}
