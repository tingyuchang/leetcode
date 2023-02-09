package Strings

import (
	"fmt"
	"math"
)

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

func TitleToNumber(columnTitle string) int {
	n := 0
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}

	for i := len(columnTitle) - 1; i >= 0; i-- {
		digit := m[string(columnTitle[i])] * int(math.Pow(26, float64(len(columnTitle)-i-1)))
		fmt.Println(digit)
		n += digit
	}

	return n
}
