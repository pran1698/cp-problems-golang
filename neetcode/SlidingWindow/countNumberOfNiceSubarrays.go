package main

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	l, r := 0, 0
	odds := 0

	for r < len(nums) {
		if nums[r]%2 == 0 {
			r++
		} else {
			r++
			odds++
		}

		for odds > k {
			if nums[l]%2 == 0 {
				l++
			} else {
				odds--
				l++
			}
		}

		tempOdds := odds
		tempL := l
		for tempOdds == k {
			if nums[tempL]%2 == 0 {
				res++
			} else {
				tempOdds--
				res++
			}
			tempL++
		}
	}

	return res
}
