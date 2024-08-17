package main

func minWindow(s string, t string) string {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	lenS := 1000000
	lS, rS := 0, 0

	for _, c := range t {
		m2[string(c)]++
	}

	l, r := 0, 0
	m1[string(s[r])]++

	for r < len(s) {
		if checkFreqForMap(m1, m2) {
			if (r-l)+1 < lenS {
				lS = l
				rS = r
				lenS = (r - l) + 1
			}
			m1[string(s[l])]--
			l++
		} else {
			r++
			if r < len(s) {
				m1[string(s[r])]++
			}
		}
	}

	if lenS == 1000000 {
		return ""
	}

	res := ""
	for i := lS; i <= rS; i++ {
		res += string(s[i])
	}
	return res
}

func checkFreqForMap(m1, m2 map[string]int) bool {
	for k, v := range m2 {
		if v > m1[k] {
			return false
		}
	}

	return true
}
