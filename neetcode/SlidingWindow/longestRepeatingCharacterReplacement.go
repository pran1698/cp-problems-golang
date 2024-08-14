package main

func characterReplacement(s string, k int) int {
	m := make(map[string]int)
	l, r := 0, 0
	res := 0

	c := string(s[r])
	m[c]++

	for r < len(s) {
		lS := (r - l) + 1

		if (lS - maxF(m)) > k {
			c1 := string(s[l])
			m[c1]--
			l++
		} else {
			res = max(res, lS)
			r++
			if r < len(s) {
				c := string(s[r])
				m[c]++
			}
		}
	}

	return res
}

func maxF(m map[string]int) int {
	mF := 0

	for _, v := range m {
		mF = max(mF, v)
	}

	return mF
}
