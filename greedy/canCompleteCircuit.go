package greedy

func CanCompleteCircuit(gas []int, cost []int) int {
	return canCompleteCircuit(gas, cost)
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, currentGas, ans := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGas += gas[i] - cost[i]

		if currentGas < 0 {
			ans = i + 1
			currentGas = 0
		}
	}

	if totalGas >= 0 {
		return ans
	}

	return -1

}

/*
approach 1

cacluate all results

TO (n^2)
SO (n)

approach 2

calculate totalGas & currentGas in the same loop

if currentGas < 0, then ans = i+1

if totalGas < 0, return -1

*/
