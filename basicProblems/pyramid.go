package basicProblems

import (
	"fmt"
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
	pyramidNum := 2*n-1
	for i:=0;i<n;i++{
		tmpStr := ""

		for j:=0;j<pyramidNum;j++{
			current := 2*(i+1)-1
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