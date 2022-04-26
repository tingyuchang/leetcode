package basicProblems

import (
	"strings"
)

/*
// Write a function that returns the number of vowels
// used in a string.  Vowels are the characters 'a', 'e'
// 'i', 'o', and 'u'.
// --- Examples
//   vowels('Hi There!') --> 3
//   vowels('Why do you ask?') --> 4
//   vowels('Why?') --> 0
 */


func checkVowels(s string) int {
	var vowels = map[string]bool{
		"a": false,
		"e": false,
		"i": false,
		"o": false,
		"u": false,
	}
	sum := 0
	for _,v := range strings.Split(s, "") {
		lowercase := strings.ToLower(v)
		exist,ok := vowels[lowercase]
		if !exist && ok {
			sum+=1
			vowels[lowercase] = true
		}
	}
	return sum
}