package main

import "strconv"

var stack1 []int

func Push1(i int) {
	stack1 = append(stack1, i)
}

func Pop1() int {
	item := stack1[len(stack)-1]
	stack1 = stack1[:len(stack)-1]
	return item
}

func Len1() int {
	return len(stack1)
}

func evalRPN(tokens []string) int {
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			num1 := Pop1()
			num2 := Pop1()

			if token == "+" {
				res := num2 + num1
				Push1(res)
			} else if token == "-" {
				res := num2 - num1
				Push1(res)
			} else if token == "*" {
				res := num2 * num1
				Push1(res)
			} else if token == "/" {
				res := num2 / num1
				Push1(res)
			}

		} else {
			num, _ := strconv.Atoi(token)
			Push1(num)
		}
	}

	return Pop1()
}
