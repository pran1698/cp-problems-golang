package main

func trap(height []int) int {
	l, r := 0, len(height)-1
	maxL, maxR := height[0], height[len(height)-1]
	trap := 0

	for l < r {
		maxL = max(maxL, height[l])
		maxR = max(maxR, height[r])

		if maxL <= maxR {
			water := maxL - height[l]
			if water > 0 {
				trap += water
			}
			l++
		} else {
			water := maxR - height[r]
			if water > 0 {
				trap += water
			}
			r--
		}
	}

	return trap
}
