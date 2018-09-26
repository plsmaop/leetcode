func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}