package main

var stack []rune

func Push(c rune) {
	stack = append(stack, c)
}

func Pop() rune {
	c := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return c
}

func Len() int {
	return len(stack)
}

func isValid(s string) bool {
	stack = make([]rune, 0)

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			Push(c)
		} else {
			if Len() > 0 {
				curr := Pop()
				if c == ')' && curr != '(' {
					return false
				} else if c == ']' && curr != '[' {
					return false
				} else if c == '}' && curr != '{' {
					return false
				}
			} else {
				return false
			}

		}
	}

	if Len() > 0 {
		return false
	}

	return true
}
