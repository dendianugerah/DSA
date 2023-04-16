package main

func passThePillow(n int, time int) int {
	step := n - 1
	time %= (step * 2)

	if time <= step {
		return time + 1
	} else {
		return step*2 - time + 1
	}
}
