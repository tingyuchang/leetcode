package Strings

import "fmt"

func BullsCowsGetHint(secret string, guess string) string {
	return getHint(secret, guess)
}

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	cache := make(map[byte]int)

	for i := 0; i < len(secret); i++ {
		cache[secret[i]]++
	}

	for i := 0; i < len(guess); i++ {

		if guess[i] == secret[i] {
			bulls++

			// remove secret[i] in result
			if cache[guess[i]] <= 0 {
				cows--
			} else {
				cache[guess[i]]--
			}
			continue
		}

		// index not match, check is cow or not
		_, ok := cache[guess[i]]

		if ok && cache[guess[i]] > 0 {
			cows++
			cache[guess[i]]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)

}
