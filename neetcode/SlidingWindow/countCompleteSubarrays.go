package main

// will traverse through every subarray
func countCompleteSubarrays(nums []int) int {
	d := 0
	m1 := make(map[int]int)
	for _, n := range nums {
		if m1[n] == 0 {
			d++
			m1[n]++
		}
	}

	m2 := make(map[int]bool)
	l, r := 0, 0
	n := len(nums)
	res := 0

	for l < n {
		r = l
		m2 = make(map[int]bool)
		for r < n {
			if !m2[nums[r]] {
				m2[nums[r]] = true
			}

			if len(m2) == d {
				res++
			}
			r++
		}
		l++
	}

	return res
}
