package main

import "strings"

var stack3 []string

func Push3(s string) {
	stack3 = append(stack3, s)
}

func Pop3() string {
	item := stack3[len(stack)-1]
	stack3 = stack3[:len(stack)-1]
	return item
}

func Len3() int {
	return len(stack3)
}

func simplifyPath(path string) string {
	s := ""
	stack3 = make([]string, 0)

	for i := 0; i < len(path); i++ {
		c := path[i]
		if c == '/' {
			if s != "" {
				if s == "." {
					// continue
				} else if s == ".." {
					//fmt.Println(stack)
					if Len3() > 0 {
						Pop3()
					}
				} else {
					Push3(s)
				}

				s = ""
			}
		} else {
			s += string(path[i])
		}
	}

	//fmt.Println(s)
	if s != "" {
		if s == "." {
			// continue
		} else if s == ".." {
			if Len3() > 0 {
				Pop3()
			}
		} else {
			Push3(s)
		}
	}

	//fmt.Println(stack)
	ans := strings.Join(stack3, "/")
	return "/" + ans
}
