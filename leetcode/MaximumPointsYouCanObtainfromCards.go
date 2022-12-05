package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	var sum int

	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		sum += cardPoints[i]
	}

	max := sum
	left, right := 0, len(cardPoints)-k

	for right < len(cardPoints) {
		sum += cardPoints[left]
		left++
		sum -= cardPoints[right]
		right++
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}
