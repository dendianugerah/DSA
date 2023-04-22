package main

import (
	"fmt"
	"strconv"
)

func Helper(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	temp := Helper(1041)
	sl := []string{}

	for _, v := range temp {
		sl = append(sl, string(v))
	}

	fmt.Println(sl)
	fmt.Println(len(sl))

}
