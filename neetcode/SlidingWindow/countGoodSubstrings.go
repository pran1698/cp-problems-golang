package main

func countGoodSubstrings(s string) int {
	m := make(map[string]bool)

	l, r := 0, 0
	res := 0

	for r < len(s) {
		c := string(s[r])
		if m[c] || ((r-l)+1 > 3) {
			c2 := string(s[l])
			m[c2] = false
			l++
		} else {
			m[c] = true
			if (r-l)+1 == 3 {
				res++
			}
			r++
		}
	}

	return res
}
