package Sort

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func LetterCasePermutation(s string) []string {
	return letterCasePermutation(s)
}

func letterCasePermutation(s string) []string {
	temp := make(map[string]struct{}, 0)

	numOfLetters := 0
	newS := make([]rune, len(s))
	for i, v := range s {
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
			newS[i] = v
			continue
		}

		newS[i] = unicode.ToLower(v)
		numOfLetters++
	}

	temp[string(newS)] = struct{}{}

	for i := 1; i <= numOfLetters; i++ {
		combinationLetterCase(newS, 0, i, &temp, make([]rune, len(s)))
	}

	res := make([]string, len(temp))
	i := 0
	for k, _ := range temp {
		res[i] = k
		i++
	}

	return res
}

func combinationLetterCase(s []rune, currentIndex, currentUppercaseNum int, res *map[string]struct{}, currentRes []rune) {
	if currentUppercaseNum < 0 {
		return
	}

	if currentUppercaseNum == 0 {
		_, ok := (*res)[string(s)]

		if !ok {
			(*res)[string(s)] = struct{}{}
		}
		return
	}

	for i := currentIndex; i < len(s); i++ {

		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') {
			combinationLetterCase(s, i+1, currentUppercaseNum, res, currentRes)
		} else {

			// to uppercase
			s[i] = unicode.ToUpper(s[i])
			combinationLetterCase(s, i+1, currentUppercaseNum-1, res, currentRes)
			// revert o lowercase
			s[i] = unicode.ToLower(s[i])
		}

	}
}

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	duplicated := make(map[string]int)
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	result = append(result, append([]int{}, nums...))
	duplicated[toString(nums)] = 1
	if len(nums) == 1 {
		return result
	}

	// n!/(a!*b!*c!), a, b, c means numbers of duplicated elements
	n := 1
	//count := make(map[int]int, 0)
	for i := 1; i <= len(nums); i++ {
		n = n * i
		//val, ok := count[nums[i-1]]
		//if !ok {
		//	count[nums[i-1]] = 1
		//} else {
		//	count[nums[i-1]] = val * (val + 1)
		//}
	}

	//for _, v := range count {
	//	if v != 1 {
	//		n = n / v
	//	}
	//}

	for n > 0 {
		n--
		i := len(nums) - 1

		for nums[i-1] >= nums[i] {
			i--
			if i == 0 {
				return result
			}
		}

		j := len(nums) - 1
		for j > i && nums[j] <= nums[i-1] {
			j--
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
		reverse(nums, i)
		temp := append([]int{}, nums...)
		unique := toString(temp)
		_, ok := duplicated[unique]
		if !ok {
			result = append(result, temp)
			duplicated[unique] = 1
		}

	}

	return result
}

func toString(nums []int) string {
	s := ""

	for i := 0; i < len(nums); i++ {
		fmt.Println(strconv.Itoa(nums[i]))
		s = s + strconv.Itoa(nums[i])
	}

	return s
}

func Permutation(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	result = append(result, append([]int{}, nums...))
	if len(nums) == 1 {
		return result
	}
	n := 1
	for i := 1; i <= len(nums); i++ {
		n = n * i
	}

	for n > 0 {
		n--
		if n == 0 {
			return result
		}
		i := len(nums) - 1

		for nums[i-1] >= nums[i] {
			i--
		}

		j := len(nums) - 1
		for j > i && nums[j] <= nums[i-1] {
			j--
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
		reverse(nums, i)
		temp := append([]int{}, nums...)
		result = append(result, temp)
	}

	return result
}

func NextPermutation(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	if len(nums) == 1 {
		return nums
	}

	i := len(nums) - 1
	// find the largest index i such that nums[i-1] is less than nums[i]
	// if nums[i-1] smaller than nums[i], means we can find next permutation
	// If i is not the first index of the slice, then sub-slice nums[i…n) is stored in descending order
	// i.e. nums[i-1] < nums[i] >= nums[i+1] >= nums[i+2]
	for nums[i-1] >= nums[i] {
		i = i - 1
		if i == 0 {
			reverse(nums, 0)
			return nums
		}
	}

	j := len(nums) - 1
	// find highest index j such nums[j] is greater than nums[i-1]
	// i.e. 1,2,3,4,6,5 => 1,2,3,5,6,4
	for j > i && nums[j] <= nums[i-1] {
		j = j - 1
	}
	// swap nums[i-1] & nums[j]
	nums[i-1], nums[j] = nums[j], nums[i-1]
	// reverse nums[i:]
	reverse(nums, i)
	return nums
}

func reverse(nums []int, n int) {
	i, j := n, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
