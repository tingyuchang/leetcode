package Tree

/*

assume we had n nums
from  0 to n, we select i as root,
then the valid tree should be
		   i
      /        \
   [0-i-1]   [i+1 - n]

so the expect result should sum all [1 - n] element as root's possible results

numsOfTree(i, n ) = numsOfTree(i-1, n) * numsOfTree(n-i-1, n)


*/
func numsTrees(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
