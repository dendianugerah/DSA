package main

import "strconv"

func isPalindrome(x int) bool {
	xString := strconv.Itoa(x)
	l := len(xString)

	for i := 0; i < len(xString)/2; i++ {
		if xString[i] != xString[l-1] {
			return false
		}
		l--
	}

	return true
}

func main() {

}
