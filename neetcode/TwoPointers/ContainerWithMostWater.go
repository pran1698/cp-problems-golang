package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxArea := 0

	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if area > maxArea {
			maxArea = area
		}

		if height[j] >= height[i] {
			i++
		} else if height[j] < height[i] {
			j--
		}
	}

	return maxArea
}
