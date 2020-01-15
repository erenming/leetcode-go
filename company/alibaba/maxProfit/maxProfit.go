package maxProfit

import (
	"math"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfits := 0
	for i := 1; i < n; i++ {
		maxProfits = int(math.Max(float64(maxProfits), float64(prices[i]-minPrice)))
		minPrice = int(math.Min(float64(minPrice), float64(prices[i])))
	}
	return maxProfits
}
