package main

func findSpecialInteger(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var curr, count int
	for i := 0; i < len(arr); i++ {
		if curr == arr[i] {
			count++
		} else {
			curr = arr[i]
			count = 1
		}

		if count > len(arr)/4 {
			return curr
		}
	}

	return 0
}
