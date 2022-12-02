package main

func halvesAreAlike(s string) bool {
	sMid := len(s) / 2
	temp := 0

	for i, v := range s {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < sMid {
				temp++
			} else {
				temp--
			}
		}
	}

	return temp == 0
}
