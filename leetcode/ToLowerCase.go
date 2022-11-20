package main

import "fmt"

func toLowerCase(s string) string {
	res := []rune{}
	for _, v := range s {
		if v >= 65 && v <= 90 {
			res = append(res, v+32)
			continue
		}
		res = append(res, v)
	}

	return string(res)
}

func main() {
	fmt.Println(toLowerCase("HDWdosiaJD"))
}
