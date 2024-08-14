package main

func searchMatrix(matrix [][]int, target int) bool {
	s, e := 0, len(matrix)-1

	for s <= e {
		xmid := (s + e) / 2

		if matrix[xmid][0] > target {
			e = xmid - 1
		} else if matrix[xmid][0] < target {
			l, r := 0, len(matrix[0])-1

			for l <= r {
				ymid := (l + r) / 2

				if matrix[xmid][ymid] > target {
					r = ymid - 1
				} else if matrix[xmid][ymid] < target {
					l = ymid + 1
				} else if matrix[xmid][ymid] == target {
					return true
				}
			}
			s++
		} else {
			return true
		}
	}

	return false
}
