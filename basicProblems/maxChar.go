package basicProblems

func maxChar(s string) string {
	tmp := make(map[string]int)
	maximum := 0
	var result string

	for _,v := range s {
		count, ok := tmp[string(v)]
		if ok {
			tmp[string(v)] = count+1
			if count+1 > maximum {
				maximum = count+1
				result = string(v)
			}
		} else {
			tmp[string(v)] = 1
		}
	}

	if maximum == 0 {
		return s
	}

	return result
}