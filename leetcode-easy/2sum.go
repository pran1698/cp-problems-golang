package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		m[v] = i
	}

	var res []int
	for i, v := range nums {
		mv, ok := m[target-v]
		if ok && (i != mv) {
			res = append(res, []int{i, mv}...)
			return res
		}
	}

	return res
}
