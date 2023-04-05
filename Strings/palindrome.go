package Strings

func LongestPalindrome(s string) int {
	return longestPalindrome(s)
}

func longestPalindrome(s string) int {
	res := 0
	cache := make(map[byte]int)
	for i, _ := range s {
		cache[s[i]]++

		if cache[s[i]] == 2 {
			res += 2
			cache[s[i]] = 0
		}
	}

	if res == len(s) {
		return res
	}

	res++

	return res
}

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	y := reverse(x)

	if x == y {
		return true
	}

	return false
}

func reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = x * -1
	}

	n := []int{}
	count, result := 0, 0
	for float64(x)/10.0 > 0 {
		n = append(n, x%10)
		x = x / 10
		count++
	}

	for i := 0; i < count; i++ {
		result = result*10 + n[i]
	}

	if result > 2147483647 {
		return 0
	}

	if isNegative {
		return result * -1
	}
	return result
}
