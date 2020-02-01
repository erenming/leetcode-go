package maxProfit122

func maxProfit(prices []int) int {
	profit, n := 0, len(prices)
	for i := 0; i < n - 1; i ++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
