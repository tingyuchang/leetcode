package Strings

func ConvertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	s := ""
	m := []string{"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y"}

	for columnNumber/26 > 0 || columnNumber%26 > 0 {
		mo := columnNumber % 26
		quotient := columnNumber / 26
		s = m[mo] + s
		if mo == 0 && quotient > 0 {
			columnNumber = quotient - 1
		} else {
			columnNumber = quotient
		}
	}

	return s
}
