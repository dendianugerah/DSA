package main

import (
	"fmt"
	"strconv"
)

func thousandSeparator(n int) string {
	nStr := strconv.Itoa(n)

	for i := len(nStr) - 3; i > 0; i -= 3 {
		nStr = nStr[:i] + "." + nStr[i:]
	}

	return nStr
}

func main() {
	fmt.Println(thousandSeparator(13893))
}
