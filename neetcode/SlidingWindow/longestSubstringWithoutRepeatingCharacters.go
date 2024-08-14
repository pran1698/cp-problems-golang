package main

func lengthOfLongestSubstring(s string) int {
	arr := make([]string, 0)

	l, r := 0, 0
	maxL := 0

	for r < len(s) {
		c := string(s[r])
		if findInSet(arr, c) {
			arr = removeFromSet(arr, l, s)
			l++
		} else {
			arr = append(arr, string(s[r]))
			maxL = max(maxL, (r-l)+1)
			r++
		}
	}

	return maxL
}

func findInSet(arr []string, s string) bool {
	for _, c := range arr {
		if c == s {
			return true
		}
	}

	return false
}

func removeFromSet(arr []string, i int, s string) []string {
	c := string(s[i])
	for i, r := range arr {
		if r == c {
			arr[i] = ""
		}
	}

	return arr
}
