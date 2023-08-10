package main

func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{}
	result = append(result, 0, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] > b[i] {
			result[0]++
		} else {
			result[1]++
		}
	}
	return result
}
