package main

import "fmt"

func TimeConvert(num int) string {
	hour := num / 60
	minute := num % 60
	return fmt.Sprintf("%d:%d", hour, minute)
}

func main() {
	fmt.Println(TimeConvert(126))
}
