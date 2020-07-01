package intToRoman

import "fmt"

func IntToRoman(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman string

	for i :=0; i < len(value) && num >= 0; i++ {
		for num >= value[i] {
			num -= value[i]
			roman = fmt.Sprintf("%s%v", roman, symbols[i])
		}
	}

	return roman
}
