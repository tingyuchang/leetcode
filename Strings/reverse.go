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