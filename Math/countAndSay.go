package Math

import (
	"strings"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	previous := CountAndSay(n - 1)

	start := -1
	var result strings.Builder
	for i := 0; i < len(previous); i++ {
		if i == 0 {
			start = 0
		}

		if previous[i] != previous[start] {
			count := i - start
			result.Write([]byte{byte('0' + count), previous[start]})
			start = i
		}

		if i == len(previous)-1 {
			count := i - start + 1
			result.Write([]byte{byte('0' + count), previous[start]})
		}
	}

	return result.String()
}
func countAndSay(n int) string {
	result := "1"

	for n > 1 {
		start := -1
		var temp strings.Builder
		for i := 0; i < len(result); i++ {
			if i == 0 {
				start = 0
			}

			if result[i] != result[start] {
				count := i - start
				temp.Write([]byte{byte('0' + count), result[start]})
				start = i
			}

			if i == len(result)-1 {
				count := i - start + 1
				temp.Write([]byte{byte('0' + count), result[start]})
			}
		}

		result = temp.String()

		n--
	}

	return result
}
