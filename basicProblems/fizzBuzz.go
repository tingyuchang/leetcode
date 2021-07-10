package basicProblems

import "fmt"

/*
// --- Directions
// Write a program that console logs the numbers
// from 1 to n. But for multiples of three print
// “fizz” instead of the number and for the multiples
// of five print “buzz”. For numbers which are multiples
// of both three and five print “fizzbuzz”.
// --- Example
//   fizzBuzz(5);
//   1
//   2
//   fizz
//   4
//   buzz
 */

const fizz, buzz, fizzbuzz = "fizz", "buzz", "fizzbuzz"

func fizzBuzz(n int) []string{
	var result []string

	for i:=0;i<n;i++ {
		if (i+1)%3 == 0 {
			if (i+1)%5 == 0 {
				result = append(result, fizzbuzz)
			} else {
				result = append(result, fizz)
			}
		} else if (i+1)%5 == 0 {
			result = append(result, buzz)
		} else {
			result = append(result, fmt.Sprintf("%v",i+1))
		}
	}

	return result
}
