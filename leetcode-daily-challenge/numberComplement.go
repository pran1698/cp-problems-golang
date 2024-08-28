package main

func findComplement(num int) int {
	i := 0
	pow := 1
	for i < 31 {
		if pow > num {
			break
		}
		i++
		pow = pow * 2
	}

	return (pow - 1) ^ num
}
