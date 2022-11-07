package main

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var res int64
	for _, v := range ar {
		res += v
	}
	return res
}
