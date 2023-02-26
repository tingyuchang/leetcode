package foodpanda

// 'a'
func Solution2(S string) string {
	// Implement your solution here
	removed := S[0]
	result := ""
	for i := 1; i < len(S); i++ {
		if S[i] >= removed {
			removed = S[i]
			if i == len(S)-1 || S[i+1] < removed {
				result = string(append([]byte(S)[:i], []byte(S)[i+1:]...))
				break
			}
		}
	}

	return result
}
