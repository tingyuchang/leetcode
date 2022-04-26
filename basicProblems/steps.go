package basicProblems

import "fmt"

func steps(n int) []string {
	result := make([]string, n)
	for i:=0;i< n;i++ {
		output := ""
		tmp := i
		for j:=0;j< n;j++ {
			if tmp > -1 {
				output = output+"#"
			} else {
				output = output+" "
			}
			tmp--
		}
		fmt.Println(output)
		result[i] = output
	}
	return result
}