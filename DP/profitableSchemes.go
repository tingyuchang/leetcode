package DP

import "fmt"

func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	return profitableSchemesTopDown(n, minProfit, group, profit)
}

func dpProfitableSchemesV1(n int, minProfit int, group []int, profit []int) int {
	var mod int = 1e9 + 7

	dp := make([][]int, minProfit+1)

	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for k := 0; k < len(group); k++ {
		g := group[k]
		p := profit[k]

		for i := minProfit; i >= 0; i-- {
			for j := n - g; j >= 0; j-- {
				dp[min(i+p, minProfit)][j+g] += dp[i][j]
			}
		}
	}
	fmt.Println(dp)
	sum := 0

	for i := 0; i < len(dp[minProfit]); i++ {
		sum = (sum + dp[minProfit][i]) & mod
	}
	return sum
}

func dpProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	var mod int = 1e9 + 7
	dp := make([][]int, n+1)

	for i, _ := range dp {
		dp[i] = make([]int, minProfit+1)

	}
	dp[0][0] = 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for k := 0; k < len(group); k++ {
		g := group[k]
		p := profit[k]

		for i := n; i >= g; i-- {
			for j := minProfit; j >= 0; j-- {
				dp[i][j] = (dp[i][j] + dp[i-g][max(0, j-p)]) % mod
				fmt.Printf("%d, %d\t", i, j)
				fmt.Println(dp)
			}
		}
	}

	sum := 0

	for i := 0; i < len(dp); i++ {
		sum = (sum + dp[i][minProfit]) % mod
	}
	return sum
}

func profitableSchemesTopDown(n int, minProfit int, group []int, profit []int) int {
	memo := make([][][]int, len(group)+1)

	for i, _ := range memo {
		memo[i] = make([][]int, n+1)
		for j, _ := range memo[i] {
			memo[i][j] = make([]int, minProfit+1)
			for k := 0; k < len(memo[i][j]); k++ {
				memo[i][j][k] = -1
			}
		}
	}

	return dpProfitableSchemesTopDown(0, 0, 0, minProfit, n, group, profit, memo)
}

func dpProfitableSchemesTopDown(position, count, profit, minProfit, n int, group, profits []int, memo [][][]int) int {
	if position == len(group) {

		if profit >= minProfit {
			return 1
		}
		return 0
	}

	if memo[position][count][profit] != -1 {
		return memo[position][count][profit]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	total := dpProfitableSchemesTopDown(position+1, count, profit, minProfit, n, group, profits, memo)
	var mod int = 1e9 + 7
	if count+group[position] <= n {
		total += dpProfitableSchemesTopDown(position+1, count+group[position], min(minProfit, profit+profits[position]), minProfit, n, group, profits, memo)
	}

	memo[position][count][profit] = total % mod
	return memo[position][count][profit]
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	res := 0
	combination(group, profit, 0, 0, 0, n, minProfit, &res)
	return res
}

func combination(group, profit []int, currentIndex, currentProfit, needManPower, numOfPeople, minProfit int, res *int) {
	if currentIndex > len(group) || needManPower > numOfPeople {
		return
	}

	if currentProfit >= minProfit {
		*res++
	}

	for i := currentIndex; i < len(group); i++ {
		currentProfit += profit[i]
		needManPower += group[i]
		combination(group, profit, i+1, currentProfit, needManPower, numOfPeople, minProfit, res)
		currentProfit -= profit[i]
		needManPower -= group[i]

	}

}
