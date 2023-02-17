package Strings

import "fmt"

func CharacterReplacement(s string, k int) int {
	l := 0
	r := 0
	mostFreq := 0
	longest := 0
	counts := make([]int, 26)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for r < len(s) {
		counts[s[r]-'A']++
		mostFreq = max(mostFreq, counts[s[r]-'A'])

		// if r increases and mostFreq also increases then we don't need to remove last element
		// mostFreq will ekkp the size of this frame
		if (r-l+1)-mostFreq > k {
			counts[s[l]-'A']--
			l++
		}
		longest = max(longest, r-l+1)
		r++
	}
	return longest
}

/*
Sliding Window?
*/

type Path struct {
	Val         string
	Start       int
	End         int
	Replacement int
}

func CharacterReplacementSlow(s string, k int) int {

	paths := make(map[string]*Path)

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		path, ok := paths[c]

		if !ok {
			path := &Path{
				Val:         c,
				Start:       i,
				End:         i,
				Replacement: k,
			}

			paths[c] = path
		} else {
			if path.End == i-1 {
				path.End = i
			} else {
				// ABAABA: start = 0 end = 3 replacement = 0
				lack := i - path.End - 1
				if path.Replacement > 0 && path.Replacement >= lack {
					path.Replacement = path.Replacement - lack
					path.End = i
				} else {
					// KRSCDCSONAJNHLBMDQGIFCPEKPOHQIHLTDIQGEKLRLCQNBOHNDQGHJPNDQPERNFSSSRDEQLFPCCCARFMDLHADJADAGNNSBNCJQOF
					// SDSSMESS
					replacement := k
					start := i - 1
					for replacement >= 0 && start >= 0 {
						c2 := string(s[start])
						if c2 != c {
							replacement--
						}
						if replacement < 0 {
							replacement = 0
							break
						}
						start--
					}
					start++
					if i-start >= path.End-path.Start {
						path.Start = start
						path.End = i
						path.Replacement = replacement
					}

				}
			}
		}
	}

	maxVal := 0
	for _, v := range paths {
		distance := (v.End - v.Start) + 1
		if distance < len(s) {
			distance += min(v.Replacement, len(s)-distance)
		}
		if distance > maxVal {
			fmt.Println(v)
			maxVal = distance
		}
	}
	return maxVal
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
