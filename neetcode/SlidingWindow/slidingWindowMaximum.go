package main

type deq []int

func (d *deq) Len() int {
	return len(*d)
}

func (d *deq) PushBack(n int) {
	*d = append(*d, n)
}

func (d *deq) PopFront() int {
	old := *d
	n := len(old)
	if n == 0 {
		return -1 // or handle underflow as needed
	}
	i := old[0]
	*d = old[1:]
	return i
}

func (d *deq) PopBack() int {
	old := *d
	n := len(old)
	if n == 0 {
		return -1 // or handle underflow as needed
	}
	i := old[n-1]
	*d = old[:n-1]
	return i
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	n := len(nums)
	if n == 0 || k == 0 {
		return res
	}

	var q deq
	for i := 0; i < k; i++ {
		for q.Len() > 0 && nums[q[q.Len()-1]] < nums[i] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	res = append(res, nums[q[0]])

	for i := k; i < n; i++ {
		for q.Len() > 0 && nums[q[q.Len()-1]] < nums[i] {
			q.PopBack()
		}
		q.PushBack(i)

		if q[0] <= i-k {
			q.PopFront()
		}
		res = append(res, nums[q[0]])
	}

	return res
}
