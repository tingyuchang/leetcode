package Math

import "math"

func BulbSwitch(n int) int {
	return bulbSwitch(n)
}

/*
TO(n^2+n)
SO(n)
*/
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
	//bulbs := make([]bool, n)
	//
	//countDown := 0
	//for countDown < n {
	//	countDown++
	//	for i := 0; i < len(bulbs); i++ {
	//		if (i+1)%countDown == 0 {
	//			bulbs[i] = !bulbs[i]
	//		}
	//	}
	//}
	//
	//ans := 0
	//for i := range bulbs {
	//	if bulbs[i] {
	//		ans++
	//	}
	//}
	//
	//return ans
}
