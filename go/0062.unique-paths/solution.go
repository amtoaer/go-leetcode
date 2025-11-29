// Created by amtoaer at 2025/11/29 11:43
// leetgo: 1.4.15
// https://leetcode.cn/problems/unique-paths/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for range m - 1 {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := Deserialize[int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := uniquePaths(m, n)

	fmt.Println("\noutput:", Serialize(ans))
}
