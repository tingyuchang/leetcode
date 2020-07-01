package intToRoman

import "fmt"

func IntToRoman(num int) string {
	thousand := num/1000
	hundred :=  (num%1000)/100
	ten := (num%100)/10
	zero := num%10

	var roman string

	for i := 0; i < thousand; i++ {
		roman = fmt.Sprintf("%s%s", roman, "M")
	}

	switch hundred {
		case 1:
			roman = fmt.Sprintf("%s%s", roman, "C")
		case 2:
			roman = fmt.Sprintf("%s%s", roman, "CC")
		case 3:
			roman = fmt.Sprintf("%s%s", roman, "CCC")
		case 4:
			roman = fmt.Sprintf("%s%s", roman, "CD")
		case 5:
			roman = fmt.Sprintf("%s%s", roman, "D")
		case 6:
			roman = fmt.Sprintf("%s%s", roman, "DC")
		case 7:
			roman = fmt.Sprintf("%s%s", roman, "DCC")
		case 8:
			roman = fmt.Sprintf("%s%s", roman, "DCCC")
		case 9:
			roman = fmt.Sprintf("%s%s", roman, "CM")
	}

	switch ten {
		case 1:
			roman = fmt.Sprintf("%s%s", roman, "X")
		case 2:
			roman = fmt.Sprintf("%s%s", roman, "XX")
		case 3:
			roman = fmt.Sprintf("%s%s", roman, "XXX")
		case 4:
			roman = fmt.Sprintf("%s%s", roman, "XL")
		case 5:
			roman = fmt.Sprintf("%s%s", roman, "L")
		case 6:
			roman = fmt.Sprintf("%s%s", roman, "LX")
		case 7:
			roman = fmt.Sprintf("%s%s", roman, "LXX")
		case 8:
			roman = fmt.Sprintf("%s%s", roman, "LXXX")
		case 9:
			roman = fmt.Sprintf("%s%s", roman, "XC")
	}

	switch zero {
		case 1:
			roman = fmt.Sprintf("%s%s", roman, "I")
		case 2:
			roman = fmt.Sprintf("%s%s", roman, "II")
		case 3:
			roman = fmt.Sprintf("%s%s", roman, "III")
		case 4:
			roman = fmt.Sprintf("%s%s", roman, "IV")
		case 5:
			roman = fmt.Sprintf("%s%s", roman, "V")
		case 6:
			roman = fmt.Sprintf("%s%s", roman, "VI")
		case 7:
			roman = fmt.Sprintf("%s%s", roman, "VII")
		case 8:
			roman = fmt.Sprintf("%s%s", roman, "VIII")
		case 9:
			roman = fmt.Sprintf("%s%s", roman, "IX")
	}


	return roman
}
