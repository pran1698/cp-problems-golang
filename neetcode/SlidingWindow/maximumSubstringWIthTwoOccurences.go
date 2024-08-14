package main

func maximumLengthSubstring(s string) int {
	m := make(map[string]int)
	l, r := 0, 0
	maxL := 0
	check := false

	for r < len(s) {
		if !check {
			c := string(s[r])
			m[c]++
		}

		if checkFreq(m) {
			maxL = max(maxL, (r-l)+1)
			r++
			check = false
		} else {
			c1 := string(s[l])
			m[c1] = m[c1] - 1
			l++
			check = true
		}
	}

	return maxL
}

func checkFreq(m map[string]int) bool {
	for _, v := range m {
		if v > 2 {
			return false
		}
	}

	return true
}
