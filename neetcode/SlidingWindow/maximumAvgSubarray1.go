package main

func findMaxAverage(nums []int, k int) float64 {
	l, r := 0, 0
	sum := nums[r]
	var maxA float64 = -1000000

	for r < len(nums) {
		if (r-l)+1 > k {
			sum = sum - nums[l]
			l++
		} else {
			if (r-l)+1 == k {
				var a float64 = float64(sum) / float64(k)
				maxA = max(maxA, a)
			}
			r++
			if r < len(nums) {
				sum += nums[r]
			}
		}
	}

	return maxA
}
