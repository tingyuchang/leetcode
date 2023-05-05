package Strings

func MaxVowels(s string, k int) int {
	return maxVowels(s, k)
}

func maxVowels(s string, k int) int {
	numsOfVowels, ans := 0, 0
	l, r := 0, 0

	for r < len(s) {
		if isVowel(s[r]) {
			numsOfVowels++
		}

		if r-l+1 == k {
			if numsOfVowels > ans {
				ans = numsOfVowels
			}
			if isVowel(s[l]) {
				numsOfVowels--
			}
			l++

		}
		r++
	}

	return ans
}

func isVowel(c byte) bool {
	if c == 'a' ||
		c == 'e' ||
		c == 'i' ||
		c == 'o' ||
		c == 'u' {
		return true
	}

	return false
}

/*
sliding window

numsOfVowels
ans
l, r := 0


for r < len(s) {
	if s[r] == VOWELS {
		numsOfVowels++
	}

	if r - l+1  == k {

	}

	r++
}


*/
