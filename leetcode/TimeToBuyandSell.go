package main

import "fmt"

func maxProfit(prices []int) int {
	down := prices[0]
	up := prices[0]
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < down {
			down = prices[i]
			up = down
		}
		if prices[i] > up {
			up = prices[i]
		}
		if up-down > profit {
			profit = up - down
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
