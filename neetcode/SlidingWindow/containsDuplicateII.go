package main

func containsNearbyDuplicate(nums []int, k int) bool {
	l, r := 0, 0
	m := make(map[int]int)

	for r < len(nums) {
		if (r - l) > k {
			m[nums[l]]--
			l++
		} else {
			if l != r && m[nums[r]] > 0 {
				return true
			}
			m[nums[r]]++
			r++
		}
	}

	return false
}
