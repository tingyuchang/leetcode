package validParentheses

func IsValid(s string) bool {
	if s == "" {
		return true
	}

	var temp []string

	for _,v := range s {
		n := string(v)
		if n == "(" || n == "[" || n == "{" {
			temp = append(temp, n)
		} else {
			if len(temp) == 0 {
				return false
			} else {
				var start string
				if n == ")" {
					start = "("
				} else if n == "]" {
					start = "["
				}  else if n == "}" {
					start = "{"
				}

				if temp[len(temp)-1] != start{
					return false
				} else {
					temp = temp[:len(temp)-1]
				}
			}
		}

	}
	

	return len(temp) == 0
}