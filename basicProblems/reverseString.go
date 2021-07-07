package basicProblems

import "strings"

func reverseString(s string) string{
	reverseString := strings.Split(s,"")

	for i,j:=0,len(reverseString)-1;i<j;i,j=i+1, j-1{
		reverseString[i], reverseString[j] = reverseString[j], reverseString[i]
	}
	return strings.Join(reverseString, "")
}

func reverseString2(s string) string {
	var reverseString string
	for i:=0;i<len(s);i++{
		reverseString+=string(s[len(s)-1-i])
	}
	return reverseString
}
