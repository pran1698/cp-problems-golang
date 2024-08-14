package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		a := 0 - nums[i]

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			sum := nums[j] + nums[k]
			if sum > a {
				k--
			} else if sum < a {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}

	return res
}
