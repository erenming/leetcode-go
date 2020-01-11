package climbStairs

// f(n) = f(n-1) + f(n-2)
// f(2):
// 		f(1) + f(0)
// func climbStairs(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
