package letterCombinations

import "fmt"

var result []string
var template map[string][]string
func LetterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	template = make(map[string][]string)
	template["2"] = []string{"a", "b", "c"}
	template["3"] = []string{"d", "e", "f"}
	template["4"] = []string{"g", "h", "i"}
	template["5"] = []string{"j", "k", "l"}
	template["6"] = []string{"m", "n", "o"}
	template["7"] = []string{"p", "q", "r", "s"}
	template["8"] = []string{"t", "u", "v"}
	template["9"] = []string{"w", "x", "y", "z"}

	var nums []string

	for _,v := range digits {
		nums = append(nums, string(v))
	}

	result = []string{}
	for _,v := range template[nums[0]]{
		combination := v
		result = append(result, combination)
	}


	for i := 1; i< len(nums); i++ {
		// 3, 4, 5, 6...etc
		options := template[nums[i]]
		temp := []string{}
		for j :=0; j < len(result); j++ {
			// a,b ,c
			for k := 0; k <len(options); k++ {
				combination := fmt.Sprintf("%v%v", result[j], options[k])
				temp = append(temp, combination)
			}
		}
		result = temp
	}


	return result
}

//func backTrack (combination string, nextDigit []string) {
//	if len(nextDigit) == 0 {
//		if len(combination) > 0 {
//			result = append(result, combination)
//		}
//	} else {
//		for _,v := range template[nextDigit[0]] {
//			backTrack(fmt.Sprintf("%v%v", combination, v), nextDigit[1:])
//		}
//	}
//}


