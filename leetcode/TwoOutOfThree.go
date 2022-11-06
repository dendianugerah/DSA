package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	one, two, three := make(map[int]int), make(map[int]int), make(map[int]int)
	result := []int{}

	for _, val := range nums1 {
		one[val] = 1
	}
	for _, val := range nums2 {
		two[val] = 1
	}
	for _, val := range nums3 {
		three[val] = 1
	}

	for i := 1; i <= 100; i++ {
		if one[i]+two[i]+three[i] > 1 {
			result = append(result, i)
		}
	}

	return result

}

func main() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
}
