package main

import "fmt"

type location struct {
	x int
	y int
}

func isPathCrossing(path string) bool {
	x, y := 0, 0
	tracks := make(map[location]bool)
	tracks[location{x, y}] = true

	for _, p := range path {
		switch p {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		if tracks[location{x, y}] {
			return true
		}

		tracks[location{x, y}] = true
	}

	return false
}

/*
Failed at 59/81 Test Case

func isPositive(num int) bool {
	if num < 0 {
		return false
	}

	return true
}

func isPathCrossing(path string) bool {
	l := 0
	pos := 0
	neg := 0

	for _, v := range path {
		switch v {
		case 'N':
			l++
		case 'E':
			l++
		default:
			l--
		}

		if isPositive(l) {
			pos++
		} else {
			neg--
		}
	}

	if neg < 0 && pos > 0 {
		return true
	}

	return false
}
*/

func main() {
	fmt.Println(isPathCrossing("NES"))
}
