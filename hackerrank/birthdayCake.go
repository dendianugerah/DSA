package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	max := candles[0]
	var result int32
	for i := 0; i < len(candles); i++ {
		if candles[i] > max {
			max = candles[i]
		}
	}

	for _, v := range candles {
		if max == v {
			result++
		}
	}
	return result
}
