package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	temp := 0

	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			temp++
		}
	}
	var res float64 = float64(temp) / float64(len(s)) * 100

	return int(res)
}

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
}
