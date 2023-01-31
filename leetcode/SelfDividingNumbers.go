package main

func selfDividingNumbers(left int, right int) []int {
	result := []int{}

	for i := left; i <= right; i++ {
		digit := getDigits(i)
		count := 0

		for _, value := range digit {
			if value != 0 && i%value == 0 {
				count++
			}
		}

		if count == len(digit) {
			result = append(result, i)
		}
	}

	return result
}

func getDigits(number int) []int {
	var result []int

	for number > 0 {
		temp := number % 10
		number /= 10
		result = append(result, temp)
	}
	return result
}
