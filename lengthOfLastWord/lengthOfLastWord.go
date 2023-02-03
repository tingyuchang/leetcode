package lengthOfLastWord

func LengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if length == 0 {
				continue
			} else {
				break
			}
		} else {
			length++
		}
	}

	return length
}
