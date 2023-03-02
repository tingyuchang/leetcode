package Strings

import (
	"fmt"
	"strconv"
)

func Compress(chars []byte) int {
	position := 0
	var current byte
	counter := 0
	for i, v := range chars {
		if i == 0 {
			position = i
			current = v
			counter++
			continue
		}

		if v == current {
			counter++
		} else {
			if counter > 1 {
				total := ""

				for counter > 0 {
					total = strconv.Itoa(counter%10) + total
					counter = counter / 10
				}

				for j := 0; j < len(total); j++ {
					chars[position+j+1] = total[j]
				}

				position += len(total) + 1
				if chars[position] != v {
					chars[position] = v
				}
				current = v
				counter = 1
			} else {
				position++
				if chars[position] != v {
					chars[position] = v
				}
				current = v
				counter = 1
			}
		}

		if i == len(chars)-1 {
			if counter > 1 {
				total := ""

				for counter > 0 {
					total = strconv.Itoa(counter%10) + total
					counter = counter / 10
				}

				for j := 0; j < len(total); j++ {
					chars[position+j+1] = total[j]
				}

				position += len(total)
			}
		}

		fmt.Println(position)
	}
	position++
	fmt.Println(position, chars)

	return len(chars[:position])
}
