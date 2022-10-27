package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		reverse(row)
		invert(row)
	}
	return image
}

func reverse(row []int) []int {
	for i := 0; i < len(row)/2; i++ {
		temp := row[i]
		row[i] = row[len(row)-i-1]
		row[len(row)-i-1] = temp
	}
	return row
}

func invert(row []int) []int {
	for i := 0; i < len(row); i++ {
		if row[i] == 1 {
			row[i] = 0
		} else {
			row[i] = 1
		}
	}
	return row
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
