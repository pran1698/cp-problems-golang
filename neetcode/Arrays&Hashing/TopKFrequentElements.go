package main

func topKFrequent(nums []int, k int) []int { // faster solution, 8ms  memory: 6mb
	var res []int

	m1 := make(map[int]int)
	m2 := make(map[int][]int)
	var arr []int

	for _, v := range nums {
		m1[v]++
	}

	for k, v := range m1 {
		m2[v] = append(m2[v], k)
	}

	for len(m2) != 0 {
		max := 0
		for f := range m2 {
			if f > max {
				max = f
			}
		}

		arr = append(arr, m2[max]...)

		delete(m2, max)
	}

	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}

	return res
}

func topKFrequentSingleMap(nums []int, k int) []int { // slower solution, 38ms less memory
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		mxCount := 0
		mxCandidate := 0
		for k, v := range m {
			if v > mxCount {
				mxCandidate = k
				mxCount = v
			}
		}
		res[i] = mxCandidate
		delete(m, mxCandidate)
	}

	return res

}
