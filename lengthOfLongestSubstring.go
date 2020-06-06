package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var tmp []string
	var result = 0
	for i,v := range s {
		log("index: %d, value: %v type:: %T\tother: %v type: %T\n", i, string(v), v, s[i+1:], s[i+1:])
		currentString := ""
		starPoint := i
		// endPoint := strings.IndexAny(s[i+1:], currentString)
		endPoint := len(s)-1
		if endPoint == -1 {
			endPoint = len(s)-1
		}
		log("Start loop\nStartPoint: %v\tEndpoint: %v\n", starPoint, endPoint)
		for starPoint <= endPoint {
			currentString = fmt.Sprintf("%s%s", currentString, string(s[starPoint]))
			checkExist := strings.IndexAny(s[starPoint+1:], currentString)
			log("CurrentString: %s\tCheckExist: %v\n", currentString, checkExist)
			if checkExist == 0 {
				log("End Loop\n")
				tmp = append(tmp, currentString)
				if len(currentString) > result {
					result = len(currentString)
				}
				break
			}


			log("At: %s\t current: %s\n", string(v), currentString)
			starPoint++
		}

	}


	fmt.Printf("Result: %d\n%v ", result, tmp)

	return result
}