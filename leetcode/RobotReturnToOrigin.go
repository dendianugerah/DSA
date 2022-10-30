package main

import "fmt"

func judgeCircle(moves string) bool {
	var x, y int
	for _, val := range moves {
		if val == 'U' {
			x++
		} else if val == 'D' {
			x--
		} else if val == 'L' {
			y++
		} else {
			y--
		}
	}

	return x == 0 && y == 0
}

func main() {
	fmt.Print(judgeCircle("UD"))
}
