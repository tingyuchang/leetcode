package basicProblems

import (
	"fmt"
	"reflect"
	"regexp"
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

func anagrams(a, b string) bool {
	re := regexp.MustCompile(`\w`)
	regexStrA := re.FindAll([]byte(a), -1)
	regexStrB := re.FindAll([]byte(b), -1)

	mapA := make(map[string]int)
	mapB := make(map[string]int)
	for _,v := range regexStrA {
		key := strings.ToLower(fmt.Sprintf("%q", v))
		_, ok := mapA[key]
		if ok {
			mapA[key] = mapA[key]+1
		}else {
			mapA[key] = 1
		}
	}
	for _,v := range regexStrB {
		key := strings.ToLower(fmt.Sprintf("%q", v))
		_, ok := mapB[key]
		if ok {
			mapB[key] = mapB[key]+1
		}else {
			mapB[key] = 1
		}
	}
	return reflect.DeepEqual(mapA, mapB)
}

