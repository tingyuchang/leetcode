package Strings

func ReverseString(s []byte) {
	l := 0
	r := len(s) - 1

	for r > l {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func ReverseVowels(s string) string {
	sb := make([]byte, len(s))

	l := 0
	r := len(s) - 1

	for r > l {
		both := true
		if s[l] != 'a' &&
			s[l] != 'e' &&
			s[l] != 'i' &&
			s[l] != 'o' &&
			s[l] != 'u' &&
			s[l] != 'A' &&
			s[l] != 'E' &&
			s[l] != 'I' &&
			s[l] != 'O' &&
			s[l] != 'U' {
			sb[l] = s[l]
			l++
			both = false
		}
		if s[r] != 'a' &&
			s[r] != 'e' &&
			s[r] != 'i' &&
			s[r] != 'o' &&
			s[r] != 'u' &&
			s[r] != 'A' &&
			s[r] != 'E' &&
			s[r] != 'I' &&
			s[r] != 'O' &&
			s[r] != 'U' {
			sb[r] = s[r]
			r--
			both = false
		}

		if !both {
			continue
		}

		sb[l], sb[r] = s[r], s[l]

		l++
		r--
	}

	if r == l {
		sb[l] = s[l]
	}

	return string(sb[:])
}
