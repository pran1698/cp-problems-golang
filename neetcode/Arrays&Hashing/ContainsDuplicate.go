package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for _, v := range m {
		if v > 1 {
			return true
		}
	}

	return false
}
