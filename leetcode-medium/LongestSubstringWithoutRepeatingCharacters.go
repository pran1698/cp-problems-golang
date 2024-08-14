package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	arr := make([]string, len(s))

	if len(s) == 0 {
		return 0
	}

	arr[0] = fmt.Sprintf("%c", s[0])
	for i := 1; i < len(s); i++ {
		c := fmt.Sprintf("%c", s[i])
		sInd := strings.Index(arr[i-1], c)
		str := arr[i-1]

		if sInd != -1 {
			chars := str[sInd+1:]
			arr[i] = chars + c
		} else {
			arr[i] = arr[i-1] + c
		}
	}

	maxInt := 0
	for _, s := range arr {
		fmt.Println(s)
		if len(s) > maxInt {
			maxInt = len(s)
		}
	}

	return maxInt
}
