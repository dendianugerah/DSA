package main

import (
	"fmt"
	"sort"
)

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	i := 0
	for j := 0; j < len(trainers); j++ {
		if players[i] <= trainers[j] {
			i++
			if i == len(players) {
				return i
			}
		}
	}

	return i
}

func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))
}
