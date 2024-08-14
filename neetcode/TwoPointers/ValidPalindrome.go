package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1

	for i < j {
		if !isAlphaNumeric(s[i]) {
			i++
		} else if !isAlphaNumeric(s[j]) {
			j--
		} else if rune(s[i]-'0') == ' ' {
			i++
		} else if rune(s[j]-'0') == ' ' {
			j--
		} else {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}

	return true
}

func isAlphaNumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9')
}
