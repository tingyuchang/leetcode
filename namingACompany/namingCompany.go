package namingACompany

func DistinctNames(ideas []string) int64 {
	var result int64
	original := make(map[string]int)
	for i, v := range ideas {
		original[v] = i
	}
	keys := make(map[string][]string)

	for _, v := range ideas {
		initial := string(v[0])
		keys[initial] = append(keys[initial], v)
	}

	for i := 0; i < len(ideas); i++ {
		initial := string(ideas[i][0])
		//exist := map[string]int{initial: 0}
		for k, v := range keys {
			// do not look for same initial letter
			if initial == k {
				continue
			}

			for j := 0; j < len(v); j++ {
				a, b := swapInitialLetter(ideas[i], v[j])
				_, ok := original[a]

				if ok {
					break
				}

				_, ok = original[b]
				if ok {
					continue
				}
				_, ok = original[a+" "+b]

				if !ok {
					result++
					original[a+" "+b] = int(result + 100)
				}
				_, ok = original[b+" "+a]

				if !ok {
					result++
					original[b+" "+a] = int(result + 100)
				}
			}
		}
	}
	return result
}

func DistinctNamesDev(ideas []string) int64 {
	result := make([]string, 0)
	original := make(map[string]int)
	for i, v := range ideas {
		original[v] = i
	}
	keys := make(map[string][]string)

	for _, v := range ideas {
		initial := string(v[0])
		keys[initial] = append(keys[initial], v)
	}

	for i := 0; i < len(ideas); i++ {
		initial := string(ideas[i][0])
		//exist := map[string]int{initial: 0}
		for k, v := range keys {
			// do not look for same initial letter
			if initial == k {
				break
			}

			for j := 0; j < len(v); j++ {
				a, b := swapInitialLetter(ideas[i], v[j])

				_, ok := original[a]

				if ok {
					break
				}

				_, ok = original[b]
				if ok {
					continue
				}
				result = append(result, a+" "+b)
				result = append(result, b+" "+a)
			}
		}
	}
	return int64(len(result))
}

func swapInitialLetter(a, b string) (string, string) {
	initialA := a[0]
	initialB := b[0]
	a = string(initialB) + a[1:]
	b = string(initialA) + b[1:]
	return a, b
}
