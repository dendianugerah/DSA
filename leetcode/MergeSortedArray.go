package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)

	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				temp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = temp
			}
		}
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{5, 4, 3, 4}, 2, []int{1, 2}, 2)
}
