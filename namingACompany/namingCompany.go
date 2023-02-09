package namingACompany

func DistinctNames(ideas []string) int64 {
	var result int64
	original := make(map[string]int64)
	for i, v := range ideas {
		original[v] = int64(i)
	}
	keys := make([]string, 26)
	buckets := make(map[string][]string)

	for _, v := range ideas {
		initial := string(v[0])

		_, ok := buckets[initial]

		if !ok {
			keys = append(keys, initial)
		}

		buckets[initial] = append(buckets[initial], v)

	}

	for i := 0; i < len(keys); i++ {
		// ex: {"a": {"abc", "ad"}}
		for j := i + 1; j < len(keys); j++ {
			mutualNum := 0
			a := buckets[keys[i]]
			b := buckets[keys[j]]
			temp := make(map[string]int, len(a)+len(b))

			for _, v := range a {
				v = v[1:]
				temp[v] = 1
			}
			for _, v := range b {
				v = v[1:]
				_, ok := temp[v]
				if ok {
					mutualNum++
				}
			}
			result += 2 * int64((len(a)-mutualNum)*(len(b)-mutualNum))

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
