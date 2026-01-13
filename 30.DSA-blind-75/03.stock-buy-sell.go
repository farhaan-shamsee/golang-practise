/*
Problem: Best Time to Buy and Sell Stock
Given an array of stock prices where prices[i] is the price on day i,
find the maximum profit you can achieve by buying on one day and selling on another later day.

Approach:
- Track the minimum price seen so far (best day to buy)
- For each day, calculate profit if we sell today (current price - min price)
- Keep updating max profit if we find a better one
- Update minimum price if current day has lower price

Algorithm:
1. Initialize maxProfit = 0 and leastPrice = prices[0]
2. Loop through prices starting from day 1
3. Calculate profit = priceToday - leastPrice
4. If profit > maxProfit, update maxProfit
5. If priceToday < leastPrice, update leastPrice (found cheaper buy day)
6. Return maxProfit

Time Complexity: O(n) - single pass through array
Space Complexity: O(1) - only using two variables

Example:
prices = [7,1,5,3,6,4]
Day 0: price=7, leastPrice=7, maxProfit=0
Day 1: price=1, profit=1-7=-6, leastPrice=1, maxProfit=0
Day 2: price=5, profit=5-1=4, maxProfit=4, leastPrice=1
Day 3: price=3, profit=3-1=2, maxProfit=4, leastPrice=1
Day 4: price=6, profit=6-1=5, maxProfit=5, leastPrice=1
Day 5: price=4, profit=4-1=3, maxProfit=5, leastPrice=1
Result: maxProfit=5 (buy at 1, sell at 6)
*/
package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	leastPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		priceToday := prices[i]

		profit := priceToday - leastPrice
		if profit > maxProfit {
			maxProfit = profit
		}

		if leastPrice > priceToday {
			leastPrice = priceToday
		}

	}

	return maxProfit
}

func stockBuySell() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Printf("Maximum profit for prices %v: %d\n", prices, profit)

	prices2 := []int{7, 6, 4, 3, 1}
	profit2 := maxProfit(prices2)
	fmt.Printf("Maximum profit for prices %v: %d\n", prices2, profit2)
}
