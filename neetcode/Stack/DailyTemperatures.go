package main

type Node struct {
	val   int
	index int
}

var stack2 []Node

func Push2(n Node) {
	stack2 = append(stack2, n)
}

func Pop2() Node {
	node := stack2[len(stack)-1]
	stack2 = stack2[:len(stack)-1]
	return node
}

func Top2() int {
	return stack2[len(stack)-1].val
}

func Len2() int {
	return len(stack)
}

func dailyTemperatures(temperatures []int) []int {
	out := make([]int, len(temperatures))
	stack2 = make([]Node, 0)

	for i, temp := range temperatures {
		if Len() > 0 {
			top := Top2()
			for temp > top && Len() > 0 {
				item := Pop2()
				out[item.index] = i - item.index
				if Len() > 0 {
					top = Top2()
				}
			}
			n := Node{
				val:   temp,
				index: i,
			}
			Push2(n)
		} else {
			n := Node{
				val:   temp,
				index: i,
			}
			Push2(n)
		}
	}

	return out
}
