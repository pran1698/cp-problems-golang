package main

func minimumCardPickup(cards []int) int {
	m := make(map[int]int)
	minL := 100000000

	for _, c := range cards {
		m[c] = -1
	}

	for i, c := range cards {
		if m[c] == -1 {
			m[c] = i
		} else {
			minL = min(minL, (i-m[c])+1)
			m[c] = i
		}
	}

	if minL == 100000000 {
		return -1
	}

	return minL
}
