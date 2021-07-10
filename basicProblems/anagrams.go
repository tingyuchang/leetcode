package basicProblems

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Check to see if two provided strings are anagrams of each other.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//   anagrams('rail safety', 'fairy tales') --> True
//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//   anagrams('Hi there', 'Bye there') --> False

func anagrams2 (a, b string) bool {
	re := regexp.MustCompile(`\w`)
	regexA := re.FindAllString(strings.ToLower(a), -1)
	regexB := re.FindAllString(strings.ToLower(b), -1)
	sort.Strings(regexA)
	sort.Strings(regexB)
	return strings.Join(regexA, "") == strings.Join(regexB, "")
}

func anagrams(a, b string) bool {
	mapA := buildCharMap(a)
	mapB := buildCharMap(b)

	for k,_ := range mapA {
		if mapA[k] != mapB[k] {
			return false
		}
	}
	for k,_ := range mapB {
		if mapA[k] != mapB[k] {
			return false
		}
	}

	return true
}

func buildCharMap(s string)map[string]int {
	charMap := make(map[string]int)
	re := regexp.MustCompile(`\w`)
	regexStr := re.FindAll([]byte(s), -1)
	for _,v := range regexStr {
		key := strings.ToLower(fmt.Sprintf("%q", v))
		_, ok := charMap[key]
		if ok {
			charMap[key] = charMap[key]+1
		}else {
			charMap[key] = 1
		}
	}

	return charMap
}

