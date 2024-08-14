package main

import "sort"

func groupAnagramsTle(strs []string) [][]string {
	isAnagramOther := make([]bool, len(strs))

	var res [][]string

	curr := 0
	for i := 0; i < len(strs); i++ {
		if !isAnagramOther[i] {
			res = append(res, []string{strs[i]})
			isAnagramOther[i] = true

			for j := i + 1; j < len(strs); j++ {
				if validAnagram(strs[i], strs[j]) {
					res[curr] = append(res[curr], strs[j])
					isAnagramOther[j] = true
				}
			}

			curr++
		}
	}

	return res
}

func validAnagram(s, t string) bool {
	arr := make([]int, 1000)

	for _, v := range s {
		arr[int(v)]++
	}

	for _, v := range t {
		arr[int(v)]--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		chars := []rune(v)

		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		sortedStr := string(chars)
		m[sortedStr] = append(m[sortedStr], v)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
