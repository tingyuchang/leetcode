package basicProblems

// Given a string, return true if the string is a palindrome
// or false if it is not.  Palindromes are strings that
// form the same word if it is reversed. *Do* include spaces
// and punctuation in determining if the string is a palindrome.
// --- Examples:
//   palindrome("abba") === true
//   palindrome("abcdefg") === false

func palindrome(s string)bool {
	for i, j  := 0, len(s)-1; i < j/2; i, j = i+1, j-1 {
		if s[i] != s[j]{
			return false
		}
	}
	return true
}
func palindrome2(s string)bool {
	return reverseString3(s) == s
}