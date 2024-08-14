package main

func isAnagram(s string, t string) bool {
	c1Map := make(map[int]int)
	c2Map := make(map[int]int)

	for _, v := range s {
		c1Map[int(v)]++
	}

	for _, v := range t {
		c2Map[int(v)]++
	}

	for k, v := range c1Map {
		if c2Map[k] != v {
			return false
		}
	}

	for k, v := range c2Map {
		if c1Map[k] != v {
			return false
		}
	}

	return true
}

func isAnagramArr(s, t string) bool {

	arr := make([]int, 1000)

	for _, v := range s {
		arr[int(v)]++
	}

	for _, v := range t {
		arr[int(v)]--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true

}
