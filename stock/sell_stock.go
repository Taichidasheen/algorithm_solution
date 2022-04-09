package stock

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
 */
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minprice := prices[0]
	maxprofit := 0
	for _, price := range prices {
		if price < minprice {
			minprice = price
		} else {
			profit := price - minprice
			if profit > maxprofit {
				maxprofit = profit
			}
		}
	}
	return maxprofit
}
