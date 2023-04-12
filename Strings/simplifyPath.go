package Strings

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{"/"}

	dirs := strings.Split(path, "/")

	for _, dir := range dirs {
		if dir == "" {
			continue
		}

		last := stack[len(stack)-1]

		switch dir {
		case "/":
			if last == "/" {
				continue
			}
		case ".":
			continue
		case "..":
			// jump to up level
			if last != "/" {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, dir)
		}
	}
	fmt.Println(stack)
	res := ""
	for i, v := range stack {
		res += v

		if i != 0 && i != len(stack)-1 {
			res += "/"
		}
	}

	return res

}
