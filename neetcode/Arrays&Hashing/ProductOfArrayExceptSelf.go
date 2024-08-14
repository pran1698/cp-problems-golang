package main

func productExceptSelfDivide(nums []int) []int {
	productN := 1
	productZ := 1
	nZeros := 0
	for _, v := range nums {
		if v != 0 {
			productN *= v
		} else {
			productZ = 0
			nZeros++
		}
	}

	var ans []int
	for _, v := range nums {
		if productZ == 1 && v != 0 {
			ans = append(ans, productN/v)
		} else if productZ == 0 && v != 0 {
			ans = append(ans, 0)
		} else if v == 0 {
			if nZeros < 2 {
				ans = append(ans, productN)
			} else {
				ans = append(ans, 0)
			}
		}
	}

	return ans
}

func productExceptSelf2Arr(nums []int) []int {
	p := make([]int, len(nums))
	s := make([]int, len(nums))

	ans := make([]int, len(nums))

	n := len(nums)
	p[0] = 1
	s[n-1] = 1

	for i := 1; i < n; i++ {
		p[i] = p[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		s[i] = s[i+1] * nums[i+1]
	}

	ans[0] = s[0]
	ans[n-1] = p[n-1]
	for i := 1; i < n-1; i++ {
		ans[i] = s[i] * p[i]
	}

	return ans
}

func productExceptSelfSingleArr(nums []int) []int {
	p := make([]int, len(nums))
	s := make([]int, len(nums))

	n := len(nums)
	p[0] = 1
	s[n-1] = 1

	for i := 1; i < n; i++ {
		p[i] = p[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		s[i] = s[i+1] * nums[i+1]
	}

	for i := 1; i < n-1; i++ {
		s[i] = s[i] * p[i]
	}

	s[n-1] = p[n-1]
	return s
}
