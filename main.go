package main

import (
	"fmt"
)

var allowLog = false

func main() {
	var in = 0
	in = lengthOfLongestSubstring("abcbbc")
	//in = lengthOfLongestSubstring("bbbb")
	//in = lengthOfLongestSubstring("pwwkew")
	in--
}


func log(format string, a ...interface{}) {
	if allowLog {
		fmt.Printf(format, a...)
	}
}