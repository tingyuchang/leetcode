package basicProblems

import (
	"fmt"
	"math"
)

/*
The function should console log a pyramid shape
with N levels using the # character.  Make sure the
pyramid has spaces on both the left *and* right hand sides
--- Examples
   pyramid(1)
       '#'
   pyramid(2)
       ' # '
       '###'
   pyramid(3)
       '  #  '
       ' ### '
       '#####'
*/
func pyramid(n int) []string {
	var results []string
	var pyramidNum int

	if n == 1 {
		pyramidNum = 1
	} else {
		pyramidNum = int(math.Pow(2, float64(n-1)))+1
	}
	for i:=0;i<n;i++{
		tmpStr := ""

		for j:=0;j<pyramidNum;j++{
			current :=0
			if i == 0 {
				current = 1
			} else {
				current = int(math.Pow(2, float64(i)))+1
			}
			if j < (pyramidNum-current)/2 || j >= (pyramidNum+current)/2   {
				tmpStr+=" "
			} else {
				tmpStr+="#"
			}
		}
		fmt.Println(tmpStr)
		results = append(results, tmpStr)
	}
	return results
}