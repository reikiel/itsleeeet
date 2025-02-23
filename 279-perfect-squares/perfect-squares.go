func numSquares(n int) int {
    dp := make([]int, n+1)

    // need make max to compare later
    for i:=1; i<=n; i++ {
        dp[i] = math.MaxInt32
    }

    // dp[0] will be 0 -> dont need anything for 0
    for i := 1; i <= n; i++ {
        for j := 1; j*j <= i; j++ {
            dp[i] = min(dp[i-j*j]+1, dp[i])

        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}