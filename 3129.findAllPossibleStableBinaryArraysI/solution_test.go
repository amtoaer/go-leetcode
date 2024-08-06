package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3129 lang=golang
 *
 * [3129] Find All Possible Stable Binary Arrays I
 */

// @lc code=start
func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][2]int, zero+1)
	mod := int(1e9 + 7)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}
	for i := 0; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 0; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			if i > limit {
				// 中间全是 0，不满足 limit 限制，需要减掉
				dp[i][j][0] -= dp[i-limit-1][j][1]
			}
			dp[i][j][0] = (dp[i][j][0]%mod + mod) % mod
			dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
			if j > limit {
				dp[i][j][1] -= dp[i][j-limit-1][0]
			}
			dp[i][j][1] = (dp[i][j][1]%mod + mod) % mod
		}
	}
	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		zero  int
		one   int
		limit int
		want  int
	}{
		{1, 1, 2, 2},
		{1, 2, 1, 1},
		{3, 3, 2, 14},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, numberOfStableArrays(tt.zero, tt.one, tt.limit))
	}
}
