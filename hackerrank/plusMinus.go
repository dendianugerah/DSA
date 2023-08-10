package main

func plusMinus(arr []int32) {
	var zero, positive, negative float64

	for _, v := range arr {
		if v == 0 {
			zero++
		} else if v > 0 {
			positive++
		} else {
			negative++
		}
	}

	positive = positive / float64(len(arr))
	negative = negative / float64(len(arr))
	zero = zero / float64(len(arr))

	println(positive)
	println(negative)
	println(zero)
}
