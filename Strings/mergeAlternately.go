package Strings

func mergeAlternately(word1 string, word2 string) string {
	newWord := make([]byte, len(word1)+len(word2))
	m, n, l := 0, 0, 0

	for m < len(word1) && n < len(word2) {
		newWord[l] = word1[m]
		m++
		l++
		newWord[l] = word2[n]
		n++
		l++
	}

	for l < len(newWord) {
		if m < len(word1) {
			newWord[l] = word1[m]
			m++
		} else {
			newWord[l] = word2[n]
			n++
		}

		l++
	}

	return string(newWord)

}
