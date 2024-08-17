package main

import (
	"reflect"
)

func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	m1 := make(map[string]int)
	for _, c := range s1 {
		m1[string(c)]++
	}
	m2 := make(map[string]int)

	l, r := 0, 0
	check := true
	for r < n2 {
		if check {
			c := string(s2[r])
			m2[c]++
		}

		if (r-l)+1 > n1 {
			c := string(s2[l])
			m2[c]--
			if m2[c] == 0 {
				delete(m2, c)
			}
			l++
			check = false
		} else {
			if reflect.DeepEqual(m1, m2) {
				return true
			}
			r++
			check = true
		}
	}

	return false
}
