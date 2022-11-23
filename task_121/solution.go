package task_121

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		} else if profit < price-minPrice {
			profit = price - minPrice
		}
	}
	return profit
}
